package db

import (
    "context"
    "time"

    "github.com/google/uuid"
    "github.com/lupguo/go_sapaude_backend_admin/domain/entity"
    "github.com/lupguo/go_sapaude_backend_admin/domain/repository"
    "gorm.io/gorm"
)

type gormUserModel struct {
    ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
    Email        string    `gorm:"uniqueIndex;not null"`
    PasswordHash string    `gorm:"not null"`
    IsActive     bool
    DeletedAt    *time.Time
    CreatedAt    time.Time
    UpdatedAt    time.Time
}

// Convert between GORM model and domain entity
func toDomain(m *gormUserModel) *entity.User {
    return &entity.User{
        ID:           m.ID,
        Email:        m.Email,
        PasswordHash: m.PasswordHash,
        IsActive:     m.IsActive,
        DeletedAt:    m.DeletedAt,
        CreatedAt:    m.CreatedAt,
        UpdatedAt:    m.UpdatedAt,
    }
}

func fromDomain(e *entity.User) *gormUserModel {
    return &gormUserModel{
        ID:           e.ID,
        Email:        e.Email,
        PasswordHash: e.PasswordHash,
        IsActive:     e.IsActive,
        DeletedAt:    e.DeletedAt,
        CreatedAt:    e.CreatedAt,
        UpdatedAt:    e.UpdatedAt,
    }
}

type userRepo struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
    return &userRepo{db}
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
    var model gormUserModel
    err := r.db.WithContext(ctx).Where("email = ? AND deleted_at IS NULL", email).First(&model).Error
    if err != nil {
        return nil, err
    }
    return toDomain(&model), nil
}

func (r *userRepo) FindByID(ctx context.Context, id string) (*entity.User, error) {
    var model gormUserModel
    err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&model).Error
    if err != nil {
        return nil, err
    }
    return toDomain(&model), nil
}

func (r *userRepo) Save(ctx context.Context, user *entity.User) error {
    model := fromDomain(user)
    return r.db.WithContext(ctx).Save(model).Error
}

func (r *userRepo) SoftDelete(ctx context.Context, id string) error {
    return r.db.WithContext(ctx).Model(&gormUserModel{}).
        Where("id = ?", id).
        Update("deleted_at", time.Now()).Error
}

func (r *userRepo) ListAll(ctx context.Context) ([]*entity.User, error) {
    var models []gormUserModel
    err := r.db.WithContext(ctx).Where("deleted_at IS NULL").Find(&models).Error
    if err != nil {
        return nil, err
    }
    users := make([]*entity.User, len(models))
    for i, m := range models {
        users[i] = toDomain(&m)
    }
    return users, nil
}
