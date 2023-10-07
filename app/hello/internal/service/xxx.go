package service

import (
	"context"
	v1 "xxx/kitex_gen/xxx/v1"
)

func (h *XXXService) XXX(ctx context.Context, req *v1.Request) (r *v1.Response, err error) {
	h.log.CtxInfof(ctx, "XXX called : %s", req)
	getMessage, err := h.biz.XXX(ctx, req.GetMessage())
	if err != nil {
		return nil, err
	}
	return &v1.Response{
		Message: getMessage,
	}, nil
}
