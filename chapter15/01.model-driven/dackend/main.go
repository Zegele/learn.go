// backend 后端创建
package main

import (
	"context"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"learn.go/chapter15/01.model-driven/models"
	"learn.go/pkg/dockertool"
	"log"
)

// 老师是依托etcd做的，视频1：47 启动了etcd ，我的没有
// 创建了一个model-driven的文件夹，并进入该文件夹，然后运行了etcd命令
// etcd  // 运行了etcd的一个节点

func main() {
	backendName := "/rankservice/backend"
	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{"http://localhost:2379"}})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background() //background 一般用于后端  todo用于其他

	dockerCli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		watcher := cli.Watch(ctx, backendName)
		for respData := range watcher {
			evs := respData.Events
			for _, ev := range evs {
				rawBackend := ev.Kv.Value

				backend := &models.RankServiceBackend{}
				json.Unmarshal(rawBackend, backend) // todo handle error
				if backend.Expected.Count != backend.Status.RunningCount {
					//todo 调用docker创建新的容器实例，或 删除一些，并更新status
					dockertool.Run(ctx)
				} else {
					fmt.Println("已经满足预期", backend.Expected.Count)
				}
			}
		}
	}()
	<-ctx.Done()
}
