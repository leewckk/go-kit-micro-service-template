// This code is genetated by protoc-gen-gokit-micro.

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

// PLUGIN: protoc-gen-gokit-micro
//		VERSION : v0.0.2-alpha
//		GIT TAG : v0.0.2-alpha
//		GIT HASH : c45d18beb986d2e8a75ce66dfa3ff17333d184e6
//		BUILDER VERSION : go version go1.16.5 linux/amd64
//		BUILD TIME: : 2022-08-12 14:19:55

// create time : 2022-08-12 15:15:43.114259796 +0800 CST m=+0.006993360

package version

import (
	context "context"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	middlewares "github.com/leewckk/go-kit-micro-service/middlewares"
	reflect "reflect"
)

/// PROTOCOL DEFINITION
type VersionRequest struct {
}

/// PROTOCOL DEFINITION
type VersionInfo struct {
	Hash           string `json:"hash" uri:"hash" form:"hash"`                               /// GIT HASH信息
	Tag            string `json:"tag" uri:"tag" form:"tag"`                                  /// git TAG信息
	Version        string `json:"version" uri:"version" form:"version"`                      /// 版本信息
	BuilderVersion string `json:"builderVersion" uri:"builderVersion" form:"builderVersion"` /// 编译器版本号
	Uptime         string `json:"uptime" uri:"uptime" form:"uptime"`                         /// 启动时间
	RunningTimes   string `json:"runningTimes" uri:"runningTimes" form:"runningTimes"`       /// 运行时间
	BuildTime      string `json:"buildTime" uri:"buildTime" form:"buildTime"`                /// 编译时间
	HostName       string `json:"hostName" uri:"hostName" form:"hostName"`                   /// HOSTNAME
	Platform       string `json:"platform" uri:"platform" form:"platform"`                   /// PLATFORM
}

/// PROTOCOL DEFINITION
type VersionResponse struct {
	Hash           string `json:"hash" uri:"hash" form:"hash"`                               /// GIT HASH信息
	Tag            string `json:"tag" uri:"tag" form:"tag"`                                  /// git TAG信息
	Version        string `json:"version" uri:"version" form:"version"`                      /// 版本信息
	BuilderVersion string `json:"builderVersion" uri:"builderVersion" form:"builderVersion"` /// 编译器版本号
	Uptime         string `json:"uptime" uri:"uptime" form:"uptime"`                         /// 启动时间
	RunningTimes   string `json:"runningTimes" uri:"runningTimes" form:"runningTimes"`       /// 运行时间
	BuildTime      string `json:"buildTime" uri:"buildTime" form:"buildTime"`                /// 编译时间
	HostName       string `json:"hostName" uri:"hostName" form:"hostName"`                   /// HOSTNAME
	Platform       string `json:"platform" uri:"platform" form:"platform"`                   /// PLATFORM
}

//// PROTOTYPE OF SERVICE: VersionService , METHOD: Get
type VersionServiceGet func(context.Context, *VersionRequest) (*VersionResponse, error)

var __VersionServiceGetProc__ VersionServiceGet

func ImplementVersionServiceGet(impl func() VersionServiceGet) {
	__VersionServiceGetProc__ = impl()
}

/// ENDPOINT MAKER

func MakeVersionServiceGetEndpoint() endpoint.Endpoint {
	return middlewares.InjectMiddleWares(func(ctx context.Context, request interface{}) (interface{}, error) {
		if r, ok := request.(*VersionRequest); ok {
			return __VersionServiceGetProc__(ctx, r)
		}
		return nil, fmt.Errorf("Error convert interface : %v", reflect.TypeOf(request))
	})
}
