package main

import (
	"context"
	"google.golang.org/grpc"
	"learn.go/zuoye/bigzuoye/api"
	"learn.go/zuoye/bigzuoye/server/client_to_server"
	"log"
	"net"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	startGRPCServer(ctx) //联通服务器，测试通过

}

// grpc
func startGRPCServer(ctx context.Context) {
	lis, err := net.Listen("tcp", "0.0.0.0:9090") //建立一个监听
	//注意监听地址要加引号'
	//"localhost:9090" 和 0.0.0.0：9090 是一样的效果
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer([]grpc.ServerOption{}...) //新建并初始化了一个服务
	api.RegisterClient_ServiceServer(s, &client_to_server.CToSServer{
		//自动生成的RegisterClient_ServiceServer，用于注册服务，它的第二个参数是要注册的server接口的实现
		// client_to_server.CToSServer是实现了该结构的结构体
		Persons: map[string]*api.PersonalInformation{},
	})
	//以上是设置要服务的内容

	go func() {
		select {
		case <-ctx.Done():
			s.Stop()
		}
	}()

	if err := s.Serve(lis); err != nil { //等待服务
		// 对s对象，进行Serve服务，服务连接是0.0.0.0：9090
		log.Fatalf("failed to serve: %v\n", err)
	}

}
