package main

import (
	"fmt"
	"math"
)

func main() {
	var hello string = "hello, golang!"
	var world = "world"
	fmt.Println(hello, world)
	小数 := 1.234
	fmt.Println(小数)

	var int3, int4 uint = 33, 44
	fmt.Println(int3 * int4)

	ho, ver := 3, 4.123
	var sc = ho * int(ver)
	fmt.Println(ho * int(ver))
	fmt.Println(sc)
	var newname string
	fmt.Println("newname = ", newname)

	var int6 uint = math.MaxUint64
	fmt.Println(int6)
	var int7 int = int(int6) //溢出
	fmt.Println(int7)        //-1 奇怪的结果

	var int8 int = math.MaxInt
	fmt.Println(int8)

	//var nameOfSquare string //驼峰命名
	//var _name string // 小组内部变量 _name

}
