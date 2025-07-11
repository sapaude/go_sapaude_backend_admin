package application

import (
    "context"
    "errors"

    "github.com/lupguo/go_sapaude_backend_admin/domain/entity"
    "github.com/lupguo/go_sapaude_backend_admin/domain/repository"
    "github.com/lupguo/go_sapaude_backend_admin/domain/service"
    "golang.org/x/crypto/bcrypt"
)

type UserService struct {
    UserRepo   repository.UserRepository
    JWTService service.JWTService
}

func NewUserService(repo repository.UserRepository, jwt service.JWTService) *UserService {
    return &UserService{
        UserRepo:   repo,
        JWTService: jwt,
    }
}

// CreateUser Register a new user
func (s *UserService) CreateUser(ctx context.Context, email string, rawPassword string) (*entity.User, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := &entity.User{
        Email:        email,
        PasswordHash: string(hashedPassword),
        IsActive:     true,
    }

    err = s.UserRepo.Save(ctx, user)
    if err != nil {
        return nil, err
    }
    return user, nil
}

// Login Authenticate a user
func (s *UserService) Login(ctx context.Context, email string, password string) (string, error) {
    user, err := s.UserRepo.FindByEmail(ctx, email)
    if err != nil {
        return "", errors.New("invalid email or password")
    }
    if user.IsDeleted() || !user.IsActive {
        return "", errors.New("user not active")
    }
    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
        return "", errors.New("invalid email or password")
    }

    // Only admin role for now
    token, err := s.JWTService.GenerateToken(user.ID.String(), "admin")
    if err != nil {
        return "", err
    }
    return token, nil
}

// ListUsers List all users
func (s *UserService) ListUsers(ctx context.Context) ([]*entity.User, error) {
    return s.UserRepo.ListAll(ctx)
}

// DeactivateUser Deactivate user (soft delete)
func (s *UserService) DeactivateUser(ctx context.Context, id string) error {
    return s.UserRepo.SoftDelete(ctx, id)
}
