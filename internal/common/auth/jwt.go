package auth

import (
	"context"
	_ "embed"
	"fmt"
	"net/http"
	"time"

	"github.com/citcho/go-gizlog/internal/auth/domain/user"
	"github.com/citcho/go-gizlog/internal/common/clock"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

//go:embed cert/secret.pem
var rawPrivKey []byte

//go:embed cert/public.pem
var rawPubKey []byte

type JWTer struct {
	privateKey, publicKey jwk.Key
	Clocker               clock.Clocker
}

func NewJWTer(c clock.Clocker) (*JWTer, error) {
	privkey, err := parse(rawPrivKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: private key: %w", err)
	}
	pubkey, err := parse(rawPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: public key: %w", err)
	}

	j := &JWTer{
		privateKey: privkey,
		publicKey:  pubkey,
		Clocker:    c,
	}

	return j, nil
}

func parse(rawKey []byte) (jwk.Key, error) {
	key, err := jwk.ParseKey(rawKey, jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}
	return key, nil
}

func (j *JWTer) GenerateToken(ctx context.Context, u *user.User) ([]byte, error) {
	token, err := jwt.NewBuilder().
		// JwtID(uuid.New().String()).
		Issuer(`github.com/citcho/go-gizlog`).
		Subject("access_token").
		IssuedAt(j.Clocker.Now()).
		Expiration(j.Clocker.Now().Add(30*time.Minute)).
		Claim("user_id", u.ID()).
		Build()
	if err != nil {
		return nil, fmt.Errorf("GenerateToken: failed to build token: %w", err)
	}

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.RS256, j.privateKey))
	if err != nil {
		return nil, err
	}

	return signed, nil
}

func (j *JWTer) GetToken(ctx context.Context, r *http.Request) (jwt.Token, error) {
	token, err := jwt.ParseRequest(
		r,
		jwt.WithKey(jwa.RS256, j.publicKey),
		jwt.WithValidate(false),
	)
	if err != nil {
		return nil, err
	}

	if err := jwt.Validate(token, jwt.WithClock(j.Clocker)); err != nil {
		return nil, fmt.Errorf("GetToken: failed to validate token: %w", err)
	}

	return token, nil
}

type userIDKey struct{}

func (j *JWTer) FillContext(r *http.Request) (*http.Request, error) {
	token, err := j.GetToken(r.Context(), r)
	if err != nil {
		return nil, err
	}

	uid, ok := token.Get("user_id")
	if !ok {
		return nil, fmt.Errorf("FillContext: failed to get user_id from token")
	}

	ctx := SetUserID(r.Context(), uid.(string))

	return r.Clone(ctx), nil
}

func SetUserID(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, userIDKey{}, uid)
}

func GetUserID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(userIDKey{}).(string)
	return id, ok
}
