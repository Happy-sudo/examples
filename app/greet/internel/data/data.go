package data

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/wire"
	"hello/kitex_gen/hello/v1/hello"
)

var ProviderSet = wire.NewSet(NewData, NewGreetRepo)

type Data struct {
	hello hello.Client
}

func NewData(log klog.CtxLogger, hello hello.Client) (*Data, func(), error) {

	d := &Data{
		hello: hello,
	}
	return d, func() {
		log.CtxInfof(context.Background(), "closing the data resources")
	}, nil
}
