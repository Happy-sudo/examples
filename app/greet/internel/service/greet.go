package service

import (
	"context"
	v1 "greet/kitex_gen/greet/v1"
)

func (h *GreetService) Greet(ctx context.Context, req *v1.Request) (r *v1.Response, err error) {
	h.log.CtxInfof(ctx, "Greet called : %s", req)
	getMessage, err := h.biz.Greet(ctx, req.GetMessage())
	if err != nil {
		return nil, err
	}
	return &v1.Response{
		Message: getMessage,
	}, nil
}
