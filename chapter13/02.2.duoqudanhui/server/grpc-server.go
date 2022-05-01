package main

// 实现客户端多次发送请求，服务端一次答复
// 客户端有一个stream client 不停地发送数据， 并告知“说完了”，然后接收消息
import (
	context2 "golang.org/x/net/context"
	"google.golang.org/grpc"
	"learn.go/chapter13/02.2.duoqudanhui/apis"
	"learn.go/chapter13/02.2.duoqudanhui/server/rankserver"
	"log"
	"net"
)

func main() {
	ctx, cancel := context2.WithCancel(context2.TODO())
	defer cancel()

	startGRPCServer(ctx)
}

func startGRPCServer(ctx context2.Context) {
	lis, err := net.Listen("tcp", "0.0.0.0:9090") //第二个参数是nodePort端口 0.0.0.0 就是localhost？
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer([]grpc.ServerOption{}...)             // 这是啥？？？
	apis.RegisterRankServiceServer(s, &rankserver.RankServer{ // RegisterRankServiceServer 注册这个接口  pb.go
		Persons: map[string]*apis.PersonalInformation{},
	})
	go func() {
		select {
		case <-ctx.Done(): //如果这个监听服务的context Done了，就停止。
			s.Stop()
		}
	}()
	if err := s.Serve(lis); err != nil { // s.Serve 是启动该服务器。
		log.Fatalf("failed to serve: %v", err)
	}
	//这种也是单次请求，单次答复。
}
