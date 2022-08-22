## go-kit-micro-service-template

[go-kit-micro-service-template](https://github.com/leewckk/go-kit-micro-service-template),是[go-kit-micro-service](https://github.com/leewckk/go-kit-micro-service)使用的参考模板，使用方法如下：

#### step 1 clone代码

````shell
git clone https://github.com/leewckk/go-kit-micro-service-template.git
````



#### step 2 初始化项目

假设项目名称为一个处理订单的业务，项目名称`service-bill`

````shell
./init.sh service-bill
````

#### step 3 定义proto文件

[`protoc-gen-gokit-micro`](https://github.com/leewckk/protoc-gen-gokit-micro)，提供了[`google.api.http`](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto)的支持，可以通过在`proto`中定义`option`来完成`http`相关参数以及`API`的定义。以下列举了一个简单查询版本信息的`proto`文件：

````protobuf
syntax = "proto3";
package version;
option go_package="pb/golang/pkg/version;version";

import "google/api/annotations.proto";

message VersionRequest {
}

message VersionInfo {
    string hash = 1 ;                       /// GIT HASH信息
    string tag = 2;                         /// git TAG信息
    string version = 3;                     /// 版本信息
    string builderVersion = 4;              /// 编译器版本号
    string uptime = 5;                      /// 启动时间
    string runningTimes = 6;                /// 运行时间
    string buildTime = 7;                   /// 编译时间
    string hostName = 8;                    /// HOSTNAME    
    string platform = 9;                    /// PLATFORM
}

message VersionResponse{
    string hash = 1 ;                       /// GIT HASH信息
    string tag = 2;                         /// git TAG信息
    string version = 3;                     /// 版本信息
    string builderVersion = 4;              /// 编译器版本号
    string uptime = 5;                      /// 启动时间
    string runningTimes = 6;                /// 运行时间
    string buildTime = 7;                   /// 编译时间
    string hostName = 8;                    /// HOSTNAME    
    string platform = 9;                    /// PLATFORM
}

service VersionService {
    rpc Get( VersionRequest ) returns (VersionResponse) {
        option(google.api.http) = {
            get : "/v1/version"
        };
    }
}
````



#### step 4 通过`auto-gen.sh`生成项目代码

````she
./auto-gen.sh
````



#### docker 测试环境

本项目模板提供了微服务测试所需环境的最低`docker-compose`[配置](docker/consul-zipkin), 具体参考: [consul-zipkin docker 环境安装](docker/consul-zipkin/README.md)

#### 配置文件

默认参考配置文件：

````yaml
project : micro-service

server:
  http:
    port : 8888
    health: "/actuator/info"
  grpc :
    address: "0.0.0.0:6666"
    port : 6666

discovery :
  servicename : "micro-service"
  type : "consul"
  host : "127.0.0.1"
  port : 8500
  api : "127.0.0.1:8500"

tracing :
  type : "zipkin"
  reporter : "http://localhost:9411/api/v2/spans"

````

#### 关键参数

**server**

`server.http.port` , 定义`http`服务的绑定端口

`server.http.health` , 定义`http`的健康接口

`server.grpc.address` , 定义`grpc`的绑定地址

`server.grpc.port` ,定义`grpc`的绑定端口



**discovery**

`discovery.servicename` ， 定义注册到注册中心的服务名；

`discovery.type`,注册中心的服务器类型, 暂时只支持`consul`

`discovery.host`,注册中心`IP`地址

`discovery.port`,注册中心端口

`discovery.api`,注册中心的API



**tracing**

`tracing.type` , 链路追踪的服务器类型, 暂时只支持`zipkin`

`tracing.reporter`, 链路追踪API接口



配置文件路径可以通过调整[$project_dir/configure/configure.go](https://github.com/leewckk/go-kit-micro-service-template/blob/master/configure/configure.go#L105)中定义的路径来调整配置路径：

````go
const (
	CONFIG_URI_DEFAULT = "/etc/micro-service/conf/conf.yaml"
)
````

