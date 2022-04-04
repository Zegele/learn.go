package main

import (
	"context"
	"google.golang.org/grpc"
	"learn.go/chapter13/02.grpc/apis"
	"log"
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
}
