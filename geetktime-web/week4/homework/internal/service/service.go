
package service

import (
	"geetktime-web/week4/homework/api"
	"geetktime-web/week4/homework/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewSceneService)

type SceneService struct {
	api.UnimplementedSceneServer

	cc  *biz.SceneUseCase
	log *log.Helper
}

func NewSceneService(cc *biz.SceneUseCase, logger log.Logger) *SceneService {
	return &SceneService{
		cc:  cc,
		log: log.NewHelper(log.With(logger, "module", "service/cart"))}
}
