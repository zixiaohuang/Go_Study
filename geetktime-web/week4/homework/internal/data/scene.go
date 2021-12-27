package data

import (
	"context"
	"geetktime-web/week4/homework/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type sceneRepo struct {
	data *Data
	sceneColl *mongo.Collection
	log *log.Helper
}

func NewSceneRepo(data *Data, logger log.Logger) biz.SceneRepo {
	return &sceneRepo{
		data: data,
		sceneColl: data.db.Collection("scene"),
		log: log.NewHelper(log.With(logger, "module", "repo/scene")),
	}
}

type Scene struct {
	UserId int64 `bson:"user_id"`
	Items  []struct {
		ItemId   int64 `bson:"role_id"`
		Blood float64 `bson:"blood_val"`
	} `bson:"items"`
}

func (r *sceneRepo)GetScene(ctx context.Context, uid int64)(*biz.Scene, error) {
	result := &Scene{}
	if err := r.sceneColl.FindOne(ctx, bson.M{"s": uid}).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return &biz.Scene{SceneId: result.UserId}, nil
		}
		return nil, err
	}
	items := make([]biz.Role, 0)
	for _, x := range result.Items {
		items = append(items, biz.Role{
			Id:       x.ItemId,
			Blood: 	  x.Blood,
		})
	}
	return &biz.Scene{SceneId: result.UserId, Roles: items}, nil
}

func (r *sceneRepo) SaveScene(ctx context.Context, c *biz.Scene) error{
	// bson 一个bson数组
	items := bson.A{}
	for _, x := range c.Roles {
		items = append(items, bson.M{"item_id": x.Id, "blood_val": x.Blood})
	}
	result := r.sceneColl.FindOneAndUpdate(ctx, bson.M{"s": c.SceneId},
		// bson 文档
		bson.D{{"role_id", c.SceneId}, {"items", items}})
	return result.Err()
}

func (r *sceneRepo)DeleteScene(ctx context.Context, uid int64) error{
	// bson 无序的map
	_, err := r.sceneColl.DeleteOne(ctx, bson.M{"s": uid})
	return err
}
