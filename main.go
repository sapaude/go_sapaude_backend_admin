package main

import (
    "context"

    "github.com/sapaude/go-shims/x/log"
    "github.com/sapaude/go_sapaude_backend_admin/conf"
    "github.com/sapaude/go_sapaude_backend_admin/routes"
    "github.com/sirupsen/logrus"
)

func main() {
    logrus.SetFormatter(&logrus.JSONFormatter{
        TimestampFormat:   `2006-01-02 15:04:05.000`,
        DisableTimestamp:  false,
        DisableHTMLEscape: false,
        DataKey:           "",
        FieldMap:          nil,
        CallerPrettyfier:  nil,
        PrettyPrint:       false,
    })
    background := context.Background()
    ctx := context.WithValue(background, "key1", "val1")
    ctx = context.WithValue(ctx, "uin", 100)
    srvLog := logrus.WithContext(ctx).WithFields(logrus.Fields{
        "modName": "fieldMod",
        "key1":    ctx.Value("key1"),
        "uin":     ctx.Value("uin"),
    })
    // log.SetDefaultLogger(srvLog.Logger)
    if err := conf.LoadConfig("conf/app.yml", "conf/route.yml"); err != nil {
        logrus.Fatalf("config load error: %v", err)
    }
    srvLog.Infof("info what ")

    // 创建Echo服务
    echoServer := routes.NewEchoServer(InitBackendAdminImpl())

    // 启动服务
    addr := conf.AppConfig.Server.Address
    log.Infof("Starting server at %s", addr)
    if err := echoServer.Start(addr); err != nil {
        log.Fatalf("server error: %v", err)
    }
}
