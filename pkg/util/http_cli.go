package util

import (
	"context"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"time"
)

var httpCli *client.Client

func InitHttpCliWithNacos(conf naming_client.INamingClient)  {
	c, err := client.NewClient (
		client.WithDialTimeout(1 * time.Second),
		client.WithWriteTimeout(500 * time.Millisecond),//写入数据超时时间，默认值：无限
		client.WithClientReadTimeout(2 * time.Second),//设置读取 response 的最长时间，默认无限长
		client.WithMaxConnWaitTimeout(100 * time.Millisecond),//设置等待空闲连接的最大时间，默认不等待
		client.WithMaxIdleConnDuration(3 * time.Second),//空闲连接超时时间
		client.WithMaxConnDuration(3 * time.Second),//设置连接存活的最大时长，超过这个时间的连接在完成当前请求后会被关闭，
		client.WithMaxConnsPerHost(1000000),
		//client.WithRetryConfig (
		//	retry.WithMaxAttemptTimes(3), // 最大的尝试次数，包括初始调用
		//	retry.WithInitDelay(1 * time.Millisecond), // 初始延迟
		//	retry.WithMaxDelay(6 * time.Millisecond), // 最大延迟，不管重试多少次，策略如何，都不会超过这个延迟
		//	retry.WithMaxJitter(2 * time.Millisecond), // 延时的最大扰动，结合 RandomDelayPolicy 才会有效果
		//	/*
		//	   配置延迟策略，你可以选择下面四种中的任意组合，最后的结果为每种延迟策略的加和
		//	   FixedDelayPolicy 使用 retry.WithInitDelay 所设置的值 ，
		//	   BackOffDelayPolicy 在 retry.WithInitDelay 所设置的值的基础上随着重试次数的增加，指数倍数增长，
		//	   RandomDelayPolicy 生成 [0，2*time.Millisecond）的随机数值 ，2*time.Millisecond 为 retry.WithMaxJitter 所设置的值，
		//	   DefaultDelayPolicy 生成 0 值，如果单独使用则立刻重试，
		//	   retry.CombineDelay() 将所设置的延迟策略所生成的值加和，最后结果即为当前次重试的延迟时间，
		//	   第一次调用失败 -> 重试延迟：1 + 1<<1 + rand[0,2)ms -> 第二次调用失败 -> 重试延迟：min(1 + 1<<2 + rand[0,2) , 6)ms -> 第三次调用成功/失败
		//	*/
		//	retry.WithDelayPolicy(retry.CombineDelay(retry.FixedDelayPolicy, retry.BackOffDelayPolicy, retry.RandomDelayPolicy)),
		//),
	)
	if err !=nil {
		hlog.Error("init client err: ", err)
		c = nil
	}
	c.Use(sd.Discovery(nacos.NewNacosResolver(conf)))
	httpCli = c
}

func InitHttpCli()  {
	c, err := client.NewClient (
		//client.WithDialTimeout(1 * time.Second),//连接建立超时时间，默认 1s
		client.WithWriteTimeout(500 * time.Millisecond),//写入数据超时时间，默认值：无限
		client.WithClientReadTimeout(2 * time.Second),//设置读取 response 的最长时间，默认无限长
		client.WithMaxConnWaitTimeout(100 * time.Millisecond),//设置等待空闲连接的最大时间，默认不等待
		client.WithMaxIdleConnDuration(3 * time.Second),//空闲连接超时时间,当超时后会关闭该连接，默认10s
		client.WithMaxConnDuration(3 * time.Second),//设置连接存活的最大时长，超过这个时间的连接在完成当前请求后会被关闭，默认无限长
		client.WithMaxConnsPerHost(1000000),
		//client.WithRetryConfig (
		//	retry.WithMaxAttemptTimes(3), // 最大的尝试次数，包括初始调用
		//	retry.WithInitDelay(1 * time.Millisecond), // 初始延迟
		//	retry.WithMaxDelay(6 * time.Millisecond), // 最大延迟，不管重试多少次，策略如何，都不会超过这个延迟
		//	retry.WithMaxJitter(2 * time.Millisecond), // 延时的最大扰动，结合 RandomDelayPolicy 才会有效果
		//	/*
		//	   配置延迟策略，你可以选择下面四种中的任意组合，最后的结果为每种延迟策略的加和
		//	   FixedDelayPolicy 使用 retry.WithInitDelay 所设置的值 ，
		//	   BackOffDelayPolicy 在 retry.WithInitDelay 所设置的值的基础上随着重试次数的增加，指数倍数增长，
		//	   RandomDelayPolicy 生成 [0，2*time.Millisecond）的随机数值 ，2*time.Millisecond 为 retry.WithMaxJitter 所设置的值，
		//	   DefaultDelayPolicy 生成 0 值，如果单独使用则立刻重试，
		//	   retry.CombineDelay() 将所设置的延迟策略所生成的值加和，最后结果即为当前次重试的延迟时间，
		//	   第一次调用失败 -> 重试延迟：1 + 1<<1 + rand[0,2)ms -> 第二次调用失败 -> 重试延迟：min(1 + 1<<2 + rand[0,2) , 6)ms -> 第三次调用成功/失败
		//	*/
		//	retry.WithDelayPolicy(retry.CombineDelay(retry.FixedDelayPolicy, retry.BackOffDelayPolicy, retry.RandomDelayPolicy)),
		//),
	)
	if err !=nil {
		hlog.Error("init client err: ", err)
		c = nil
	}
	httpCli = c
}

