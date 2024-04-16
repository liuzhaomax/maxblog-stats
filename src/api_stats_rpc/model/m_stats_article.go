package model

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ModelStatsArticleSet = wire.NewSet(wire.Struct(new(ModelStatsArticle), "*"))

type ModelStatsArticle struct {
	DB *gorm.DB
}
