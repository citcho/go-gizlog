package application

type LoginCommand struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
