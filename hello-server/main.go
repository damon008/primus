package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	"primus/hello-server/handler"
	hello "primus/kitex_gen/hello/helloservice"
)

func main() {
	svr := hello.NewServer(
		new(handler.HelloServiceImpl),
		server.WithServiceAddr(&net.TCPAddr {
			Port: 1000,
		}),
		//server.WithListener(),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
