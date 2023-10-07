// Code generated by Kitex v0.4.3. DO NOT EDIT.

package xxx

import (
	server "github.com/cloudwego/kitex/server"
	v1 "xxx/kitex_gen/xxx/v1"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler v1.XXX, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}