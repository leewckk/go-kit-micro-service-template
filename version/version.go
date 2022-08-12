// The MIT License (MIT)
//
// Copyright (c) 2022 leewckk@126.com
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
//

package version

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type VersionInfo struct {
	GitHash      string    `json:"hash"`
	GitTag       string    `json:"tag"`
	Version      string    `json:"version"`
	GoVersion    string    `json:"builderVersion"`
	Uptime       time.Time `json:"uptime"`
	RunningTimes string    `json:"runningTimes"`
	BuildTime    string    `json:"build_time"`
	HostName     string    `json:"hostname"`
	Platform     string    `json:"platform"`
}

var __ver__ VersionInfo

func RegisterVersionInfo(hash string, tag string, version string, gover string, buildTime string) {
	__ver__.GitHash = hash
	__ver__.GitTag = tag
	__ver__.Version = version
	__ver__.GoVersion = gover
	__ver__.BuildTime = buildTime
	__ver__.Uptime = time.Now()
	__ver__.HostName, _ = os.Hostname()
	__ver__.Platform = runtime.GOOS
}

func GetVersion() *VersionInfo {
	__ver__.RunningTimes = fmt.Sprintf("%vs", int64(time.Now().Sub(__ver__.Uptime).Seconds()))
	return &__ver__
}
