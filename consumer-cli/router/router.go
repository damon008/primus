package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"primus/consumer-cli/api"
)

func Init(r *server.Hertz) {
	v1 := r.Group("/api/v1") //, license.LicenceIssued()...)

	hello := v1.Group("/hello")
	hello.GET("", api.GetH)
}
