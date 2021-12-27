package homework

import (
	"geetktime-web/week4/homework/internal/biz"
	"geetktime-web/week4/homework/internal/conf"
	"geetktime-web/week4/homework/internal/data"
	"geetktime-web/week4/homework/internal/server"
	"geetktime-web/week4/homework/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	database := data.NewMongo(confData)
	dataData, cleanup, err := data.NewData(database, logger)
	if err != nil {
		return nil, nil, err
	}
	sceneRepo := data.NewSceneRepo(dataData, logger)
	cartUseCase := biz.NewSceneUseCase(sceneRepo, logger)
	cartService := service.NewSceneService(cartUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, tracerProvider, cartService)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}