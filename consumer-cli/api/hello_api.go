package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"primus/consumer-cli/hello/hello_service"
	"time"
)

func GetH(ctx context.Context, c *app.RequestContext) {
	var q int32 = 100000
	//cc, _ := nacos.NewNacosConfig("121.37.173.206", 8848)
	//r := nacoscli.NewNacosResolver(cc)

	cli, _ := hello_service.NewHelloServiceClient(
		"http://producer-service",
		//"http://127.0.0.1:9809",
		//hello_service.WithHertzClientMiddleware(), // 指定 client 的中间件
	)
	resp, rawResp, err := cli.GetH(
		context.Background(),
		&q,
		// 在发起调用的时候可指定请求级别的配置
		//config.
		config.WithSD(true),                      // 指定请求级别的设置，用来开启服务发现
		config.WithReadTimeout(2000*time.Second), // 指定请求读超时
	)
	if err != nil {
		hlog.Error(err)
		//c.JSON
		return
	}
	hlog.Info(rawResp.StatusCode())
	hlog.Info(resp)
	c.JSON(consts.StatusOK, resp)
}
