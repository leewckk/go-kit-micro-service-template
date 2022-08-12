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

// create time : 2022-08-12 15:15:43.114428401 +0800 CST m=+0.007161932

package version

import (
	context "context"
	fmt "fmt"
	grpc "github.com/go-kit/kit/transport/grpc"
	grpc1 "google.golang.org/grpc"
	version1 "micro-service-sample/endpoint/version"
	version "micro-service-sample/pb/golang/pkg/version"
	reflect "reflect"
)

type VersionServiceServiceTransport struct {
	__Get__ grpc.Handler
}

func RegisterVersionServiceProc(s *grpc1.Server, opts ...grpc.ServerOption) {
	version.RegisterVersionServiceServer(s, NewVersionServiceTransport(opts...))
}

func NewVersionServiceTransport(opts ...grpc.ServerOption) *VersionServiceServiceTransport {

	var t VersionServiceServiceTransport
	t.__Get__ = grpc.NewServer(
		version1.MakeVersionServiceGetEndpoint(),
		DecodeVersionServiceGetRequest,
		EncodeVersionServiceGetResponse,
		opts...,
	)
	return &t
}

///
func (this *VersionServiceServiceTransport) Get(ctx context.Context, request *version.VersionRequest) (*version.VersionResponse, error) {
	_, resp, err := this.__Get__.ServeGRPC(ctx, request)
	if nil != err {
		return nil, err
	}
	return resp.(*version.VersionResponse), nil
}

///

//// THIS IS FOR SERVICE DEOCDE REQUEST
//// PB.REQUEST -> SERVICE.REQUEST
func DecodeVersionServiceGetRequest(ctx context.Context, request interface{}) (interface{}, error) {
	return &version1.VersionRequest{}, nil
}

////

//// THIS IS FOR SERVICE ENCODE RESPONSE
//// SERVICE.RESPONSE -> PB.RESPONSE
func EncodeVersionServiceGetResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	if response, ok := resp.(*version1.VersionResponse); ok {
		r := version.VersionResponse{}
		/// TODO
		r.Hash = response.Hash
		r.Tag = response.Tag
		r.Version = response.Version
		r.BuilderVersion = response.BuilderVersion
		r.Uptime = response.Uptime
		r.RunningTimes = response.RunningTimes
		r.BuildTime = response.BuildTime
		r.HostName = response.HostName
		r.Platform = response.Platform
		return &r, nil
	}
	return nil, fmt.Errorf("Error convert interface : %v -> *%v", reflect.TypeOf(resp), "version1.VersionResponse")
}

////
