package models

type User struct {
	ID       int    `json:"user_id" db:"user_id"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role"`
}

type UserUsecaseI interface {
	GetUserById(id int) (*User, error)
	GetUserByLogin(login string) (*User, error)
	Authenticate(login string, password string) (bool, error)
	Create(user *User) (int, error)
	Update(id int, newUser *User) error
	Delete(id int) error
}

type UserRepositoryI interface {
	GetUserById(id int) (*User, error)
	GetUserByLogin(login string) (*User, error)
	Create(user *User) (int, error)
	Update(newUser *User) error
	Delete(id int) error
}
