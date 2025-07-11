package dbs

import (
    "github.com/sapaude/go-shims/x/log"
    "github.com/sapaude/go_sapaude_backend_admin/conf"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

// AdminDB 后台数据库
type AdminDB struct {
    db *gorm.DB
}

func NewAdminDB() *AdminDB {
    db, err := gorm.Open(mysql.Open(conf.AppConfig.Database.DSN), &gorm.Config{})
    if err != nil {
        log.Fatalf("db init error: %v", err)
    }

    return &AdminDB{db: db}
}
