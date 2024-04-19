package model

import (
	"context"
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-stats/src/api_stats_rpc/pb"
	"gorm.io/gorm"
)

var ModelStatsArticleSet = wire.NewSet(wire.Struct(new(ModelStatsArticle), "*"))

type ModelStatsArticle struct {
	DB *gorm.DB
}

func (m *ModelStatsArticle) GetStatsArticleMain(ctx context.Context, data *pb.StatsArticleMainRes) error {
	var view, like, quantity int64
	viewResult := m.DB.WithContext(ctx).Table("article").Select("SUM(view)").Scan(&view)
	if viewResult.RowsAffected == 0 {
		return viewResult.Error
	}
	likeResult := m.DB.WithContext(ctx).Table("article").Select("SUM(`like`)").Scan(&like)
	if likeResult.RowsAffected == 0 {
		return likeResult.Error
	}
	quantityResult := m.DB.WithContext(ctx).Table("article").Count(&quantity)
	if quantityResult.RowsAffected == 0 {
		return quantityResult.Error
	}
	data.Quantity = int32(quantity)
	data.View = int32(view)
	data.Like = int32(like)
	return nil
}
