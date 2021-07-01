// +build wireinject

package main

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"aurora/blog/api/app/service/manager/internal/conf"
	"aurora/blog/api/app/service/manager/internal/data"
	"aurora/blog/api/app/service/manager/internal/server"
	"aurora/blog/api/app/service/manager/internal/service"
	"aurora/blog/api/pkg/lifecycle"
	"github.com/google/wire"
)

func initApp() (*lifecycle.App, func(), error) {
	panic(wire.Build(server.ProvideSet, service.ProvideSet, biz.ProvideSet, data.ProvideSet, conf.ProvideSet, newApp))
}
