package biz

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type XXXRepo interface {
	XXX(ctx context.Context, message string) (string, error)
}

type XXXUseCase struct {
	repo XXXRepo
	log  klog.CtxLogger
}

func NewXXXUseCase(repo XXXRepo, log klog.CtxLogger) *XXXUseCase {
	return &XXXUseCase{repo: repo, log: log}
}
func (u *XXXUseCase) XXX(ctx context.Context, message string) (string, error) {
	return u.repo.XXX(ctx, message)
}
