// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/liuzhaomax/maxblog-stats/internal/api"
	"github.com/liuzhaomax/maxblog-stats/internal/core"
	"github.com/liuzhaomax/maxblog-stats/internal/middleware_rpc"
	"github.com/liuzhaomax/maxblog-stats/internal/middleware_rpc/auth"
	"github.com/liuzhaomax/maxblog-stats/internal/middleware_rpc/tracing"
	"github.com/liuzhaomax/maxblog-stats/internal/middleware_rpc/validator"
	"github.com/liuzhaomax/maxblog-stats/src/api_stats_rpc/business"
	"github.com/liuzhaomax/maxblog-stats/src/api_stats_rpc/model"
)

// Injectors from wire.go:

func InitInjector() (*Injector, func(), error) {
	registry := core.InitPrometheusRegistry()
	logger := core.InitLogrus()
	coreLogger := &core.Logger{
		Logger: logger,
	}
	client, cleanup, err := core.InitRedis()
	if err != nil {
		return nil, nil, err
	}
	authRPC := &auth.AuthRPC{
		Logger: coreLogger,
		Redis:  client,
	}
	validatorRPC := &validator.ValidatorRPC{
		Logger: coreLogger,
		Redis:  client,
	}
	configuration := core.InitTracer()
	tracingRPC := &tracing.TracingRPC{
		Logger:       coreLogger,
		TracerConfig: configuration,
	}
	middlewareRPC := &middleware_rpc.MiddlewareRPC{
		AuthRPC:      authRPC,
		ValidatorRPC: validatorRPC,
		TracingRPC:   tracingRPC,
	}
	db, cleanup2, err := core.InitDB()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	modelStatsArticle := &model.ModelStatsArticle{
		DB: db,
	}
	trans := &core.Trans{
		DB: db,
	}
	response := &core.Response{
		Logger: coreLogger,
	}
	rocketMQ := &core.RocketMQ{}
	businessStatsArticle := &business.BusinessStatsArticle{
		Model:    modelStatsArticle,
		Tx:       trans,
		Redis:    client,
		IRes:     response,
		RocketMQ: rocketMQ,
	}
	handlerRPC := &api.HandlerRPC{
		PrometheusRegistry: registry,
		MiddlewareRPC:      middlewareRPC,
		BusinessRPC:        businessStatsArticle,
	}
	injector := &Injector{
		HandlerRPC: handlerRPC,
		DB:         db,
		Redis:      client,
	}
	return injector, func() {
		cleanup2()
		cleanup()
	}, nil
}
