package main

import "fmt"

func main() {
	fmt.Println("1st version")
	//数组难以长期维护
	xqInfo := [3]string{"小强", "男", "在职"}
	xlInfo := [3]string{"小李", "男", "在职"}
	xhInfo := [3]string{"小红", "男", "在职"}

	fmt.Println(xqInfo)
	fmt.Println(xlInfo)
	fmt.Println(xhInfo)
	//...

	fmt.Println("2nd version")
	//难点：数组长度管理
	newPersonInfos := [3][3]string{
		[3]string{"小强", "男", "在职"}, //注意这个最后的 , 号
		[3]string{"小李", "男", "在职"},
		[3]string{"小红", "男", "在职"},
	}
	for _, val := range newPersonInfos {
		fmt.Println(val)
	}

	fmt.Println("3rd version")
	// 支持动态添加
	newPersonInfos2 := [][3]string{ //[...][3]string
		[3]string{"小强", "男", "在职"},
		[3]string{"小李", "男", "在职"},
		[3]string{"小红", "男", "在职"},
		[3]string{"小李2", "女", "在职"},
		[3]string{"小红4", "女", "在职"},
	}
	for _, val := range newPersonInfos2 {
		fmt.Println(val)
	}

	newPersonInfos2 = append(newPersonInfos2, [3]string{"大牛", "不明", "未知"}) //在数组中，无法append，换成slice后可以更改

	fmt.Println("用降维方式输出：")
	for d1, d1val := range newPersonInfos2 {
		for d2, d2val := range d1val {
			fmt.Println(d1, d1val, d2, "val:", d2val)
		}
	}
}
