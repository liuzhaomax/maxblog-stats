package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	businessRpc "github.com/liuzhaomax/maxblog-stats/src/api_stats_rpc/business"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	RPCService         *businessRpc.BusinessStatsArticle
	Engine             *gin.Engine
	DB                 *gorm.DB
	Redis              *redis.Client
	PrometheusRegistry *prometheus.Registry
}
