package register

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

//注册1000人的信息
// 核4个 4个人注册 把注册信息装到channel中，然后把channel的值传入 allmember切片
// channel

type Member struct {
	Name string
	FatR float64
	Rank int
}

type Allmember struct {
	Members []*Member //为什么要用指针类型，因为member中的值可能会变化，所以这里用指针的值。
}

func Register(number int, c chan *Member) (m *Member) { //第一步：把数据注入channel
	m = &Member{} //需要实例化一下
	m.Name = "a" + fmt.Sprint(number)
	c <- m
	//fmt.Println("?", m) // 在多线程中，函数内部尽量不要打印，数据呈现可能会有问题
	return
}

// 注意： 函数体中有 for range channel ，调用该函数时（go routine时），结果可能会有错。
func (all *Allmember) AddMember(c chan *Member) { //函数中套for range channel 会使数据出错。
	//all = &allmember{members: []*member{}}
	for omember := range c {
		//	fmt.Println(omember)
		all.Members = append(all.Members, omember)
	}
	//	return
}

func (allm *Allmember) MakeFatRAndRank(name string) {
	//makeFatr

	rand.Seed(time.Now().UnixNano()) //参数类型是int64
	r := rand.Intn(40)               //范围是0-39 不包括40
	rToF := float64(r)/100 + 0.01    //加1是避免随机出现0的情况。 处理之后的范围是：0.01-0.4
	//time.Sleep(1 * time.Millisecond)

	//参考chapter04-05
	for _, v := range allm.Members {
		if v.Name == name {
			v.FatR = rToF // 赋值给fatR
			break         //break 放的位置一定要仔细考虑
		}
		// break 放这就循环一次。。。找了好久，没看到这个错误。
	}

	//makeRank
	sort.Slice(allm.Members, func(i, j int) bool { //每次赋值FatR后，都要排序一次
		return allm.Members[i].FatR < allm.Members[j].FatR
	})

	for i, _ := range allm.Members {
		allm.Members[i].Rank = i + 1
	}
}

func (allm *Allmember) makeFatR3(i int) {
	//makeFatr

	rand.Seed(time.Now().UnixNano())             //参数类型是int64
	r := rand.Intn(40)                           //范围是0-39 不包括40
	allm.Members[i].FatR = float64(r)/100 + 0.01 //加1是避免随机出现0的情况。 处理之后的范围是：0.01-0.4

	//makeRank
	//sort.Slice(allm.Members, func(i, j int) bool { //每次赋值FatR后，都要排序一次
	//	return allm.Members[i].FatR < allm.Members[j].FatR
	//})
	//
	//for i, _ := range allm.Members {
	//	allm.Members[i].Rank = i + 1
	//}
}

func makeFatR(allm *Allmember, i int) {
	//makeFatr
	rand.Seed(time.Now().UnixNano()) //参数类型是int64
	r := rand.Intn(39)
	rToF := float64(r)/100 + 0.01 //加1是避免随机出现0的情况。
	//time.Sleep(1 * time.Millisecond)
	allm.Members[i].FatR = rToF

	//makeRank
	//sort.Slice(allm.Members, func(i, j int) bool { //每次赋值FatR后，都要排序一次
	//	return allm.Members[i].FatR < allm.Members[j].FatR
	//})

	//for i, _ := range allm.Members {
	//	allm.Members[i].Rank = i + 1
	//}
}

func (allm Allmember) GetMap(sMap *sync.Map, memNum int) {
	for i := 0; i < memNum; i++ {
		go func(i int) {
			sMap.Store(allm.Members[i].Name, allm.Members[i].Rank)
		}(i)
	}
	//time.Sleep(2 * time.Second)
}

//func (allm Allmember) GetMap2(sMap *sync.Map, memNum int) {
//	go func(i int) {
//		sMap.Store(allm.Members[i].Name, allm.Members[i].Rank)
//		v, ok := sMap.Load("a" + fmt.Sprint(i))
//		fmt.Printf("%s的体脂率是：%f; 排名是：%d。 %teacher\n", allm.Members[i].Name, allm.Members[i].FatR, v, ok)
//	}(memNum)
//}
