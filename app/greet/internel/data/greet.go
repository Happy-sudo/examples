package data

import (
	"context"
	"greet/internel/biz"
	v1 "hello/kitex_gen/hello/v1"
)

type GreetRepo struct {
	data *Data
}

func (h *GreetRepo) Greet(ctx context.Context, message string) (string, error) {
	hello, err := h.data.hello.Hello(ctx, &v1.Request{
		Message: message,
	})
	if err != nil {
		return "", err
	}

	return hello.GetMessage(), nil
}

func NewGreetRepo(data *Data) biz.GreetRepo {
	return &GreetRepo{data: data}
}
