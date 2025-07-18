// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/sapaude/go_sapaude_backend_admin/api"
	"github.com/sapaude/go_sapaude_backend_admin/application"
	"github.com/sapaude/go_sapaude_backend_admin/domain/repository"
	"github.com/sapaude/go_sapaude_backend_admin/domain/service"
	"github.com/sapaude/go_sapaude_backend_admin/infra/auth"
	"github.com/sapaude/go_sapaude_backend_admin/infra/dbs"
)

// Injectors from wire.go:

func InitBackendAdminImpl() *api.UserAPI {
	adminDB := dbs.NewAdminDB()
	userModeInfra := dbs.NewUserDBInfra(adminDB)
	jwtInfra := auth.NewJWTInfra()
	userService := service.NewUserService(userModeInfra, jwtInfra)
	userApp := application.NewUserApp(userService)
	crontabTaskApp := application.NewCrontabTaskApp()
	userAPI := api.NewUserAPI(userApp, crontabTaskApp)
	return userAPI
}

// wire.go:

// api
var apiSet = wire.NewSet(api.NewUserAPI)

// app
var appSet = wire.NewSet(application.NewCrontabTaskApp, application.NewUserApp)

// srv
var srvSet = wire.NewSet(wire.Bind(new(service.IServiceUser), new(*service.UserService)), service.NewUserService)

// infra
var infraSet = wire.NewSet(wire.Bind(new(repository.IReposAUTH), new(*auth.JWTInfra)), auth.NewJWTInfra, wire.Bind(new(repository.IReposUserMode), new(*dbs.UserModeInfra)), dbs.NewUserDBInfra)

// base
var baseSet = wire.NewSet(dbs.NewAdminDB)
