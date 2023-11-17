package TL_GoPrivate_Contract

import (
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewGoPrivateServiceServer(handler GoPrivateService, opts ...server.Option) server.Server {
	var options []server.Option
	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(NewGoPrivateServiceServiceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
