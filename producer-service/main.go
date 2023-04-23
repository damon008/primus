// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/registry/nacos"
	nacosCli "primus/pkg/nacos"
	"primus/pkg/util"
	"primus/producer-service/biz/router"
	"primus/producer-service/biz/rpc"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func main() {
	conf,err := nacosCli.NewNacosConfig("121.37.173.206", 8848)
	if err != nil {
		hlog.Error(err)
	}


	h := server.Default(
		server.WithHostPorts(":9809"),
		server.WithRegistry(nacos.NewNacosRegistry(conf), &registry.Info{
			ServiceName: "producer-service",
			Addr:        utils.NewNetAddr("tcp", util.GetIpAddr()+":9809"), //&net.TCPAddr{Port: 3000},
			Weight:      10,
			Tags:        nil,
		}),
	)

	rpc.InitRPC()
	router.GeneratedRegister(h)
	h.Spin()
}
