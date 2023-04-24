package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/registry/nacos"
	"k8s.io/apimachinery/pkg/util/wait"
	"primus/consumer-cli/router"
	"primus/pkg/license"
	nacosCli "primus/pkg/nacos"
	"primus/pkg/util"
	"primus/pkg/util/singleton"
	"time"

	//"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	//"primus/pkg/nacos"
	//nacoscli "github.com/hertz-contrib/registry/nacos"
)

func main() {
	conf, err := nacosCli.NewNacosConfig("121.37.173.206", 8848)
	if err != nil {
		hlog.Error(err)
	}

	singleton.GetHttpCli()

	stop := make(chan struct{})
	//服务启动后会检测5分钟检测一次软件授权，第一次看授权,全局变量
	go wait.Until(license.LicenceChecker, 5*time.Minute, stop)

	h := server.Default(
		server.WithHostPorts(":10000"),
		server.WithRegistry(nacos.NewNacosRegistry(conf), &registry.Info{
			ServiceName: "consumer-cli",
			Addr:        utils.NewNetAddr("tcp", util.GetIpAddr()+":10000"),
			Weight:      10,
			Tags:        nil,
		}),
	)
	router.Init(h)
	h.Spin()
}
