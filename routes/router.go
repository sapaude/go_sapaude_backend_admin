package routes

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/sapaude/go_sapaude_backend_admin/api"
    "github.com/sapaude/go_sapaude_backend_admin/conf"
)

type Router struct {
    UserHandler *api.UserAPI
}

func NewRouter(uh *api.UserAPI) *Router {
    return &Router{UserHandler: uh}
}

func (r *Router) InitRoutes(e *echo.Echo) {
    // Middlewares
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Public
    for _, route := range conf.AppConfig.Route.PublicRoutes {
        switch route {
        case "/login":
            e.POST("/login", r.UserHandler.Login)
        }
    }

    // auth := e.Group("/users", middleware.JWTWithConfig(middleware.JWTConfig{
    //     SigningKey:  []byte(conf.AppConfig.JWT.Secret),
    //     TokenLookup: "header:Authorization",
    //     AuthScheme:  "Bearer",
    // }))
    auth := e.Group("/users")
    auth.POST("", r.UserHandler.CreateUser)
    auth.GET("", r.UserHandler.ListUsers)
    auth.DELETE("/:id", r.UserHandler.DeactivateUser)
}

// NewEchoServer Provide Echo instance with route registration
func NewEchoServer(uh *api.UserAPI) *echo.Echo {
    e := echo.New()
    r := NewRouter(uh)
    r.InitRoutes(e)
    return e
}
