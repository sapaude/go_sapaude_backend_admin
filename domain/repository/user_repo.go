package repository

import (
    "context"

    "github.com/lupguo/go_sapaude_backend_admin/domain/entity"
)

// UserRepository 用户仓储接口
type UserRepository interface {
    FindByEmail(ctx context.Context, email string) (*entity.User, error)
    FindByID(ctx context.Context, id string) (*entity.User, error)
    Save(ctx context.Context, user *entity.User) error
    SoftDelete(ctx context.Context, id string) error
    ListAll(ctx context.Context) ([]*entity.User, error)
}
