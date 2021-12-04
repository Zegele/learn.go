package main

import (
	"fmt"
)

type Point struct {
	x float64
	y float64
}

func main() {
	var a, b, c, d Point

	fmt.Print("请输入A点坐标(x, y)：")
	fmt.Scanf("%f,%f\n", &a.x, &a.y)
	fmt.Println(a)

	fmt.Print("请输入B点坐标：")
	for {
		fmt.Scanf("%f,%f\n", &b.x, &b.y)
		if b.x != a.x || b.y != a.y {
			break
		} else {
			fmt.Print("有重复的点，请重新输入B点坐标：")
		}
	}
	fmt.Println(b)

	fmt.Print("请输入C点坐标：")
	for {
		fmt.Scanf("%f,%f", &c.x, &c.y)
		if (c.x != b.x || c.y != b.y) && (c.x != a.x || c.y != a.y) {
			break
		} else {
			fmt.Print("有重复的点，请重新输入C点坐标：")
		}
	}
	fmt.Println(c)

	fmt.Print("请输入D点坐标(x,y):")

	for {
		fmt.Scanf("%f,%f", &d.x, &d.y)
		if ((d.x != c.x) || (d.x != b.x) || (d.x != a.x)) && ((d.y != c.y) || (d.y != b.y) || (d.y != a.y)) {
			fmt.Scanf("%f,%f", &d.x, &d.y) //为什么第一个会被跳过？
			break
		}
	}
	fmt.Println(d)

	fmt.Print("你想让A点与B、C、D 中的哪个点组成直线？") //选一个点与A组成直线，另两个点组成一条直线
	var k1, k2 float64
	var d1, d2 float64
	var s1, s2 string
	var selectP string
	for {
		fmt.Scanln(&selectP)
		if selectP == "B" || selectP == "b" {
			k1, d1, s1 = line(a, b)
			k2, d2, s2 = line(c, d)
			break
		} else if selectP == "C" || selectP == "c" {
			k1, d1, s1 = line(a, c)
			k2, d2, s2 = line(b, d)
			break
		} else if selectP == "D" || selectP == "d" {
			k1, d1, s1 = line(a, d)
			k2, d2, s2 = line(b, c)
			break
		} else {
			fmt.Println("输入错误，请重新选择输入其中一个（B、C、D）：")
		}
	}

	//判断是否平行
	if s1 != "" && s2 != "" {
		fmt.Println("这两条直线平行")
	} else if (d1 != d2) && (k1 == k2) {
		fmt.Println("这两条直线平行")
	} else if (d1 == d2) && (k1 == k2) {
		fmt.Println("这两条直线相交得都重合了！")
	} else {
		fmt.Println("这两条直线相交")
	}

}

func line(p1, p2 Point) (k float64, d float64, nok string) {
	if p1.x-p2.x == 0 {
		fmt.Printf("直线表达式是：x = %v\n", p1.x)
		nok = "k不存在"
		return k, d, nok

	} else if p1.y-p2.y == 0 {
		fmt.Printf("直线表达式是：y = %v\n", p1.y)
		k = 0.0
		d = p1.y
		return
	} else {
		k = (p1.y - p2.y) / (p1.x - p2.x)
		d = p1.y - k*p1.x
		fmt.Printf("直线表达式是：y = %vx + %v\n", k, d)
		return
	}

}
