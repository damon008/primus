package rpc

import (
	"context"
	"github.com/bytedance/gopkg/cloud/circuitbreaker"
	"github.com/bytedance/gopkg/lang/fastrand"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/codec"
	"github.com/cloudwego/kitex/pkg/remote/codec/thrift"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	trace "github.com/kitex-contrib/tracer-opentracing"
	api "primus/kitex_gen/hello"
	helloService "primus/kitex_gen/hello/helloservice"
	//"primus/pkg/bound"
	cb "primus/pkg/circuitbreak"
	"primus/pkg/constants"
	nacosCli "primus/pkg/nacos"
	"primus/kitex_gen/hello"
	"strconv"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

var helloClient helloService.Client

func initHelloRpc() {
	//r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})

	conf,err := nacosCli.NewNacosConfig("121.37.173.206", 8848)
	if err != nil {
		klog.Error(err)
	}

	r1 := resolver.NewNacosResolver(conf)

	lb := loadbalance.NewConsistBalancer(loadbalance.ConsistentHashOption {
		GetKey: func(ctx context.Context, request interface{}) string {
			return strconv.Itoa(fastrand.Intn(100000))
		},
		/*GetKey:  func(ctx context.Context, request interface{}) string {
			key, _ := ctx.Value(ctxConsistentKey).(string)
			if len(key) == 0 {
				return "1234"
			}
			return key
		},*/
		// 是否使用 replica
		// 如果使用，当请求失败（连接失败）后会依次尝试 replica
		// 会带来额外内存和计算开销
		// 如果不设置，那么请求失败（连接失败）后直接返回
		Replica:        1,
		// 虚拟节点数
		// 每个真实节点对应的虚拟节点的数量
		// 这个数值越大，内存和计算代价越大，负载越均衡
		// 当节点数多时，可以适当设小一些；反之可以适当设大一些
		// 推荐 VirtualFactor * Weight（如果 Weighted 为 true）的中位数在 1000 左右，负载应当已经很均衡了
		// 推荐 总虚拟节点数 在 2000W 以内（1000W 情况之下 build 一次需要 250ms，不过为后台 build 理论上 3s 内均无问题）
		VirtualFactor:  1000,
		// 是否要遵循 Weight 进行负载均衡
		// 如果为 false，对于每个 instance 都会忽略 Weight，均生成 VirtualFactor 个虚拟节点，进行无差别负载均衡
		// 如果为 true，对于每个 instance 会生成 instance.Weight() * VirtualFactor 个虚拟节点
		// 需要注意，对于 weight 为 0 的 instance，无论 VirtualFactor 为多少，均不会生成虚拟节点
		// 建议设为 true，不过要注意适当调小 VirtualFactor
		Weighted:       true,
		// 是否进行过期处理
		// 实现会缓存所有的 Key
		// 如果永不过期会导致内存一直增长
		// 设置过期会导致额外性能开销
		// 目前的实现是每分钟扫描删除一次，以及实例发生变动 rebuild 时删除一次
		// 建议一定要设置，值不要小于一分钟
		ExpireDuration: 5 * time.Minute,
	})

	// Use kitex client to make rpc calls and request back-end services

	//异常重试：提高服务整体的成功率
	//Backup Request：减少服务的延迟波动
	fp := retry.NewFailurePolicy()
	// 重试次数, 默认2，不包含首次请求
	fp.WithMaxRetryTimes(2) // 合法值：[0-5]
	// 总耗时，包括首次失败请求和重试请求耗时达到了限制的duration，则停止后续的重试。
	fp.WithMaxDurationMS(8000)

	// 关闭链路中止
	fp.DisableChainRetryStop()

	// 开启DDL中止
	fp.WithDDLStop()

	// 退避策略，默认无退避策略
	fp.WithFixedBackOff(1000) // 固定时长退避
	fp.WithRandomBackOff(1000, 10000) // 随机时长退避

	// 开启重试熔断
	//fp.WithRetryBreaker(float64(1))
	// 同一节点重试
	//fp.WithRetrySameNode()

	bp := retry.NewBackupPolicy(50)
	bp.WithMaxRetryTimes(2)
	// 关闭链路中止
	bp.DisableChainRetryStop()

	// 开启重试熔断
	//bp.WithRetryBreaker(float64(1))

	// 同一节点重试
	//bp.WithRetrySameNode()

	// TODO 熔断机制
	opt := circuitbreaker.Options {
		BucketTime:                0,
		BucketNums:                0,
		CoolingTimeout:            0,
		DetectTimeout:             0,
		HalfOpenSuccesses:         0,
		//ShouldTrip:                circuitbreaker.ConsecutiveTripFunc(5),//连续错误数
		ShouldTrip:                circuitbreaker.RateTripFunc(0.5, 10),//错误率
		//ShouldTrip:                circuitbreaker.ThresholdTripFunc(5),//错误数达到阈值
		ShouldTripWithKey:         nil,
		BreakerStateChangeHandler: nil,
		EnableShardP:              false,
		Now:                       nil,
	}
	cbPanel, err := circuitbreaker.NewPanel(cb.ChangeHandler, opt)
	if err != nil {
		klog.Error(err)
	}
	cbCtrl := circuitbreak.Control{GetKey: cb.GetKey, GetErrorType: cb.GetErrorType, DecorateError: cb.DecorateError}
	cbMW := circuitbreak.NewCircuitBreakerMW(cbCtrl, cbPanel)


	//cbs := circuitbreak.NewCBSuite(circuitbreak.GenServiceCBKeyFunc)

	c, err := helloService.NewClient (
		constants.HelloServiceName,
		client.WithSuite(tracing.NewClientSuite()),
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.ApiServiceName}),
		//client.WithCircuitBreaker(cbs),
		//client.WithMiddleware(cb.FailMW),
		client.WithMiddleware(cbMW),

		//client.WithDialer(netpoll.NewDialer()),
		client.WithPayloadCodec(thrift.NewThriftCodecWithConfig(thrift.FastRead | thrift.FastWrite)),
		client.WithLoadBalancer(lb),

		client.WithTracer(prometheus.NewClientTracer(":9091", "/kitexHelloclient")),

		client.WithErrorHandler(func(err error) error {
			//for thrift、KitexProtobuf
			if e, ok := err.(*remote.TransError); ok {
				return e
				//return e.TypeID()
			}
			//for gRPC
			/*if s, ok := status.FromError(err); ok {
				return s
				//return s.Code(),s
			}*/
			return kerrors.ErrRemoteOrNetwork.WithCause(err)
		}),
		//client.WithLoadBalancer(loadbalance.NewWeightedBalancer()),
		//指定这个服务的ip和端口
		//client.withHostports(":8888"),
		/**
		传输协议
		通常 RPC 协议中包含 RPC 消息协议和应用层传输协议，RPC 消息协议看做是传输消息的 Payload，传输协议额外传递一些元信息通常会用于服务治理，框架的 MetaHandler 也是和传输协议搭配使用。在微服务场景下，传输协议起到了重要的作用，如链路跟踪的透传信息通常由传输协议进行链路传递。

		Kitex 目前支持两种传输协议：TTHeader、HTTP2，但实际提供配置的 Transport Protocol 是：TTHeader、GRPC、Framed、TTHeaderFramed、PurePayload。

		这里做一些说明：

		因为 Kitex 对 Protobuf 的支持有 Kitex Protobuf 和 gRPC
		为方便区分将 gRPC 作为传输协议的分类，框架会根据是否有配置 gRPC 决定使用哪个协议：Kitex Protobuf还是gRPC
		Framed 严格意义上并不是传输协议，只是标记 Payload Size 额外增加的 4 字节头，但消息协议对是否有 Framed 头并不是强制的，PurePayload 即没有任何头部的，所以将 Framed 也作为传输协议的分类；
		Framed 和 TTHeader 也可以结合使用，所以有 TTHeaderFramed 。
		消息协议可选的传输协议组合如下：

		Thrift: TTHeader(建议)、Framed、TTHeaderFramed
		KitexProtobuf: TTHeader(建议)、Framed、TTHeaderFramed
		gRPC: HTTP2

		传输协议封装消息协议进行 RPC 互通，传输协议可以额外透传元信息，用于服务治理，Kitex 支持的传输协议有 TTHeader、HTTP2。
		TTHeader 可以和 Thrift、Kitex Protobuf 结合使用；HTTP2 目前主要是结合 gRPC 协议使用，后续也会支持 Thrift。
		*/
		//配置项
		//Client 初始化时通过 WithTransportProtocol 配置传输协议：
		//
		//// client option
		//client.WithTransportProtocol(transport.XXX)
		//Server 支持协议探测（在 Kitex 默认支持的协议内），无需配置传输协议
		//在创建客户端的时候需指定使用 gRPC 协议
		//目前 Kitex 支持的消息类型、编解码协议和传输协议
		//
		//消息类型	编码协议	传输协议
		//PingPong	Thrift / Protobuf	TTHeader / HTTP2(gRPC)
		//Oneway	Thrift	TTHeader
		//Streaming	Protobuf	HTTP2(gRPC)
		//PingPong：客户端发起一个请求后会等待一个响应才可以进行下一次请求
		//Oneway：客户端发起一个请求后不等待一个响应
		//Streaming：客户端发起一个或多个请求 , 等待一个或多个响应
		//连接复用不可用于GRPC(用于Streaming	Protobuf	HTTP2(gRPC))

		//所以只能是pb的IDL才能客户端开启grpc
		//client.WithTransportProtocol(transport.GRPC),
		//client.WithGRPCConnPoolSize(6), // the cpu cores of server is 4, and 4*3/2 = 6
		//client.WithRPCTimeout(3 * time.Second),              // rpc timeout

		//Thrift + TTHeader
		client.WithTransportProtocol(transport.TTHeaderFramed),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),

		// 10 MB
		client.WithCodec(codec.NewDefaultCodecWithSizeLimit(1024 * 1024 * 10)),

		//client.WithLongConnection(
		//	connpool.IdleConfig{MaxIdlePerAddress: 1000, MaxIdleGlobal: 1000, MaxIdleTimeout: time.Minute}),

		// client.WithMuxConnection(opt.PoolSize))
		// netpoll 设计上，2 个 connection 最优
		//这里的连接多路复用是针对于 Thrift 和 Kitex Protobuf，如果配置 gRPC 协议(client.WithTransportProtocol(transport.GRPC))，默认是连接多路复用。
		//Client 开启连接多路复用，Server 必须也开启，否则会导致请求超时；Server 开启连接多路复用对 Client 没有限制，可以接受短连接、长连接池、连接多路复用的请求。
		client.WithMuxConnection(2),
		client.WithConnectTimeout(20 * time.Millisecond),//20ms

		client.WithFailureRetry(fp),
		client.WithResolver(r1),
		//client.WithResolver(dns.NewDNSResolver()),
		//client.WithBackupRequest(bp),
		//client.WithMiddleware(failure.NewFailureMW()),
		//client.WithMiddleware(failure.NewDelayMW(60 * time.Millisecond)),//毫秒
		client.WithSuite(trace.NewDefaultClientSuite()),
		//client.WithBoundHandler(bound.NewCpuLimitHandler()),

	)

	//基于K8s服务发现
	/**
	service-host为hosts域名配置：service-name.{namespace}.svc.{cluster.domain}
	hello-service-svc.default.svc.cluster.local
	c1, err := hello.NewClient (
		constants.HelloServiceName,
		client.WithLoadBalancer(lb),
		client.WithPayloadCodec(thrift.NewThriftCodecWithConfig(thrift.FastRead | thrift.FastWrite)),
		client.WithFailureRetry(fp),
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithHostPorts("hello-service-svc.default.svc.cluster.local:9000"),//此处为服务的svc的信息
		client.WithMuxConnection(2),                       // mux
		client.WithConnectTimeout(50 * time.Millisecond),    // conn timeout
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer

		//callopt.WithURL("http://myserverdomain.com:8888"),
		//client.WithHTTPResolver(myResolver),


		//client.WithTransportProtocol(transport.GRPC),
		//client.WithGRPCConnPoolSize(6), // the cpu cores of server is 4, and 4*3/2 = 6
		//client.WithRPCTimeout(5*time.Second),
	)*/

	//也可以直接通过RPC协议进行访问，但需要指定hello的服务ip和端口
	//c1, err := hello.NewClient(constants.HelloServiceName, client.WithHostPorts("${helloServiceIp}:2008"))//client.WithHostPorts("[::1]:2008"))

	if err != nil {
		klog.Error(err)
	}
	helloClient = c
}

func Echo(ctx context.Context, req *api.Request) (r *hello.Response, err error) {
	return helloClient.Echo(ctx, req)
}

func GetH(ctx context.Context, id int32) (r *hello.Response, err error) {
	return helloClient.GetH(ctx, id)
}
