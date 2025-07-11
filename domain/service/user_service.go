package service

import (
    "context"
    "errors"
    "strconv"

    "github.com/sapaude/go_sapaude_backend_admin/domain/entity"
    "github.com/sapaude/go_sapaude_backend_admin/domain/repository"
    "golang.org/x/crypto/bcrypt"
)

// IServiceUser 用户服务接口
type IServiceUser interface {
    // CreateUser 新增后台用户
    CreateUser(ctx context.Context, email string, rawPassword string) (*entity.User, error)

    // Login 后台用户登录
    Login(ctx context.Context, email string, password string) (string, error)

    // ListUsers 罗列后台用户
    ListUsers(ctx context.Context, pg *entity.PageSetting) ([]*entity.User, error)

    // DeactivateUser 注销后台用户
    DeactivateUser(ctx context.Context, uid uint64) error
}

// UserService 后台用户服务
type UserService struct {
    UserRepo   repository.IReposUserMode
    JWTService repository.IReposAUTH
}

func NewUserService(repo repository.IReposUserMode, jwt repository.IReposAUTH) *UserService {
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
    if !user.IsActive {
        return "", errors.New("user not active")
    }
    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
        return "", errors.New("invalid email or password")
    }

    // Only admin role for now
    token, err := s.JWTService.GenerateToken(strconv.FormatUint(user.UID, 10), "admin")
    if err != nil {
        return "", err
    }
    return token, nil
}

// ListUsers List all users
func (s *UserService) ListUsers(ctx context.Context, pg *entity.PageSetting) ([]*entity.User, error) {
    return s.UserRepo.ListPageUsers(ctx, pg)
}

// DeactivateUser Deactivate user (soft delete)
func (s *UserService) DeactivateUser(ctx context.Context, uid uint64) error {
    return s.UserRepo.SoftDelete(ctx, uid)
}
