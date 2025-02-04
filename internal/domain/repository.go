package domain

type UserRepository interface {
	Save(user *User) error
	FindByID(id uint) (*User, error)
	FindAll() ([]User, error)
	Delete(id uint) error
}
