package user

type User struct {
	id       string
	email    string
	password string
}

func (r User) ID() string {
	return r.id
}

func (r User) Email() string {
	return r.email
}

func (r User) Password() string {
	return r.password
}

func NewUser(
	id string,
	email string,
	password string,
) (*User, error) {
	u := &User{
		id:       id,
		email:    email,
		password: password,
	}

	return u, nil
}
