package main

import (
	"fmt"
	"learn.go/zuoye/tizhi_register5/register"
	"sync"
	"time"
)

func main() {
	registerToAll := make(chan *register.Member, 1000)
	all := &register.Allmember{Members: []*register.Member{}}
	totalMember := 1000
	sMap := &sync.Map{}

	for i := 0; i < totalMember; i++ { // 注册1000个
		go register.Register(i, registerToAll) //注册1000个，并注入到channel中。
	}

	time.Sleep(200 * time.Millisecond) //给注入数据留点时间
	close(registerToAll)

	// 先close channel，然后再for range channel
	for omember := range registerToAll { // 通过range把channel中的数据，传入all.Members
		all.Members = append(all.Members, omember)
	}

	for i := 0; i < 3; i++ { //跟新3轮数据
		fmt.Printf("**第%d轮开始!**********************\n", i+1)
		go all.MakeFatR4(totalMember) //尝试fatr和rank分离 这个只有fatr
		//time.Sleep(2 * time.Second)

		all.MakeRank() //这个是排名

		//for i, v := range all.Members { //测试通过
		//	fmt.Println(i, v)
		//}

		all.GetMap(sMap, totalMember) //构建一个map key结构体，目的是要用名字，v是排名

		for _, key := range all.Members {
			v, ok := sMap.Load(key.Name)
			fmt.Printf(" %s的排名是：%d。 %teacher\n", key.Name, v, ok)
		}
		time.Sleep(1 * time.Second)
	}
}
