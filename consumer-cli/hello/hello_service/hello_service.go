// Code generated by hertz generator.

package hello_service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	nacoscli "github.com/hertz-contrib/registry/nacos"
	"primus/pkg/nacos"
	"strconv"

	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol"
	hello "primus/consumer-cli/model/hello"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
)

type Client interface {
	GetByParams(context context.Context, req *string, reqOpt ...config.RequestOption) (resp *hello.Response, rawResponse *protocol.Response, err error)

	GetH(context context.Context, req *int32, reqOpt ...config.RequestOption) (resp *hello.Response, rawResponse *protocol.Response, err error)

	Post(context context.Context, req *hello.Request, reqOpt ...config.RequestOption) (resp *hello.Response, rawResponse *protocol.Response, err error)
}

type HelloServiceClient struct {
	client *cli
}

func NewHelloServiceClient(hostUrl string, ops ...Option) (Client, error) {
	cc, _ := nacos.NewNacosConfig("121.37.173.206", 8848)

	opts := getOptions(append(ops, withHostUrl(hostUrl))...)
	cli, err := newClient(opts)
	cli.Use(sd.Discovery(nacoscli.NewNacosResolver(cc)))
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{
		client: cli,
	}, nil
}

func (s *HelloServiceClient) GetByParams(context context.Context, req *string, reqOpt ...config.RequestOption) (resp *hello.Response, rawResponse *protocol.Response, err error) {
	httpResp := &hello.Response{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/api/v1/getByParams/:data")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *HelloServiceClient) GetH(context context.Context, req *int32, reqOpt ...config.RequestOption) (resp *hello.Response, rawResponse *protocol.Response, err error) {
	httpResp := &hello.Response{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{
			"id": strconv.Itoa(int(*req)),
		}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/api/v1/getH/:id")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *HelloServiceClient) Post(context context.Context, req *hello.Request, reqOpt ...config.RequestOption) (resp *hello.Response, rawResponse *protocol.Response, err error) {
	httpResp := &hello.Response{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("POST", "/api/v1/create")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

var defaultClient, _ = NewHelloServiceClient(
	"http://127.0.0.1:9809",
	 //WithHertzClientMiddleware(), // 指定 client 的中间件
	)

func ConfigDefaultClient(ops ...Option) (err error) {
	defaultClient, err = NewHelloServiceClient("", ops...)
	return
}

func GetByParams(context context.Context, req *string, reqOpt ...config.RequestOption) (resp *hello.Response, rawResponse *protocol.Response, err error) {
	return defaultClient.GetByParams(context, req, reqOpt...)
}

func GetH(context context.Context, req *int32, reqOpt ...config.RequestOption) (resp *hello.Response, rawResponse *protocol.Response, err error) {
	return defaultClient.GetH(context, req, reqOpt...)
}

func Post(context context.Context, req *hello.Request, reqOpt ...config.RequestOption) (resp *hello.Response, rawResponse *protocol.Response, err error) {
	return defaultClient.Post(context, req, reqOpt...)
}
