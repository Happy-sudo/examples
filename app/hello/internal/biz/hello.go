package biz

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type HelloRepo interface {
	Hello(ctx context.Context, message string) (string, error)
}

type HelloUseCase struct {
	repo HelloRepo
	log  klog.CtxLogger
}

func NewHelloUseCase(repo HelloRepo, log klog.CtxLogger) *HelloUseCase {
	return &HelloUseCase{repo: repo, log: log}
}
func (u *HelloUseCase) Hello(ctx context.Context, message string) (string, error) {
	return u.repo.Hello(ctx, message)
}
