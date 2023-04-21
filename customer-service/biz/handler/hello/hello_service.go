// Code generated by hertz generator.

package hello

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	hello "primus/customer-service/biz/model/hello"
)

// GetByParams .
// @router /api/v1/getByParams/:data [GET]
func GetByParams(ctx context.Context, c *app.RequestContext) {
	var err error
	var req string
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(hello.Response)

	c.JSON(consts.StatusOK, resp)
}

// GetH .
// @router /api/v1/getH/:id [GET]
func GetH(ctx context.Context, c *app.RequestContext) {
	//var err error
	//var req int32
	hlog.Info("app: ", c)
	id := c.Param("id")
	hlog.Info("id: ", id)
	/*err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}*/

	resp := new(hello.Response)
	resp.Data = id
	c.JSON(consts.StatusOK, resp)
}

// Post .
// @router /api/v1/create [POST]
func Post(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hello.Request
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(hello.Response)

	c.JSON(consts.StatusOK, resp)
}