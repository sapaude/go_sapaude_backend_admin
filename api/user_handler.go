package api

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/lupguo/go_sapaude_backend_admin/application"
)

type UserHandler struct {
    UserService *application.UserService
}

func NewUserHandler(svc *application.UserService) *UserHandler {
    return &UserHandler{UserService: svc}
}

// Login POST /login
func (h *UserHandler) Login(c echo.Context) error {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
    }

    token, err := h.UserService.Login(c.Request().Context(), req.Email, req.Password)
    if err != nil {
        return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
    }

    return c.JSON(http.StatusOK, map[string]string{"token": token})
}

// CreateUser POST /users
func (h *UserHandler) CreateUser(c echo.Context) error {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
    }

    user, err := h.UserService.CreateUser(c.Request().Context(), req.Email, req.Password)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusCreated, user)
}

// ListUsers GET /users
func (h *UserHandler) ListUsers(c echo.Context) error {
    users, err := h.UserService.ListUsers(c.Request().Context())
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, users)
}

// DeactivateUser DELETE /users/:id
func (h *UserHandler) DeactivateUser(c echo.Context) error {
    id := c.Param("id")
    if err := h.UserService.DeactivateUser(c.Request().Context(), id); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return c.NoContent(http.StatusNoContent)
}
