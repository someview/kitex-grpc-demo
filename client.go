// Code generated by Kitex. DO NOT EDIT.

package TL_GoPrivate_Contract

import (
	"context"
	"fmt"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	transport "github.com/cloudwego/kitex/transport"
)



type GoPrivateServiceClient interface {
	OnPrivateMessage(ctx context.Context , callOptions ...callopt.Option) (stream GoPrivateService_OnPrivateMessageClient, err error)
	PostMessage(ctx context.Context, req *PrivateRequest, callOptions ...callopt.Option) (resp *PrivateResponse, err error)
}

type GoPrivateService_OnPrivateMessageClient interface {
	streaming.Stream
	Send(*PrivateRequest) error
	Recv() (*PrivateResponse, error)
}  


// KGoPrivateServiceOnPrivateMessageClient is the client implementation for the stream method OnPrivateMessage of GoPrivateService.
type KGoPrivateServiceOnPrivateMessageClient struct {
	streaming.Stream
}
func (c *KGoPrivateServiceOnPrivateMessageClient) Send(req *PrivateRequest) error {
	return c.Stream.SendMsg(req)
}
func (c *KGoPrivateServiceOnPrivateMessageClient) Recv() (*PrivateResponse, error) {
	m := new(PrivateResponse)
	return m, c.Stream.RecvMsg(m)
}

// kGoPrivateServiceClient is an implementation of the GoPrivateServiceClient interface.
type KGoPrivateServiceClient struct {
	client.Client // The actual Kitex client
}


func (c *KGoPrivateServiceClient) OnPrivateMessage(ctx context.Context , callOptions ...callopt.Option) (stream GoPrivateService_OnPrivateMessageClient, err error) {
	streamClient, ok := c.Client.(client.Streaming)
	if !ok {
		return nil, fmt.Errorf("client not support streaming")
	}
	res := new(streaming.Result)
	err = streamClient.Stream(ctx, "OnPrivateMessage", nil, res)
	if err != nil {
		return nil, err
	}
	stream = &KGoPrivateServiceOnPrivateMessageClient{res.Stream}
	return stream, nil
}

func (c *KGoPrivateServiceClient) PostMessage(ctx context.Context, req *PrivateRequest, callOptions ...callopt.Option) (resp *PrivateResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	err = c.Client.Call(ctx, "PostMessage", req, resp)
	return
}

// NewGoPrivateServiceClient creates a client for the service defined in IDL.
func NewGoPrivateServiceClient(destService string, opts ...client.Option) (GoPrivateServiceClient, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))
	options = append(options, client.WithTransportProtocol(transport.GRPC))
	options = append(options, opts...)
	kc, err := client.NewClient(NewGoPrivateServiceServiceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &KGoPrivateServiceClient{
		Client: kc,
	}, nil
}