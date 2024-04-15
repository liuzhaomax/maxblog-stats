//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-stats/internal/api"
	"github.com/liuzhaomax/maxblog-stats/internal/core"
	"github.com/liuzhaomax/maxblog-stats/internal/middleware"
	"github.com/liuzhaomax/maxblog-stats/src/set"
)

func InitInjector() (*Injector, func(), error) {
	wire.Build(
		core.InitLogrus,
		core.InitGinEngine,
		core.InitDB,
		core.InitRedis,
		core.InitTracer,
		core.InitPrometheusRegistry,
		api.APISet,
		set.HandlerSet,
		set.BusinessSet,
		set.ModelSet,
		core.LoggerSet,
		core.ResponseSet,
		core.TransactionSet,
		core.RocketMQSet,
		middleware.MwsSet,
		middleware.MiddlewareSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
