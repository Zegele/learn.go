package main

import "fmt"

func main() {
	hello()

	helloToSomeone("aa")
	helloToSomeone("bb")

	msg := constructHelloMessage("cc")
	fmt.Println(msg)

	bmi := calcBMI(65, 1.7)
	fmt.Println(bmi)
}

func hello() { //无参数无返回值
	fmt.Println("你好，golang！")
}

func helloToSomeone(name string) { //有参数无返回值
	fmt.Println("你好，", name)
}

func constructHelloMessage(name string) string { //有参数有返回值
	//fmt.Println("你好，", name)
	return "你好，" + name + "，再来一个"
}

func calcBMI(weight, tall float64) float64 {
	return weight / (tall * tall)
}
