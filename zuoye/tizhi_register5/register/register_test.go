package register

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	//memberNum := 100
	registerToAll := make(chan *Member, 1000) // 当这个buffer设置较小时，会出错。可能的原因时，空间小，1000个数据还没有导入完毕，channel就关闭了。所以会让数据丢失。

	all := &Allmember{Members: []*Member{}}

	//for i := 0; i < 1000; i++ { // 为什么带来 for 最终的数据是有问题的？？？？
	// 可能的原因，for循环和函数体内的channel for循环，可能有不同routine相互影响。
	//	go all.AddMember(registerToAll) //从channel中读取值，录入到allmember中
	//}

	// 注册1000个， 并把1000个数据导入到channel中
	for i := 0; i < 1000; i++ { // 注册1000个
		go Register(i, registerToAll) //注册10个，并注入到channel中。
	}

	time.Sleep(200 * time.Millisecond) //给注入数据留点时间
	close(registerToAll)               //关闭 注册channel ，关闭的channel才能正常遍历

	// 先close channel，然后再for range channel
	for omember := range registerToAll { // 通过range把channel中的数据，传入all.Members
		//如果把该for range 包装到函数体内，调用函数使用该for range时，会出错。
		all.Members = append(all.Members, omember)
	}
	// 把拿数据的放在这里，当channel中没有数据的时候，会等待，直到channel中有数据，它才运行。

	//for o := range registerToAll { //测试 通道里有东西。记得注释掉，否则通道是空的
	//	fmt.Println("??", o)
	//}
	//---------------以上通过

	for i, v := range all.Members { //打印1000个注册数据。测试通过！
		fmt.Println(i, v)
	}
}

func TestFarTAndRank(t *testing.T) {
	m := &Member{Name: "a0", Rank: 1}
	m1 := &Member{Name: "a", Rank: 0}
	allm := &Allmember{Members: []*Member{m, m1}}

	fmt.Println(*allm.Members[0], *allm.Members[1])

	allm.MakeFatRAndRank(m.Name)
	time.Sleep(1 * time.Millisecond)
	allm.MakeFatRAndRank(m1.Name)

	fmt.Println(*allm.Members[0], *allm.Members[1])
	// 通过 测试 写入FatR

}

func TestRegisterToRank(t *testing.T) {
	registerToAll := make(chan *Member, 1000)
	all := &Allmember{Members: []*Member{}}

	for i := 0; i < 1000; i++ { // 注册1000个
		go Register(i, registerToAll)
	}

	time.Sleep(200 * time.Millisecond)
	close(registerToAll)

	for omember := range registerToAll {
		all.Members = append(all.Members, omember)
	}

	//for i, _ := range all.Members {
	//	go all.makeFatR3(i)
	//	//time.Sleep(1 * time.Millisecond)
	//}

	for i := 0; i < 1000; i++ {
		go all.MakeFatRAndRank("a" + fmt.Sprint(i))
		time.Sleep(1 * time.Millisecond) // 如果不sleep， goroutine 会出错。 如果不sleep，且不go关键字，不出错。 如果sleep，且带go 不出错，但是程序运行慢。没理解哪里的问题。
	}

	for i, v := range all.Members { //打印1000个注册名，体脂，排名。测试通过！
		fmt.Println(i, v)
	}
}

func TestFromToRank(t *testing.T) {
	m := &Member{Name: "a0", Rank: 1}
	m1 := &Member{Name: "a1", Rank: 2}
	allm := &Allmember{Members: []*Member{m, m1}}
	sMap := &sync.Map{}
	allm.GetMap(sMap, 2)
	for i := 0; i < 2; i++ {
		v, ok := sMap.Load(allm.Members[i].Name)
		fmt.Printf("%s的排名是：%d。 %teacher\n", allm.Members[i].Name, v, ok)
	}
}
