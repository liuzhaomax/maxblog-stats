package set

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-stats/src/api_user/business"
	businessRpc "github.com/liuzhaomax/maxblog-stats/src/api_user_rpc/business"
)

var BusinessSet = wire.NewSet(
	business.BusinessUserSet,
	businessRpc.BusinessUserSet,
)
