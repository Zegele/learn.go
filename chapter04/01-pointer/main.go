package main

import "fmt"

//指针 &a, *a *int 区别这3个东西
func main() {
	a, b := 1, 2
	add(&a, &b) //&a 是a盒子， 不是a内部的值
	fmt.Println(a)

	c := &a // c 的类型是 *int， c指向a的盒子， *c就可以拿到a里边的东西。 3
	// c 有自己的地址，但c中储存的内容是a的地址。只要有a的地址，就可以找到a中的值。
	d := &c // d 的类型是 **int，d指向c的盒子，d 本身是指针，它存的东西，也是指针。
	fmt.Println("d=", d, "*d=", *d, "**d=", **d)

	m := map[string]string{} //定义，并初始化了map
	mp1 := &m                // mp1 的类型就是： *map[string]string
	fmt.Println(mp1)
	put(m)
	fmt.Println("*mp1=", *mp1)

	f1 := add // f1 = func(int, int)
	f1(&a, &b)
	fmt.Println("f1, add= ", a)
	f1p1 := &f1     // f1p1 = *func(int, int) f1p1是这个函数的地址。
	(*f1p1)(&a, &b) // f1p1(&a, &b) 错误； *f1p1(&a, &b) 错误； (*f1p1(&a, &b))错误。
	fmt.Println("f1p1, add = ", a)

	{
		var nothing *int
		// *nothing = 3 // 注意： 这里是没有指向任何东西的int指针
		fmt.Println("nothing=", nothing) //nothing= <nil>
	}

	{
		var nothingMap map[string]string //没有初始化的map  var nothingMap = map[string]string{}
		//nothingMap["a"] = "BBB"
		fmt.Println("nothingMap=", nothingMap) //nothingMap= map[]
	}

	{
		var nothingSlice []int //var nothingSlice []int = []int{} 实例化
		//nothingSlice[0] = 100 // index out of range [0] with length 0 空切片不能这样赋值？
		nothingSlice = append(nothingSlice, 100) //切片也是引用类型，但是不用实例化，可以使用append
		fmt.Println(nothingSlice)                //[100]
	}
}

/*
	func add(a, b int){
		a = a+b
	}
*/

func add(a, b *int) { //指针类型
	*a = *a + *b //*a a盒子内的东西
}

func put(m map[string]string) {
	m["a"] = "AAA"
}
