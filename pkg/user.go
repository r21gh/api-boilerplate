package root

// User : This struct to implemenet the user fields
type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserService interface is responsible for
// creating an interface for User functionalities
type UserService interface {
	CreateUser(u *User) error
	GetByUsername(username string) (*User, error)
}