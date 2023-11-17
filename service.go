package TL_GoPrivate_Contract

import (
	"context"

	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
)

// GoPrivateService is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type GoPrivateService interface {
	OnPrivateMessage(stream GoPrivateService_OnPrivateMessageServer) (err error)
	PostMessage(ctx context.Context, req *PrivateRequest) (resp *PrivateResponse, err error)
}

type GoPrivateService_OnPrivateMessageServer interface {
	streaming.Stream
	Recv() (*PrivateRequest, error)
	Send(*PrivateResponse) error
}

var GoPrivateServiceServiceInfo = NewGoPrivateServiceServiceInfo()

func NewGoPrivateServiceServiceInfo() *kitex.ServiceInfo {
	serviceName := "GoPrivateService"
	handlerType := (GoPrivateService)(nil)
	methods := map[string]kitex.MethodInfo{
		"OnPrivateMessage": kitex.NewMethodInfo(OnPrivateMessageHandler, func() any { return new(PrivateRequest) }, func() any { return new(PrivateResponse) }, false),
		"PostMessage":      kitex.NewMethodInfo(PostMessageHandler, func() any { return new(PrivateRequest) }, func() any { return new(PrivateResponse) }, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "TL_GoPrivate_Contract",
		"ServiceFilePath": "",
	}
	extra["streaming"] = true
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.7.3",
		Extra:           extra,
	}
	return svcInfo
}

// stream IDL
type GoPrivateServiceOnPrivateMessageServer struct {
	streaming.Stream
}

func (c *GoPrivateServiceOnPrivateMessageServer) Recv() (*PrivateRequest, error) {
	m := new(PrivateRequest)
	return m, c.Stream.RecvMsg(m)
}

func (c *GoPrivateServiceOnPrivateMessageServer) Send(res *PrivateResponse) error {
	return c.Stream.SendMsg(res)
}

func OnPrivateMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	st := arg.(*streaming.Args).Stream
	stream := &GoPrivateServiceOnPrivateMessageServer{st}
	return handler.(GoPrivateService).OnPrivateMessage(stream)
}
func PostMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(PrivateRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(GoPrivateService).PostMessage(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PrivateRequest:
		req := arg.(*PrivateRequest)
		resultPtr, ok := result.(*PrivateResponse)
		if !ok {
			panic("generator code not compatiable")
		}
		res, err := handler.(GoPrivateService).PostMessage(ctx, req)
		if err != nil {
			return err
		}
		*resultPtr = *res // how to avoid alloc buffer
	}
	return nil
}
