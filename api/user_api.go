package api

import (
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
    "github.com/sapaude/go_sapaude_backend_admin/application"
    "github.com/sapaude/go_sapaude_backend_admin/domain/entity"
)

type UserAPI struct {
    userApp    *application.UserApp
    crontabApp *application.CrontabTaskApp
}

func NewUserAPI(userApp *application.UserApp, crontabApp *application.CrontabTaskApp) *UserAPI {
    return &UserAPI{userApp: userApp, crontabApp: crontabApp}
}

// Login POST /login
func (h *UserAPI) Login(c echo.Context) error {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
    }

    token, err := h.userApp.Login(c.Request().Context(), req.Email, req.Password)
    if err != nil {
        return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
    }

    return c.JSON(http.StatusOK, map[string]string{"token": token})
}

// CreateUser POST /users
func (h *UserAPI) CreateUser(c echo.Context) error {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
    }

    user, err := h.userApp.CreateUser(c.Request().Context(), req.Email, req.Password)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusCreated, user)
}

// ListUsers GET /users
func (h *UserAPI) ListUsers(c echo.Context) error {
    users, err := h.userApp.ListUsers(c.Request().Context(), &entity.PageSetting{
        PageNum:  0,
        PageSize: 10,
    })
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, users)
}

// DeactivateUser DELETE /users/:id
func (h *UserAPI) DeactivateUser(c echo.Context) error {
    userId := c.Param("uid")
    uid, _ := strconv.ParseUint(userId, 10, 64)

    if err := h.userApp.DeactivateUser(c.Request().Context(), uid); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return c.NoContent(http.StatusNoContent)
}