// Use SetQueryString to set query parameters
func Post(uri string, query string) error {
	req := &protocol.Request{}
	res := &protocol.Response{}
	req.Header.SetMethod(consts.MethodPost)
	req.SetRequestURI(uri)
	req.SetQueryString(query)//a=1&b=2
	req.Options().Apply([]config.RequestOption{config.WithSD(true)})
	return httpCli.Do(context.Background(), req, res)
}

// Send "www-url-encoded" request
func FormData4Post(uri string, body map[string]string) error {
	req := &protocol.Request{}
	res := &protocol.Response{}
	req.Header.SetMethod(consts.MethodPost)
	req.SetRequestURI(uri)
	req.Options().Apply([]config.RequestOption{config.WithSD(true)})
	req.SetMultipartFormData(body)
	return httpCli.Do(context.Background(), req, res)
}

// Send "multipart/form-data" request
func MultipartFormData4Post(uri string, body map[string]string) error {
	req := &protocol.Request{}
	res := &protocol.Response{}
	req.Header.SetMethod(consts.MethodPost)
	req.SetRequestURI(uri)
	req.Options().Apply([]config.RequestOption{config.WithSD(true)})
	req.SetMultipartFormData(body)
	return httpCli.Do(context.Background(), req, res)
}

// Send "Json" request
func Json4Post(uri string, body interface{}) error {
	req := &protocol.Request{}
	res := &protocol.Response{}
	req.Header.SetMethod(consts.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.SetRequestURI(uri)
	req.Options().Apply([]config.RequestOption{config.WithSD(true)})
	jsonByte, _ := sonic.Marshal(body)
	req.SetBody(jsonByte)
	return httpCli.Do(context.Background(), req, res)
}

// ?/path/encode uri , method: get/post/delete/put
func HttpDo(uri string, method string) (*HTTPResp, error) {
	req := &protocol.Request{}
	res := &protocol.Response{}
	req.Header.SetMethod(method)
	req.SetRequestURI(uri)
	//req.Header.SetHostBytes(req.URI().Host())
	req.Options().Apply([]config.RequestOption{config.WithSD(true)})
	err := httpCli.Do(context.Background(), req, res)
	if err != nil {
		hlog.Error(err.Error())
		return nil, err
	}
	return &HTTPResp {
		Status: res.StatusCode(),
		Data:   res.Body(),
	},nil
}

//form submit by post/put/update
func DoByForm(uri string, method string, formData map[string]string) (*HTTPResp, error) {
	req := &protocol.Request{}
	res := &protocol.Response{}
	req.Header.SetMethod(method)
	req.SetRequestURI(uri)
	//req.SetFormData(formData)
	req.SetMultipartFormData(formData)
	//req.Header.SetHostBytes(req.URI().Host())
	req.Options().Apply([]config.RequestOption{config.WithSD(true)})
	err := httpCli.Do(context.Background(), req, res)
	if err != nil {
		hlog.Error(err.Error())
		return nil, err
	}
	return &HTTPResp {
		Status: res.StatusCode(),
		Data:   res.Body(),
	},nil
}
