package db

import (
    "github.com/lupguo/go_sapaude_backend_admin/conf"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type BackendDB struct {
    db *gorm.DB
}

func NewBackendDB() *BackendDB {
    dbConn, err := gorm.Open(mysql.Open(conf.AppConfig.Database.DSN), &gorm.Config{})
    if err != nil {
        log.Fatalf("db init error: %v", err)
    }

    return &BackendDB{db: db}
}
