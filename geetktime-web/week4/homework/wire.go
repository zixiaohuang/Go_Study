package homework

import (
	"geetktime-web/week4/homework/internal/biz"
	"geetktime-web/week4/homework/internal/data"
	"github.com/google/wire"
)

func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger, *tracesdk.TracerProvider) (*App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}