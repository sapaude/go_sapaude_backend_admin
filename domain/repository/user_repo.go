package repository

import (
    "context"

    "github.com/sapaude/go_sapaude_backend_admin/domain/entity"
)

// IReposUserMode 用户仓储接口
type IReposUserMode interface {
    FindByEmail(ctx context.Context, email string) (*entity.User, error)
    FindByID(ctx context.Context, id uint64) (*entity.User, error)
    Save(ctx context.Context, user *entity.User) error
    SoftDelete(ctx context.Context, id uint64) error
    ListPageUsers(ctx context.Context, pg *entity.PageSetting) ([]*entity.User, error)
}
