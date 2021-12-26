package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Role struct {
	Id int64
	Blood float64
}

type Scene struct {
	SceneId int64
	Roles []Role
}

type SceneRepo interface {
	GetScene(ctx context.Context, uid int64)(*Scene, error)
	SaveScene(ctx context.Context, c *Role) error
	DeleteScene(ctx context.Context, uid int64) error
}

type SceneUseCase struct {
	repo SceneRepo
	log *log.Helper
}

func NewSceneUseCase(repo SceneRepo, logger log.Logger) *SceneUseCase {
	return &SceneUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/scene"))}
}

