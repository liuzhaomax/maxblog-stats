package set

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-stats/src/api_stats_rpc/business"
)

var BusinessSet = wire.NewSet(
	business.BusinessStatsArticleSet,
)
