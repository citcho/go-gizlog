package user

import (
	"errors"
	"regexp"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id       string
	name     string
	email    string
	password string
}

func ReConstruct(
	id string,
	name string,
	email string,
	password string,
) (*User, error) {
	u := &User{
		id:       id,
		name:     name,
		email:    email,
		password: password,
	}

	return u, nil
}

func NewUser(
	id string,
	name string,
	email string,
	password string,
) (*User, error) {
	if utf8.RuneCountInString(name) > 50 {
		return &User{}, errors.New("名前は50文字未満で入力してください。")
	}
	if err := isValidEmail(email); err != nil {
		return &User{}, err
	}
	if err := isValidPassword(password); err != nil {
		return &User{}, err
	}

	u := &User{
		id:       id,
		name:     name,
		email:    email,
		password: encrypt(password),
	}

	return u, nil
}

func (r User) ID() string {
	return r.id
}

func (r User) Name() string {
	return r.name
}
func (r User) Email() string {
	return r.email
}

func (r User) Password() string {
	return r.password
}

func isValidEmail(e string) error {
	if !(regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$`).Match([]byte(e))) {
		return errors.New("正しいメールアドレスの形式で入力してください。")
	}

	return nil
}

func isValidPassword(pw string) error {
	if len(pw) < 7 {
		return errors.New("パスワードは8文字以上入力する必要があります。")
	}
	// 英数字記号以外を使っているか判定
	if !(regexp.MustCompile("^[0-9a-zA-Z!-/:-@[-`{-~]+$").Match([]byte(pw))) {
		return errors.New("パスワードは大文字小文字英数字と記号のみサポートしています。")
	}
	reg := []*regexp.Regexp{
		// 英字が含まれるか
		regexp.MustCompile(`[[:alpha:]]`),
		// 数字が含まれるか
		regexp.MustCompile(`[[:digit:]]`),
		// 記号が含まれるか
		regexp.MustCompile(`[[:punct:]]`),
	}
	for _, r := range reg {
		if r.FindString(pw) == "" {
			return errors.New("パスワードは大文字小文字英字、数字、記号をそれぞれ1文字以上含む必要があります。")
		}
	}

	return nil
}

func encrypt(s string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
