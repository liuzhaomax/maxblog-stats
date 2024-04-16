package business

import (
	"context"
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-stats/internal/core"
	"github.com/liuzhaomax/maxblog-stats/src/api_stats_rpc/model"
	"github.com/liuzhaomax/maxblog-stats/src/api_stats_rpc/pb"
	"github.com/redis/go-redis/v9"
)

var BusinessStatsArticleSet = wire.NewSet(wire.Struct(new(BusinessStatsArticle), "*"))

type BusinessStatsArticle struct {
	Model    *model.ModelStatsArticle
	Tx       *core.Trans
	Redis    *redis.Client
	IRes     core.IResponse
	RocketMQ core.IRocketMQ
}

func (b *BusinessStatsArticle) GetStatsArticleMain(ctx context.Context, empty *pb.Empty) (*pb.StatsArticleMainRes, error) {
	// TODO implement me
	panic("implement me")
}
