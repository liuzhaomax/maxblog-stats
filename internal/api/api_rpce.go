package api

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-stats/internal/core"
	"github.com/liuzhaomax/maxblog-stats/internal/middleware_rpc"
	"github.com/liuzhaomax/maxblog-stats/src/api_stats_rpc/business"
	"github.com/liuzhaomax/maxblog-stats/src/api_stats_rpc/pb"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

var APIRPCSet = wire.NewSet(wire.Struct(new(HandlerRPC), "*"), wire.Bind(new(APIRPC), new(*HandlerRPC)))

type APIRPC interface {
	Register() *grpc.Server
}

type HandlerRPC struct {
	PrometheusRegistry *prometheus.Registry
	MiddlewareRPC      *middleware_rpc.MiddlewareRPC
	BusinessRPC        *business.BusinessStatsArticle
}

func (h *HandlerRPC) Register() *grpc.Server {
	interceptorsBasicChoice := []grpc.UnaryServerInterceptor{
		core.LoggerForRPC,
		h.MiddlewareRPC.TracingRPC.Trace,
		h.MiddlewareRPC.ValidatorRPC.ValidateMetadata,
		h.MiddlewareRPC.AuthRPC.ValidateSignature,
	}
	// TODO prometheus metrics 接口
	interceptorMap := map[string][]grpc.UnaryServerInterceptor{
		"/StatsService/GetStatsArticleMain": interceptorsBasicChoice,
	}

	// 连接多个中间件
	serverOpts := []grpc.ServerOption{}
	serverOpts = append(serverOpts, grpc.UnaryInterceptor(middleware_rpc.ChainUnaryInterceptors(interceptorMap)))

	// 创建gRPC服务
	server := grpc.NewServer(serverOpts...)
	// 注册接口
	pb.RegisterStatsServiceServer(server, h.BusinessRPC)

	return server
}
