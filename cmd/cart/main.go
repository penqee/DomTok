/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"

	"github.com/west2-online/DomTok/app/cart"
	"github.com/west2-online/DomTok/config"
	"github.com/west2-online/DomTok/kitex_gen/cart/cartservice"
	"github.com/west2-online/DomTok/pkg/base"
	"github.com/west2-online/DomTok/pkg/constants"
	"github.com/west2-online/DomTok/pkg/logger"
	"github.com/west2-online/DomTok/pkg/middleware"
	"github.com/west2-online/DomTok/pkg/utils"
)

var serviceName = constants.CartServiceName

func init() {
	config.Init(serviceName)
	logger.Init(serviceName, config.GetLoggerLevel())
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{config.Etcd.Addr})
	if err != nil {
		logger.Fatalf("Cart: new etcd registry failed, err: %v", err)
	}
	listenAddr, err := utils.GetAvailablePort()
	if err != nil {
		logger.Fatalf("Cart: get available port failed, err: %v", err)
	}
	addr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		logger.Fatalf("Cart: resolve tcp addr failed, err: %v", err)
	}

	p := base.TelemetryProvider(serviceName, config.Otel.CollectorAddr)
	defer func() { logger.LogError(p.Shutdown(context.Background())) }()

	svr := cartservice.NewServer(
		// 注入 controller 依赖
		cart.InjectCartHandler(),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: serviceName,
		}),
		server.WithMuxTransport(),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{
			MaxConnections: constants.MaxConnections,
			MaxQPS:         constants.MaxQPS,
		}),

		server.WithMiddleware(middleware.ErrorLog()),
		server.WithMiddleware(middleware.Respond()),
	)
	if err = svr.Run(); err != nil {
		logger.Fatalf("Cart: run server failed, err: %v", err)
	}
}
