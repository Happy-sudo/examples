package data

import (
	"context"
	"hello/internal/biz"
)

type HelloRepo struct {
	data *Data
}

func NewHelloRepo(data *Data) biz.HelloRepo {
	return &HelloRepo{data: data}
}

func (h *HelloRepo) Hello(ctx context.Context, message string) (string, error) {
	return message, nil
}
