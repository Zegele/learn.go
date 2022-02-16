package main

import (
	"fmt"
	"learn.go/zuoye/tizhi_register3/register"
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

	go all.MakeFatR4(totalMember) //尝试fatr和rank分离
	//time.Sleep(5 * time.Second)

	all.MakeRank()

	for i, v := range all.Members { //打印1000个注册数据。测试通过！
		fmt.Println(i, v)
	}
	//for i := 0; i < totalMember; i++ {
	//	go all.MakeFatRAndRank("a" + fmt.Sprint(i)) //生成 fatr 和 rank
	//	//time.Sleep(1 * time.Millisecond)
	//}

	all.GetMap(sMap, totalMember)
	//for i := 0; i < totalMember; i++ {
	//	go func(i int) {
	//		v, ok := sMap.Load(all.Members[i].Name)
	//		fmt.Println(v, ok)
	//	}(i)
	//}

	//time.Sleep(5 * time.Second)
	for i, key := range all.Members {
		v, ok := sMap.Load(key.Name)
		fmt.Printf("%d: %s的体脂率是：%f; 排名是：%d。 %t\n", i, key.Name, key.FatR, v, ok)
	}
}
