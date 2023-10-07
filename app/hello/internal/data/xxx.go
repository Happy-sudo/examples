package data

import (
	"context"
	"xxx/internal/biz"
)

type XXXRepo struct {
	data *Data
}

func NewXXXRepo(data *Data) biz.XXXRepo {
	return &XXXRepo{data: data}
}

func (h *XXXRepo) XXX(ctx context.Context, message string) (string, error) {
	return message, nil
}
