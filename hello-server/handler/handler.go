package handler

import (
	"context"
	"primus/kitex_gen/hello"
)

// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct{}

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
	return
}

// Post implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) Post(ctx context.Context, req *hello.Request) (resp *hello.Response, err error) {
	// TODO: Your code here...
	return
}