package main

// 客户端单词请求，服务端单次答复。单去单回
import (
	context2 "golang.org/x/net/context"
	"google.golang.org/grpc"
	"learn.go/chapter13/02.grpc/apis"
	"learn.go/chapter13/02.grpc/server/rankserver"
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
		// 不需要考虑url，关注函数的参数和返回值，参数和返回值都被定义在payload中。
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
