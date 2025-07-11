//go:build wireinject

package di

import (
    "github.com/google/wire"
    "github.com/labstack/echo/v4"
    "github.com/lupguo/go_sapaude_backend_admin/conf"
    "github.com/lupguo/go_sapaude_backend_admin/infra/auth"
    "github.com/lupguo/go_sapaude_backend_admin/infra/db"
    "gorm.io/gorm"
)

func InitApp(conf *conf.Config, dbConn *gorm.DB) (*echo.Echo, error) {
    wire.Build(
        // Repos
        db.NewUserRepository,

        // Services
        auth.NewJWTService,
        application.NewUserService,

        // Handlers
        api.NewUserHandler,
        routes.NewRouter,

        // Echo instance
        newEchoServer,
    )
    return &echo.Echo{}, nil
}

// Provide Echo instance with route registration
func newEchoServer(r *routes.Router) *echo.Echo {
    e := echo.New()
    r.InitRoutes(e)
    return e
}
