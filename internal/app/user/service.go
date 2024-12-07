// Domain Service
package user

type UserService struct {
    repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) Register(name, email string) (*User, error) {
    user := &User{
        ID:        generateUUID(), // mock gen uuid function :D
        Name:      name,
        Email:     email,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
    if err := s.repo.Save(user); err != nil {
        return nil, err
    }
    return user, nil
}
