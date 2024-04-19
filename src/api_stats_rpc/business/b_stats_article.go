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
	// TODO prometheus
	data := &pb.StatsArticleMainRes{}
	err := b.Tx.ExecTrans(ctx, func(ctx context.Context) error {
		err := b.Model.GetStatsArticleMain(ctx, data)
		if err != nil {
			b.IRes.ResFailureForRPC(ctx, core.InternalServerError, "DB查询文章主要统计失败", err)
			return core.FormatError(core.DBDenied, "DB查询文章主要统计失败", err)
		}
		return nil
	})
	if err != nil {
		b.IRes.ResFailureForRPC(ctx, core.InternalServerError, "DB事务失败", err)
		return nil, core.FormatError(core.DBDenied, "DB事务失败", err)
	}
	b.IRes.ResSuccessForRPC(ctx)
	return data, nil
}
