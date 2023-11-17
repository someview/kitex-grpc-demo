package main

import (
	"context"
	"fmt"
	"log"
	"net"
	. "pbtest"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/server"
)

type App struct{}

// OnPrivateMessage implements TL_GoPrivate_Contract.GoPrivateService.
func (*App) OnPrivateMessage(stream GoPrivateService_OnPrivateMessageServer) (err error) {
	return fmt.Errorf("now not supported")
}

// PostMessage implements TL_GoPrivate_Contract.GoPrivateService.
func (*App) PostMessage(ctx context.Context, req *PrivateRequest) (resp *PrivateResponse, err error) {
	fmt.Println("App req:", req)
	return &PrivateResponse{}, nil
}

var _ GoPrivateService = (*App)(nil)

func echoServer() {
	app := &App{}
	addr, err := net.ResolveTCPAddr("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	server := NewGoPrivateServiceServer(app, server.WithServiceAddr(addr))
	if err := server.Run(); err != nil {
		log.Fatalln("err:", err)
	}
}

func echoClient() {
	client, err := NewGoPrivateServiceClient("GoPrivateService", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := client.PostMessage(ctx, &PrivateRequest{})
	fmt.Println("client, res:", res, "err:", err)
}

func main() {
	go echoServer()
	echoClient()
	time.Sleep(time.Second * 10)
}
