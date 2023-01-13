package main

import (
	"flag"
	"fmt"
	"strings"
)

//  解析喜欢的编程语言，并直接解析到slice中，我们可以定义如下sliceValue类型，然后实现Value接口：

//定义一个类型，用于增加该类型方法
type sliceValue []string

//new一个存放命令行参数值的slice
func newSliceValue(vals []string, p *[]string) *sliceValue {
	*p = vals
	return (*sliceValue)(p) //把p的*[]string类型转化为*sliceValve类型
}

/*
//Value 接口：
type Value interface{
	String() string
	Set(string) error
}
实现flag包中的Value接口，将命令行接收到的值用 ， 分隔，存到slice里。
*/
func (s *sliceValue) Set(val string) error {
	*s = sliceValue(strings.Split(val, ","))
	return nil
}

// flag为slice的默认值default is me， 和 return 返回值没有关系
func (s *sliceValue) String() string {
	*s = sliceValue(strings.Split("defalut is me", ","))
	return "it's none of my business"
}

/*
可执行文件名 -slice="java, go" 最后将输出[java, go]
可执行文件名 最后将输出[default is me]

*/

func main() {
	var languages []string
	flag.Var(newSliceValue([]string{}, &languages), "slice", "i like programming `languages`")
	flag.Parse()

	//打印结果slice接收到的值
	fmt.Println(languages)
}
