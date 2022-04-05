package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"learn.go/chapter13/02.3.daiwatch/apis"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure()) //这里WithInsecure是必需的  grpc是基于https的，必须是以https访问的 ，
	// 如果强制不用https，就必须要WithInsecure
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := apis.NewRankServiceClient(conn)
	w, err := c.WatchPersons(context.TODO(), &apis.Null{})
	if err != nil {
		log.Fatal("启动watcher失败：", err)
	}
	for {
		pi, err := w.Recv() //接收 ， 在for循环里不停接收
		if err != nil {
			if err == io.EOF {
				log.Println("服务器告知说完了")
				break
			}
			log.Fatal("接收异常:", err)
		}
		log.Println("收到新变动：", pi.String())
	}
	//然后应该是给个end ，但这里没有给end，让一直跑。一直看有没有变化？
}
