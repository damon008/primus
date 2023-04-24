package handler

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"primus/kitex_gen/hello"
)

// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct{}

func (s *HelloServiceImpl) GetByParams(ctx context.Context, data string, msg string) (r *hello.Response, err error) {
	//TODO implement me
	panic("implement me")
}

// Echo implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) Echo(ctx context.Context, req *hello.Request) (resp *hello.Response, err error) {
	// TODO: Your code here...

	return
}

// Get implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) Get(ctx context.Context) (resp *hello.Response, err error) {
	// TODO: Your code here...
	return
}

// GetH implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) GetH(ctx context.Context, id int32) (resp *hello.Response, err error) {
	// TODO: Your code here...
	klog.Info("param: ", id)

	return &hello.Response{
		Msg: &hello.Msg{
			Code: 0,
			Msg:  "success",
		},
		Data: fmt.Sprintf("result is: %d", id),
	}, nil
}

// Post implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) Post(ctx context.Context, req *hello.Request) (resp *hello.Response, err error) {
	// TODO: Your code here...
	return
}
