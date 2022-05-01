package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learn.go/chapter13/02.2.duoqudanhui/apis"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure()) //这里WithInsecure是必需的  grpc是基于https的，必须是以https访问的 ，
	// 如果强制不用https，就必须要WithInsecure
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := apis.NewRankServiceClient(conn)
	ret, err := c.Register(context.TODO(), &apis.PersonalInformation{Name: "Tom"})
	if err != nil {
		log.Fatal("注册失败：", err)
	}
	log.Println("注册成功", ret)

	log.Println("开始批量注册")
	regCli, err := c.RegisterPersons(context.TODO())
	//c.RegisterPersons 的返回值是一个接口。这个接口有Send(*PersonalInformation) error 和CloseAndRecv()方法
	if err != nil {
		log.Fatal("获取批量注册客户端失败：", err)
	}

	if err := regCli.Send(&apis.PersonalInformation{Name: fmt.Sprintf("tom-%d", time.Now().Nanosecond())}); err != nil {
		log.Fatal("注册时失败：", err)
	}
	time.Sleep(1 * time.Second)

	if err := regCli.Send(&apis.PersonalInformation{Name: fmt.Sprintf("tom-%d", time.Now().Nanosecond())}); err != nil {
		log.Fatal("注册时失败：", err)
	}
	time.Sleep(1 * time.Second)

	if err := regCli.Send(&apis.PersonalInformation{Name: fmt.Sprintf("tom-%d", time.Now().Nanosecond())}); err != nil {
		log.Fatal("注册时失败：", err)
	}
	time.Sleep(1 * time.Second)

	if err := regCli.Send(&apis.PersonalInformation{Name: fmt.Sprintf("tom-%d", time.Now().Nanosecond())}); err != nil {
		log.Fatal("注册时失败：", err)
	}

	resp, err := regCli.CloseAndRecv() //客户端发完消息了，并等待服务端接收。
	// 应该是这里把结果都聚合在一起了。

	if err != nil {
		log.Fatal("无法接收结果：", err) //说明服务端接收失败
	}

	log.Println("批量注册结果：", resp.String())
}
