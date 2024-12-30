package svc

import (
	configurator "github.com/zeromicro/go-zero/core/configcenter"

	"github.com/jzero-io/jzero-admin-plugins/hello/internal/config"
	"github.com/jzero-io/jzero-admin-plugins/hello/internal/custom"
	"github.com/jzero-io/jzero-admin-plugins/hello/internal/middleware"
)

type ServiceContext struct {
	Config configurator.Configurator[config.Config]
	middleware.Middleware
	Custom *custom.Custom
}

func NewServiceContext(cc configurator.Configurator[config.Config]) *ServiceContext {
	sc := &ServiceContext{
		Config:     cc,
		Custom:     custom.New(),
		Middleware: middleware.New(),
	}
	sc.SetConfigListener()
	return sc
}
