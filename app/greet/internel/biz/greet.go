package biz

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GreetRepo interface {
	Greet(ctx context.Context, message string) (string, error)
}

type GreetUseCase struct {
	repo GreetRepo
	log  klog.CtxLogger
}

func NewGreetUseCase(repo GreetRepo, log klog.CtxLogger) *GreetUseCase {
	return &GreetUseCase{repo: repo, log: log}
}
func (u *GreetUseCase) Greet(ctx context.Context, message string) (string, error) {
	return u.repo.Greet(ctx, message)
}
