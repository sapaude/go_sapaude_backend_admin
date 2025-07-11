package dbs

import (
    "context"

    "github.com/sapaude/go_sapaude_backend_admin/domain/entity"
    "gorm.io/gorm"
)

// UserModeInfra 用户Infra
type UserModeInfra struct {
    db *gorm.DB
}

// NewUserDBInfra 基于后台数据库初始化用户Infra
func NewUserDBInfra(adminDB *AdminDB) *UserModeInfra {
    return &UserModeInfra{adminDB.db}
}

func (r *UserModeInfra) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
    var user entity.User
    err := r.db.WithContext(ctx).
        Where("email = ? AND deleted_at IS NULL", email).
        First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserModeInfra) FindByID(ctx context.Context, id uint64) (*entity.User, error) {
    var user entity.User
    err := r.db.WithContext(ctx).
        Where("id = ? AND deleted_at IS NULL", id).
        First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserModeInfra) Save(ctx context.Context, user *entity.User) error {
    return r.db.WithContext(ctx).Save(user).Error
}

func (r *UserModeInfra) SoftDelete(ctx context.Context, id uint64) error {
    return r.db.WithContext(ctx).Delete(&entity.User{ID: id}).Error
}

// ListPageUsers 罗列指定分页中所有用户
func (r *UserModeInfra) ListPageUsers(ctx context.Context, pg *entity.PageSetting) ([]*entity.User, error) {
    var users []*entity.User
    err := r.db.WithContext(ctx).
        Where("deleted_at IS NULL").
        Find(&users).
        Limit(pg.PageSize).
        Offset((pg.PageNum - 1) * pg.PageSize).
        Error
    if err != nil {
        return nil, err
    }
    return users, nil
}
