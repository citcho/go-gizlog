package application

type StoreUserCommand struct {
	ID       string
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
