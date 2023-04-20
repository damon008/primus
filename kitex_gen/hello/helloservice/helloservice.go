// Code generated by Kitex v0.5.2. DO NOT EDIT.

package helloservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	hello "primus/kitex_gen/hello"
)

func serviceInfo() *kitex.ServiceInfo {
	return helloServiceServiceInfo
}

var helloServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "HelloService"
	handlerType := (*hello.HelloService)(nil)
	methods := map[string]kitex.MethodInfo{
		"echo": kitex.NewMethodInfo(echoHandler, newHelloServiceEchoArgs, newHelloServiceEchoResult, false),
		"Get":  kitex.NewMethodInfo(getHandler, newHelloServiceGetArgs, newHelloServiceGetResult, false),
		"GetH": kitex.NewMethodInfo(getHHandler, newHelloServiceGetHArgs, newHelloServiceGetHResult, false),
		"Post": kitex.NewMethodInfo(postHandler, newHelloServicePostArgs, newHelloServicePostResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "hello",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.2",
		Extra:           extra,
	}
	return svcInfo
}

func echoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*hello.HelloServiceEchoArgs)
	realResult := result.(*hello.HelloServiceEchoResult)
	success, err := handler.(hello.HelloService).Echo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newHelloServiceEchoArgs() interface{} {
	return hello.NewHelloServiceEchoArgs()
}

func newHelloServiceEchoResult() interface{} {
	return hello.NewHelloServiceEchoResult()
}

func getHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {

	realResult := result.(*hello.HelloServiceGetResult)
	success, err := handler.(hello.HelloService).Get(ctx)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newHelloServiceGetArgs() interface{} {
	return hello.NewHelloServiceGetArgs()
}

func newHelloServiceGetResult() interface{} {
	return hello.NewHelloServiceGetResult()
}

func getHHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*hello.HelloServiceGetHArgs)
	realResult := result.(*hello.HelloServiceGetHResult)
	success, err := handler.(hello.HelloService).GetH(ctx, realArg.Id)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newHelloServiceGetHArgs() interface{} {
	return hello.NewHelloServiceGetHArgs()
}

func newHelloServiceGetHResult() interface{} {
	return hello.NewHelloServiceGetHResult()
}

func postHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*hello.HelloServicePostArgs)
	realResult := result.(*hello.HelloServicePostResult)
	success, err := handler.(hello.HelloService).Post(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newHelloServicePostArgs() interface{} {
	return hello.NewHelloServicePostArgs()
}

func newHelloServicePostResult() interface{} {
	return hello.NewHelloServicePostResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Echo(ctx context.Context, req *hello.Request) (r *hello.Response, err error) {
	var _args hello.HelloServiceEchoArgs
	_args.Req = req
	var _result hello.HelloServiceEchoResult
	if err = p.c.Call(ctx, "echo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Get(ctx context.Context) (r *hello.Response, err error) {
	var _args hello.HelloServiceGetArgs
	var _result hello.HelloServiceGetResult
	if err = p.c.Call(ctx, "Get", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetH(ctx context.Context, id int32) (r *hello.Response, err error) {
	var _args hello.HelloServiceGetHArgs
	_args.Id = id
	var _result hello.HelloServiceGetHResult
	if err = p.c.Call(ctx, "GetH", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Post(ctx context.Context, req *hello.Request) (r *hello.Response, err error) {
	var _args hello.HelloServicePostArgs
	_args.Req = req
	var _result hello.HelloServicePostResult
	if err = p.c.Call(ctx, "Post", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
