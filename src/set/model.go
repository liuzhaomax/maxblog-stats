package set

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-stats/src/api_stats_rpc/model"
)

var ModelSet = wire.NewSet(
	model.ModelStatsArticleSet,
)
