// go语言接口的nil判断
// www.kancloud.cn/imdszxs/golang/1509740
package main

import "fmt"

// nil在go语言中只能被赋值给指针和接口
//接口在底层的实现有两个部分：type和data
//在源码中，显式地将赋值给接口时，接口的type和data都将为nil
//此时，接口与nil值判断是相等的
//但如果将一个带有类型的nil赋值给接口时，
//只有data为nil
//而type不是nil
//此时接口与nil判断将不相等

//接口与nil不相等
//下面代码使用MyImplement()实现fmt包中的Stringer接口，这个接口的定义如下：
// type Stringer interface { String()string}
//在GetStringer()函数中将返回这个接口，
//通过*MyImplement指针变量置为nil提供GetStringer的返回值
//在main()中 判断GetStringer与nil是否相等，如下：

// 定义一个结构体
type MyImplement struct{}

// 实现fmt.Stringer的String方法
func (m *MyImplement) String() string {
	return "hi"
}

// 在函数中返回fmt.Stringer接口
func GetStringer() fmt.Stringer {
	// 赋nil
	var s *MyImplement = nil
	// 返回变量
	return s //s变量此时被fmt.Stringer接口包装后，实际类型为*MyImplement，值为nil的接口
}

func main() {
	// 判断返回值是否为nil
	if GetStringer() == nil {
		fmt.Println("GetStringer() == nil")
	} else {
		fmt.Println("GetStringer() != nil")
	}
	// GetStringer()的返回值是接口
	// 虽然接口里的value为nil，但type带有*MyImplement信息，使用==判断，依然不为nil
}

// 发现nil类型值返回时直接返回nil
// 为了避免这类误判的问题，可以在函数返回时，发现带有nil的指针时直接返回nil
//如下：

func GetStringerGai() fmt.Stringer {
	var s *MyImplement = nil
	if s == nil { //如果s的值是nil，则返回nil，这时，s的值和类型就都是nil
		return nil
	}

	return s
}
