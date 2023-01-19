package data

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewData, NewHelloRepo)

type Data struct {
}

func NewData(log klog.CtxLogger) (*Data, func(), error) {
	d := &Data{}
	return d, func() {
		log.CtxInfof(context.Background(), "closing the data resources")
	}, nil
}
