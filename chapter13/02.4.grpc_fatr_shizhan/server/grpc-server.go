package main

import (
	context2 "golang.org/x/net/context"
	"google.golang.org/grpc"
	"learn.go/chapter13/02.4.grpc_fatr_shizhan/apis"
	"learn.go/chapter13/02.4.grpc_fatr_shizhan/ranks"
	"learn.go/chapter13/02.4.grpc_fatr_shizhan/server/rankserver"
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
		RankS:    ranks.NewFatRateRank(),
		PersonCh: make(chan *apis.PersonalInformation, 1024),
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
