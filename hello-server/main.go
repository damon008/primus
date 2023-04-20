package main

import (
	"log"
	"primus/hello-server/handler"
	hello "primus/kitex_gen/hello/helloservice"
)

func main() {
	svr := hello.NewServer(new(handler.HelloServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
