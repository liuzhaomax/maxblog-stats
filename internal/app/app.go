package app

import (
	"context"
	"fmt"
	"github.com/liuzhaomax/maxblog-stats/internal/api"
	"github.com/liuzhaomax/maxblog-stats/internal/core"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Option func(*options)

type options struct {
	ConfigFile string
	WWWDir     string
}

func SetConfigFile(configFile string) Option {
	return func(opts *options) {
		opts.ConfigFile = configFile
	}
}

func SetWWWDir(wwwDir string) Option {
	return func(opts *options) {
		opts.WWWDir = wwwDir
	}
}

func InitConfig(opts *options) func() {
	cfg := core.GetConfig()
	cleanLogger := cfg.LoadConfig(opts.ConfigFile)
	cfg.App.Logger.WithField("path", opts.ConfigFile).Info(core.FormatInfo("配置文件加载成功"))
	cfg.App.Logger.Info(core.FormatInfo("系统启动"))
	if cfg.App.Enabled.ServiceDiscovery {
		// register service
		err := cfg.Lib.Consul.ServiceRegister()
		if err != nil {
			cfg.App.Logger.WithField(core.FAILURE, core.GetFuncName()).Fatal(core.FormatError(core.Unknown, "服务注册失败", err))
		}
		cfg.App.Logger.Info(core.FormatInfo("服务注册成功"))
		// discover services
		go func() {
			for {
				err = cfg.Lib.Consul.ServiceDiscover()
				if err != nil {
					cfg.App.Logger.WithField(core.FAILURE, core.GetFuncName()).Warn(core.FormatError(core.Unknown, "下游服务发现失败", err))
				}
				time.Sleep(time.Duration(cfg.Lib.Consul.Interval) * time.Second)
			}
		}()
	}
	return func() {
		cleanLogger()
	}
}

func InitRpcServer(ctx context.Context, handlerRPC *api.HandlerRPC) func() {
	cfg := core.GetConfig()
	cfg.App.Logger.Info(core.FormatInfo("服务启动开始"))
	addr := fmt.Sprintf("%s:%s", "0.0.0.0", cfg.Server.Port)
	server := handlerRPC.Register()
	go func() {
		listen, err := net.Listen("tcp", addr)
		if err != nil {
			cfg.App.Logger.WithField(core.FAILURE, core.GetFuncName()).Fatal(core.FormatError(core.Unknown, "服务监听失败", err))
		}
		cfg.App.Logger.WithContext(ctx).Infof("Service %s is running at %s", cfg.App.Name, addr)
		err = server.Serve(listen)
		if err != nil {
			cfg.App.Logger.WithField(core.FAILURE, core.GetFuncName()).Fatal(core.FormatError(core.Unknown, "服务启动失败", err))
		}
	}()
	return func() {
		cfg.App.Logger.Info(core.FormatInfo("服务关闭开始"))
		_, cancel := context.WithTimeout(ctx, time.Second*time.Duration(cfg.Server.ShutdownTimeout))
		defer cancel()
		server.Stop()
		cfg.App.Logger.Info(core.FormatInfo("服务关闭成功"))
	}
}

func Init(ctx context.Context, optFuncs ...Option) func() {
	// initialising options
	opts := options{}
	for _, optFunc := range optFuncs {
		optFunc(&opts)
	}
	// init conf
	cleanConfig := InitConfig(&opts)
	cfg := core.GetConfig()
	// init injector
	injector, cleanInjection, _ := InitInjector()
	// init server by protocol
	cleanServer := InitRpcServer(ctx, injector.HandlerRPC)
	cfg.App.Logger.WithFields(logrus.Fields{
		"app_name": cfg.App.Name,
		"version":  cfg.App.Version,
		"pid":      os.Getpid(),
		"host":     cfg.Server.Host,
		"port":     cfg.Server.Port,
		"protocol": cfg.Server.Protocol,
	}).Info(core.FormatInfo("服务启动成功"))
	return func() {
		cleanServer()
		cleanInjection()
		cleanConfig()
	}
}

func Launch(ctx context.Context, opts ...Option) {
	cfg := core.GetConfig()
	clean := Init(ctx, opts...)
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
EXIT:
	for {
		sig := <-sc
		cfg.App.Logger.WithContext(ctx).Infof("%s [%s]", core.FormatInfo("服务中断信号收到"), sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}
	defer time.Sleep(time.Second)
	defer os.Exit(state)
	defer clean()
	defer cfg.App.Logger.WithContext(ctx).Infof(core.FormatInfo("系统停止"))
}
