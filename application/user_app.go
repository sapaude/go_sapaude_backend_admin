package application

import (
    "context"

    "github.com/sapaude/go_sapaude_backend_admin/domain/entity"
    "github.com/sapaude/go_sapaude_backend_admin/domain/service"
)

// UserApp 后台用户管理
type UserApp struct {
    userSrv service.IServiceUser
}

func NewUserApp(userSrv service.IServiceUser) *UserApp {
    return &UserApp{userSrv: userSrv}
}

// Login 登录
func (u *UserApp) Login(ctx context.Context, email, pass string) (token string, err error) {
    return u.userSrv.Login(ctx, email, pass)
}

// CreateUser 创建用户
func (u *UserApp) CreateUser(ctx context.Context, email, pass string) (token string, err error) {
    return "", nil
}

func (u *UserApp) Register(ctx context.Context, email, pass string) (token string, err error) {

    return "", nil
}

func (u *UserApp) Logout(ctx context.Context, token string) error {
    return nil
}

// GetUser 获取用户信息
func (u *UserApp) GetUser(ctx context.Context, token string) (*entity.User, error) {
    return nil, nil
}

func (u *UserApp) ListUsers(ctx context.Context, pg *entity.PageSetting) ([]*entity.User, error) {
    return u.userSrv.ListUsers(ctx, pg)
}

func (u *UserApp) DeactivateUser(ctx context.Context, uid uint64) error {
    return u.userSrv.DeactivateUser(ctx, uid)
}
