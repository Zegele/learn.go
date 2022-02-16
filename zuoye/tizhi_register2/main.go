package main

import (
	"fmt"
	"learn.go/zuoye/tizhi_register2/register"
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

	for i := 0; i < totalMember; i++ {
		go all.MakeFatRAndRank("a" + fmt.Sprint(i)) //生成 fatr 和 rank
		//time.Sleep(1 * time.Millisecond)
	}

	all.GetMap(sMap, totalMember)
	for i := 0; i < totalMember; i++ {
		v, ok := sMap.Load("a" + fmt.Sprint(i))
		fmt.Printf("%s的体脂率是：%f; 排名是：%d。 %t\n", all.Members[i].Name, all.Members[i].FatR, v, ok)
	}
}
