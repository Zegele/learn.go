package main

import "fmt"

func main() {
	var left, right int = 1, 2
	// var op string = "+"

	c := Calculator{ // Calculator是结构体
		left:  left,
		right: right,
	}
	fmt.Printf("c的地址是 =%p\n", &c)      //c的地址是 =0xc000070480
	fmt.Println(c.Add())               //c的地址是=0xc0000704b0  地址是不一样的,\
	fmt.Println("c.result=", c.result) //c.result= 0

	newC := NewCalculator{}
	newC.left = 100
	newC.right = 200
	fmt.Println(newC.Add())

	mc := MyCommand{}
	mc.commandOptions["aa"] = "AAA"
	fmt.Println(mc.ToCmdStr())
}

type MyCommand struct {
	mainCommand    *string //这是什么意思？？？
	commandOptions map[string]string
}

func (my MyCommand) ToCmdStr() string {
	out := ""
	for k, v := range my.commandOptions {
		out = out + fmt.Sprintf("--%s=%s", k, v)
	}
	return out
}
