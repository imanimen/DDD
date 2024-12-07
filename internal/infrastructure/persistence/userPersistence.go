package persistence


import (
	"errors"
	"github.com/imanimen/ddd-pr/internal/user"
)

type InMemoryUserRepository struct {
	users map[string]*user.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
    return &InMemoryUserRepository{users: make(map[string]*user.User)}
}

func (repo *InMemoryUserRepository) FindByID(id string) (*user.User, error) {
    if user, exists := repo.users[id]; exists {
        return user, nil
    }
    return nil, errors.New("user not found")
}

func (repo *InMemoryUserRepository) Save(user *user.User) error {
    repo.users[user.ID] = user
    return nil
}