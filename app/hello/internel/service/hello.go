package service

import (
	"context"
	v1 "hello/kitex_gen/hello/v1"
)

func (h *HelloService) Hello(ctx context.Context, req *v1.Request) (r *v1.Response, err error) {
	h.log.CtxInfof(ctx, "Hello called : %s", req)
	getMessage, err := h.biz.Hello(ctx, req.GetMessage())
	if err != nil {
		return nil, err
	}
	return &v1.Response{
		Message: getMessage,
	}, nil
}
