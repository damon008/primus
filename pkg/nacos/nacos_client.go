package nacos

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/registry/nacos/common"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"os"
)

func NewNacosConfig(ip string) (naming_client.INamingClient, error) {
	hlog.Debugf("connect Nacos start")
	sc := []constant.ServerConfig{
		//*constant.NewServerConfig(os.Getenv("serverAddr"), uint64(NacosPort())),
		*constant.NewServerConfig(ip, uint64(8848)),
	}
	cc := constant.ClientConfig {
		NamespaceId:         os.Getenv("namespace"),
		RegionId:            "cn-hangzhou",
		CustomLogger:        common.NewCustomNacosLogger(),
		NotLoadCacheAtStart: true,
	}
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}

/*func NacosPort() int64 {
	portText := os.Getenv(NACOS_ENV_PORT)
	if len(portText) == 0 {
		return NACOS_DEFAULT_PORT
	}
	port, err := strconv.ParseInt(portText, 10, 64)
	if err != nil {
		klog.Errorf("ParseInt failed,err:%s", err.Error())
		return NACOS_DEFAULT_PORT
	}
	return port
}

// NacosAddr Get Nacos addr from environment variables
func NacosAddr() string {
	addr := os.Getenv(NACOS_ENV_SERVER_ADDR)
	if len(addr) == 0 {
		return NACOS_DEFAULT_SERVER_ADDR
	}
	return addr
}*/
