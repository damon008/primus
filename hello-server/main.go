package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote/codec/thrift"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"net"
	"primus/hello-server/handler"
	hello "primus/kitex_gen/hello/helloservice"
	"primus/pkg/constants"
	nacosCli "primus/pkg/nacos"
)

func main() {
	conf, err := nacosCli.NewNacosConfig("121.37.173.206", 8848)
	if err != nil {
		klog.Error(err)
	}

	svr := hello.NewServer(
		new(handler.HelloServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.HelloServiceName}),
		server.WithServiceAddr(&net.TCPAddr{
			Port: 1000,
		}),
		//server.WithListener(),
		server.WithRegistry(registry.NewNacosRegistry(conf)),

		//server.WithCodec(codec.NewDefaultCodecWithSizeLimit(1024*1024*10)),

		server.WithPayloadCodec(thrift.NewThriftCodecWithConfig(thrift.FrugalRead|thrift.FrugalWrite|thrift.FastRead|thrift.FastWrite)),
		//连接多路复用(mux)，GRPC默认多路复用
		server.WithMuxTransport(),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
	)

	err = svr.Run()

	if err != nil {
		klog.Fatal(err.Error())
	}
}
