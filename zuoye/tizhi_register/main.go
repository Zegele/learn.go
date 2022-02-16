package main

import (
	"fmt"
	"time"
)

//注册1000人的信息
// 核4个 4个人注册 把注册信息装到channel中，然后把channel的值传入 allmember切片
// channel

type member struct {
	name string
	fatR float64
	rank uint8
}

type allmember struct {
	members []*member //为什么要用指针类型，因为member中的值可能会变化，所以这里用指针的值。
}

func Register(number int) (m *member) { //第一步：把数据注入channel
	m = &member{} //需要实例化一下
	m.name = "a" + fmt.Sprint(number)
	//fmt.Println("?", m) // 在多线程中，函数内部尽量不要打印，数据呈现可能会有问题
	return
}

func (all *allmember) AddMember(c chan *member) {
	//all = &allmember{members: []*member{}}
	for omember := range c {
		//	fmt.Println(omember)
		all.members = append(all.members, omember)
	}
	//	return
}

func main() {
	registerToAll := make(chan *member, 1000)

	for i := 0; i < 1000; i++ { // 注册10个
		//fmt.Println(i)
		registerToAll <- Register(i) //注册10个，并注入到channel中。
	}
	//time.Sleep(1 * time.Second)
	close(registerToAll) //关闭 注册channel

	//for o := range registerToAll { //测试 通道里有东西。记得注释掉，否则通道是空的
	//	fmt.Println("??", o)
	//}
	//---------------以上通过

	all := &allmember{members: []*member{}}
	//for i := 0; i < 4; i++ { // 为什么带来 for 最终的数据是有问题的？？？？
	//	go all.AddMember(registerToAll) //从channel中读取值，录入到allmember中
	//}
	go all.AddMember(registerToAll) //for  的意义是什么？
	time.Sleep(1 * time.Second)

	for i, v := range all.members { // 单个go all.AddMember(registerToAll) 结果正确，但是for 多个goroutine后，结果就会出错。
		fmt.Println(i, v)
	}

	//time.Sleep(1 * time.Second)
}

//func main() {
//	registerToAll := make(chan *member, 1000) // 把注册信息传入allMembers的channel
//	//wg := sync.WaitGroup{}
//	//heshu := 4
//	//wg.Add(heshu)
//	var all allmember
//	//for i := 0; i < heshu; i++ {
//	//	go all.AddMember(registerToAll)
//	//}
//
//	for i := 0; i < 10; i++ {
//		registerToAll <- Register(i)
//	}
//	close(registerToAll)
//	//	wg.Wait()
//
//	for i, v := range all.members {
//		fmt.Println(i, v)
//	}
//
//	//time.Sleep(10 * time.Second)
//
//}

//func main() {
//	//var all allmember // 只是声明，没有实例化
//	all := &allmember{members: []*member{}} // 声明，并实例化
//
//	for i := 0; i < 3; i++ {
//		all.members = append(all.members, &member{name: "a" + fmt.Sprint(i)})
//		fmt.Println(all.members[i])
//	}
//	fmt.Println(all)
//}
