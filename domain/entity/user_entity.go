package entity

import (
    "time"

    "gorm.io/gorm"
)

// User 用户实体
type User struct {
    ID           uint64 `gorm:"primaryKey"`
    UID          uint64 `gorm:"index"`
    Email        string `gorm:"uniqueIndex;not null"`
    PasswordHash string
    IsActive     bool
    CreatedAt    time.Time
    UpdatedAt    time.Time
    DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (u *User) TableName() string {
    return "t_users"
}
