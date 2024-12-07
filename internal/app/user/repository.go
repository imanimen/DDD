package user

type UserRepository interface {
	FindByID(id string) (*User, error)
	Save(user *User) error
}