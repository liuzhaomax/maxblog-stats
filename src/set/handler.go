package set

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-stats/src/api_user/handler"
)

var HandlerSet = wire.NewSet(
	handler.HandlerUserSet,
)
