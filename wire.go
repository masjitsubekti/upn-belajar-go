//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"gitlab.com/upn-belajar-go/configs"
	"gitlab.com/upn-belajar-go/infras"
	"gitlab.com/upn-belajar-go/internal/domain/master"
	"gitlab.com/upn-belajar-go/internal/domain/orm"
	"gitlab.com/upn-belajar-go/internal/handlers"
	"gitlab.com/upn-belajar-go/transport/http"
	"gitlab.com/upn-belajar-go/transport/http/middleware"
	"gitlab.com/upn-belajar-go/transport/http/router"
)

// Wiring for configurations.
var configurations = wire.NewSet(
	configs.Get,
)

// Wiring for persistences.
var persistences = wire.NewSet(
	infras.ProvidePostgreSQLConn,
)

// Wiring for domain Master
var domainMaster = wire.NewSet(
	// JenisMitraService interface and implementation
	master.ProvideJenisMitraServiceImpl,
	wire.Bind(new(master.JenisMitraService), new(*master.JenisMitraServiceImpl)),
	// JenisMitraRepository interface and implementation
	master.ProvideJenisMitraRepositoryPostgreSQL,
	wire.Bind(new(master.JenisMitraRepository), new(*master.JenisMitraRepositoryPostgreSQL)),

	// SiswaService interface and implementation
	master.ProvideSiswaServiceImpl,
	wire.Bind(new(master.SiswaService), new(*master.SiswaServiceImpl)),
	// SiswaRepository interface and implementation
	master.ProvideSiswaRepositoryPostgreSQL,
	wire.Bind(new(master.SiswaRepository), new(*master.SiswaRepositoryPostgreSQL)),

	// KelasService interface and implementation
	orm.ProvideKelasServiceImpl,
	wire.Bind(new(orm.KelasService), new(*orm.KelasServiceImpl)),
)

// Wiring for all domains.
var domains = wire.NewSet(
	domainMaster,
)

// Wiring for HTTP routing.
var routing = wire.NewSet(
	wire.Struct(new(router.DomainHandlers), "*"),
	handlers.ProvideJenisMitraHandler,
	handlers.ProvideSiswaHandler,
	handlers.ProvideKelasHandler,
	// jwt
	middleware.ProvideJWTMiddleware,
	router.ProvideRouter,
)

// Wiring for everything.
func InitializeService() *http.HTTP {
	wire.Build(
		// configurations
		configurations,
		// persistences
		persistences,
		// domains
		domains,
		// routing
		routing,
		// selected transport layer
		http.ProvideHTTP)
	return &http.HTTP{}
}
