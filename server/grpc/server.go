// The MIT License (MIT)

// Copyright (c) 2022 leewckk@126.com

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package grpc

import (
	"micro-service-sample/configure"
	"micro-service-sample/transport/grpc/version"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/leewckk/go-kit-micro-service/middlewares/tracing/report"
	"github.com/leewckk/go-kit-micro-service/router/grpc"
)

func Run(cfg *configure.Config) chan error {

	trapi := cfg.Tracing.Reporter
	port := cfg.Server.Grpc.Port
	name := cfg.Discovery.ServiceName
	center := cfg.Discovery.Host
	centerPort := cfg.Discovery.Port

	opts := make([]kitgrpc.ServerOption, 0, 0)
	traceReporter := report.NewZipkinReporter(trapi)
	if nil != traceReporter {
		// defer traceReporter.Close()
		opts = append(opts, report.NewGrpcServerTracer(traceReporter))
	}

	procs := make([]grpc.RegisterRouteProc, 0, 0)
	/// TODO
	// 注册服务
	procs = append(procs, version.RegisterVersionServiceProc)

	router := grpc.NewRouter(name,
		grpc.RouterServerOptions(opts...), //// 插件
		grpc.RouterProcs(procs...),        //// 服务注册
	)

	return grpc.Run(router, port, name, center, centerPort)
}
