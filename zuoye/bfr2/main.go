package main

/*
计算多个人的平均体脂
1. 实现完整的体脂计算器
2. 连续输入3人的姓名、性别、身高、体重、年龄信息
3. 计算每个人的BMI，体脂率
输出：
1. 每个人姓名、BMI、体脂率、建议。
2. 总人数，平均体脂率
*/
import (
	"fmt"
	"learn.go/zuoye/bfr_rely_on/a_bfr"
	"learn.go/zuoye/bfr_rely_on/bfr_suggest"
	"learn.go/zuoye/bfr_rely_on/bmi"
	"learn.go/zuoye/bfr_rely_on/continue_ornot"
	"learn.go/zuoye/bfr_rely_on/putin"
)

/*
项目A:
	a文件夹:
		main.go
	b文件夹:
		bb.go
	c文件夹:
		cc.go
	mod.go

这个项目A结构中，如果main引用bb或cc中的函数，bb或cc中的函数首字母都得用大写（公有函数）么？
试过小写不行。

*/

func main() {
	var (
		name   string
		sexval float64
		tall   float64
		weight float64
		age    int
	)
	fmt.Println("体脂计算器")
	for {
		//输入：几个人参与体脂计算
		num := hownum()
		totalbfr := 0.0 //初始化总体脂率为0
		for i := 1; i <= num; i++ {
			//录入数据
			name, age, sexval, weight, tall = putin.Putin()

			//计算体脂率
			//calcbmi := calcbmi.Calcbmi(weight, tall) //其实不需要这个中间件
			bfr := bfr1.Calcbfr(bmi.Calcbmi(weight, tall), age, sexval) //体脂率
			totalbfr += bfr                                             //总体脂率累加

			//suggest
			sug := bfr1.BfrSuggest(bfr, sexval, age)

			//输出信息
			fmt.Printf("%s的体脂是：%f,给您的建议是：%s。\n", name, bfr, sug)
		}

		//计算平均体脂率
		a_bfr.ABfr(totalbfr, num)

		//是否继续
		if cont := continue_ornot.ContinueOrNot(); !cont {
			break
		}
	}
}

func hownum() (num int) {
	fmt.Print("你要录入几个人的数据？请输入数字：")
	for {
		fmt.Scanln(&num)
		if num <= 0 {
			fmt.Println("请输入一个大于0的数字")
		} else {
			return
		}
	}
}
