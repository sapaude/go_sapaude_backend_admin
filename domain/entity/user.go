package entity

import (
    "time"

    "github.com/google/uuid"
)

type User struct {
    ID           uuid.UUID
    Email        string
    PasswordHash string
    IsActive     bool
    DeletedAt    *time.Time
    CreatedAt    time.Time
    UpdatedAt    time.Time
}

// Business logic example
func (u *User) IsDeleted() bool {
    return u.DeletedAt != nil
}
