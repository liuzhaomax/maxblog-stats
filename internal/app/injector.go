package app

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-stats/internal/api"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	HandlerRPC *api.HandlerRPC
	DB         *gorm.DB
	Redis      *redis.Client
}
