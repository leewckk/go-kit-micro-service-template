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

package http

import (
	"fmt"
	"micro-service-sample/configure"
	"micro-service-sample/transport/gin/version"

	"github.com/leewckk/go-kit-micro-service/middlewares/tracing/report"
	transhttp "github.com/leewckk/go-kit-micro-service/middlewares/transport/http"
	"github.com/leewckk/go-kit-micro-service/router/http"
)

func Run(cfg *configure.Config) chan error {

	trapi := cfg.Tracing.Reporter
	port := cfg.Server.Http.Port
	name := cfg.Discovery.ServiceName
	center := cfg.Discovery.Host
	centerPort := cfg.Discovery.Port

	opts := make([]transhttp.ServerOption, 0, 0)
	traceReporter := report.NewZipkinReporter(trapi)
	if nil != traceReporter {
		// defer traceReporter.Close()
		opts = append(opts, report.NewGinServerTracer(traceReporter))
	}
	procs := make([]http.RegisterRouteProc, 0, 0)
	/// TODO
	/// 添加服务
	procs = append(procs, version.RegisterVersionServiceProc)

	router := http.NewRouter(fmt.Sprintf(":%v", port),
		http.RouterVersion(""),
		http.RouterServerOptions(opts...), //// 注册一些中间件
		http.RouterProcs(procs...),        //// 注册服务
		http.RouterHealth(http.DEFAULT_HTTP_HEART_BEAT, "/info", version.NewVersionServiceTransport()[0]),
	)
	return http.Run(router, port, name, center, centerPort)
}
