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

package configure

import (
	"github.com/spf13/viper"
)

type HttpServerConfig struct {
	Port   int    `json:"port"`
	Health string `json:"health"`
}

type GrpcServerConfig struct {
	Port   int    `json:"port"`
	Health string `json:"health"`
}

type ServerConfig struct {
	Http HttpServerConfig `json:"http"`
	Grpc GrpcServerConfig `json:"grpc"`
}

type DiscoveryConfig struct {
	ServiceName string `json:"serviceName"`
	Type        string `json:"type"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Api         string `json:"api"`
}

type TracingConfig struct {
	Type     string `json:"type"`
	Reporter string `json:"reporter"`
}

type DBConfig struct {
	Driver string `json:"driver"`
	Conn   string `json:"conn"`
	Debug  bool   `json:"debug"`
}

type RedisConfig struct {
	Server   string `json:"server"`
	Conn     string `json:"conn"`
	Dbnum    int    `json:"dbnum"`
	Password string `json:"password"`
}

type CalculatorConfig struct {
	Host string `json:"host"`
}

type LogConfig struct {
	LogstoreProject string `json:"logstoreProject"`
	LogstoreName    string `json:"logstoreName"`
	FormatJson      bool   `json:"formatJson"`
	ReportCaller    bool   `json:"reportCaller"`
}

type Config struct {
	Project    string           `json:"project"`
	Runmode    string           `json:"runmode"`
	Area       string           `json:"area"`
	Server     ServerConfig     `json:"server"`
	Discovery  DiscoveryConfig  `json:"discovery"`
	Tracing    TracingConfig    `json:"tracing"`
	Db         DBConfig         `json:"db"`
	Redis      RedisConfig      `json:"redis"`
	Calculator CalculatorConfig `json:"calculator"`
	Log        LogConfig        `json:"log"`
}

type Configure struct {
	Uri  string
	conf Config
	cfg  *viper.Viper
}

var __config_default__ Configure

const (
	CONFIG_URI_DEFAULT = "/etc/micro-service/conf/conf.yaml"
)

func init() {
	__config_default__.cfg = viper.New()
	__config_default__.Uri = CONFIG_URI_DEFAULT
	__config_default__.load()
}

func NewConfigure(uri string) *Configure {
	c := &Configure{Uri: uri,
		cfg: viper.New(),
	}
	err := c.load()
	if nil != err {
		panic(err)
	}
	return c
}

func GetDefault() *Configure {
	return &__config_default__
}

func (this *Configure) SetServiceName(name string) error {
	this.conf.Discovery.ServiceName = name
	return nil
}

func (this *Configure) load() error {
	this.cfg.SetConfigFile(this.Uri)

	this.cfg.SetDefault("area", "china")
	this.cfg.SetDefault("server.http.port", 8888)
	this.cfg.SetDefault("server.http.health", "/actuator/info")
	this.cfg.SetDefault("server.grpc.port", 6666)

	this.cfg.SetDefault("tracing.type", "zipkin")
	this.cfg.SetDefault("tracing.reporter", "http://127.0.0.1:9411/api/v2/spans")
	err := this.cfg.ReadInConfig()
	if nil != err {
		return err
	}

	err = this.cfg.Unmarshal(&this.conf)
	if nil != err {
		return err
	}

	return nil
}

func (this *Configure) Config() *Config {
	return &this.conf
}
