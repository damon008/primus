# Copyright Statement

Open source does not mean free of charge. This project is released under the GPL-3.0 open source license and provides technical exchange and learning, but it is strictly forbidden to modify and derive code as closed-source commercial software for release and sales! If you need to use this product with any attached commercialization behavior locally, please contact the project leader for commercial authorization to comply with the GPL agreement to ensure your normal use.

Currently in China, the GPL agreement has contractual characteristics and is a civil legal act, within the scope of China's "Contract Law". The original team of this project reserves all litigation rights.

[Related Case: Violation of GPL Agreement Compensation 500,000, the First Case in China!]((https://mp.weixin.qq.com/s/YQ6sNjbDS-P7BViLZIsaoA))

**PS:Our team has the final interpretation right of this open source agreement.**

## Related Technical Support

| 拓展                                                                                                 | 描述                                                                                                |
|----------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------|
| [Autotls](https://github.com/hertz-contrib/autotls)                                                | 为 Hertz 支持 Let's Encrypt 。                                                                        |
| [Http2](https://github.com/hertz-contrib/http2)                                                    | 提供对 HTTP2 的支持。                                                                                    |
| [Websocket](https://github.com/hertz-contrib/websocket)                                            | 使 Hertz 支持 Websocket 协议。                                                                          |
| [Etag](https://github.com/hertz-contrib/etag)                                                      | 提供 ETag HTTP 响应标头。                                                                                |
| [Limiter](https://github.com/hertz-contrib/limiter)                                                | 提供了基于 bbr 算法的限流器。                                                                                 |
| [Monitor-prometheus](https://github.com/hertz-contrib/monitor-prometheus)                          | 提供基于 Prometheus 服务监控功能。                                                                           |
| [Obs-opentelemetry](https://github.com/hertz-contrib/obs-opentelemetry)                            | Hertz 的 Opentelemetry 扩展，支持 Metric、Logger、Tracing 并且达到开箱即用。                                       |
| [Opensergo](https://github.com/hertz-contrib/opensergo)                                            | Opensergo 扩展。                                                                                     |
| [Pprof](https://github.com/hertz-contrib/pprof)                                                    | Hertz 集成 Pprof 的扩展。                                                                               |
| [Registry](https://github.com/hertz-contrib/registry)                                              | 提供服务注册与发现功能。到现在为止，支持的服务发现拓展有 nacos， consul， etcd， eureka， polaris， servicecomb， zookeeper， redis。 |
| [Sentry](https://github.com/hertz-contrib/hertzsentry)                                             | Sentry 拓展提供了一些统一的接口来帮助用户进行实时的错误监控。                                                                |
| [Tracer](https://github.com/hertz-contrib/tracer)                                                  | 基于 Opentracing 的链路追踪。                                                                             |
| [Basicauth](https://github.com/cloudwego/hertz/tree/develop/pkg/app/middlewares/server/basic_auth) | Basicauth 中间件能够提供 HTTP 基本身份验证。                                                                    |
| [Jwt](https://github.com/hertz-contrib/jwt)                                                        | Jwt 拓展。                                                                                           |
| [Keyauth](https://github.com/hertz-contrib/keyauth)                                                | 提供基于 token 的身份验证。                                                                                 |
| [Requestid](https://github.com/hertz-contrib/requestid)                                            | 在 response 中添加 request id。                                                                        |
| [Sessions](https://github.com/hertz-contrib/sessions)                                              | 具有多状态存储支持的 Session 中间件。                                                                           |
| [Casbin](https://github.com/hertz-contrib/casbin)                                                  | 通过 Casbin 支持各种访问控制模型。                                                                             |
| [Cors](https://github.com/hertz-contrib/cors)                                                      | 提供跨域资源共享支持。                                                                                       |
| [Csrf](https://github.com/hertz-contrib/csrf)                                                      | Csrf 中间件用于防止跨站点请求伪造攻击。                                                                            |
| [Secure](https://github.com/hertz-contrib/secure)                                                  | 具有多配置项的 Secure 中间件。                                                                               |
| [Gzip](https://github.com/hertz-contrib/gzip)                                                      | 含多个可选项的 Gzip 拓展。                                                                                  |
| [I18n](https://github.com/hertz-contrib/i18n)                                                      | 可帮助将 Hertz 程序翻译成多种语言。                                                                             |
| [Lark](https://github.com/hertz-contrib/lark-hertz)                                                | 在 Hertz 中处理 Lark/飞书的卡片消息和事件的回调。                                                                   |
| [Loadbalance](https://github.com/hertz-contrib/loadbalance)                                        | 提供适用于 Hertz 的负载均衡算法。                                                                              |
| [Logger](https://github.com/hertz-contrib/logger)                                                  | Hertz 的日志拓展，提供了对 zap、logrus、zerologs 日志框架的支持。                                                     |
| [Recovery](https://github.com/cloudwego/hertz/tree/develop/pkg/app/middlewares/server/recovery)    | Hertz 的异常恢复中间件。                                                                                   |
| [Reverseproxy](https://github.com/hertz-contrib/reverseproxy)                                      | 实现反向代理。                                                                                           |
| [Swagger](https://github.com/hertz-contrib/swagger)                                                | 使用 Swagger 2.0 自动生成 RESTful API 文档。                                                               |
| [Cache](https://github.com/hertz-contrib/cache)                                                    | 用于缓存 HTTP 接口内容的 Hertz 中间件，支持多种客户端。                                                                |


### Stress Test Conclusion
Under normal requests, Kitex-mux performs the best, followed by Kitex. It is recommended to use Thrift encoding and decoding first.

Under the GRPC protocol, Kitex performs the best and is stronger than native gRPC. It is recommended to use Thrift encoding and decoding first.

Under Streaming, Kitex performs the best based on Protobuf encoding, with two options: Kitex Protobuf and gRPC.

Under Thrift and Protobuf, Kitex-mux performs the best, followed by Kitex. It is recommended to use Thrift encoding and decoding first.

When using Thrift as Payload, Kitex's performance exceeds that of official gRPC, with throughput close to twice that of gRPC; In addition, when Kitex uses customized Protobuf protocol, its performance is also better than gRPC.

### Project Structure

- hello-server: This is the kitex server-side demo. 
- producer-service: This is the kitex cli-side demo and also the hertz server-side demo. 
- consumer-cli: This is the hertz cli-side demo and also the hertz server-side demo.
- idl: This is where the project code generates the idl file, including Thrift and pb files.
- kitex_gen: This is the middleware required for kitex cli and server-side communication generated based on idl files, making communication more convenient.
- pkg: This is a resource pool containing all dependent packages, functions, tools, and other resources.

### Code Generation

```
go install github.com/cloudwego/hertz/cmd/hz@latest
go install github.com/cloudwego/kitex/tool/server/kitex@latest
go install github.com/cloudwego/thriftgo@latest

#For kitex server, remove ',template=slim' because it is not friendly to discover.
kitex -module "primus" -thrift frugal_tag -service hello idl/hello.thrift
mkdir hello-server
cd hello-server
...

# Code server for hertz
# No need to add the parameter -t=template=slim for encoding and decoding code. The same applies to pb files.
cd producer-service
hz new -module producer-service --idl=../idl/hello.thrift -t=template=slim
hz update --idl=../idl/hello.thrift
# Hertz client
cd consumer-cli
hz client --mod=primus/consumer-cli --idl=../idl/hello.thrift --model_dir=model -t=template=slim --client_dir=hello
```
### Build
```
cd primus/producer-service
go build -o producer-service
```

## License Scene
During this development, a license scene was designed for software licensing, such as a service producer that specializes in deploying software on customer machines to check whether they have been licensed.

In the service producer, start a timer in main. The logic is to perform authorization authentication after obtaining the authorization license to see if it has been authorized: status.
Generate an authorization interface in the service to provide other services with a query of the authorization status (also generate a license in the service).
When all services perform global routing, query the authorization interface in the previous service to see if the status is OK.

## API requests

### Authorization Code Generation

```
curl -H "Content-Type: application/json" -X POST -d '{"appId": "***.com","issuedTime": 1669082100,"notBefore": 1669082100,"notAfter": 1669116589,"customerInfo": "***公司","authorization": "all,training,inference","machineCodes": ["JL32YL2"],"nodeNum":1}' "http://localhost:2999/v1/createLicence"
```

### Special thanks

This framework is based on the open-source hertz and kitex frameworks developed by ByteDance. Thank you for ByteDance's strong support: https://github.com/cloudwego.
