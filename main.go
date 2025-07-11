package main

import (
    "strconv"

    "github.com/labstack/echo/v4"
    "github.com/lupguo/go_sapaude_backend_admin/conf"
    "github.com/lupguo/go_sapaude_backend_admin/infra/logger"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func main() {
    log.InitLogger()
    log := log.Log

    if err := conf.LoadConfig("conf/app.yml", "conf/route.yml"); err != nil {
        log.Fatalf("config load error: %v", err)
    }

    var app *echo.Echo
    app, err := di.InitApp(conf.AppConfig, dbConn)
    if err != nil {
        log.Fatalf("wire init error: %v", err)
    }

    addr := ":" + strconv.Itoa(conf.AppConfig.Server.Port)
    log.Infof("Starting server at %s", addr)
    if err := app.Start(addr); err != nil {
        log.Fatalf("server error: %v", err)
    }
}
