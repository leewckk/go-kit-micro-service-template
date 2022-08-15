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
  host : "127.0.0.1
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

