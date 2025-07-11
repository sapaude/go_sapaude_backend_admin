package main

import (
    "strconv"

    "github.com/sapaude/go-shims/x/log"
    "github.com/sapaude/go_sapaude_backend_admin/conf"
    "github.com/sapaude/go_sapaude_backend_admin/routes"
)

func main() {
    if err := conf.LoadConfig("conf/app.yml", "conf/route.yml"); err != nil {
        log.Fatalf("config load error: %v", err)
    }

    // 创建Echo服务
    server := routes.NewEchoServer(InitBackendAdminImpl())

    // 启动服务
    addr := ":" + strconv.Itoa(conf.AppConfig.Server.Port)
    log.Infof("Starting server at %s", addr)
    if err := server.Start(addr); err != nil {
        log.Fatalf("server error: %v", err)
    }
}
