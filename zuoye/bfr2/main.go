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
	gobmi "github.com/armstrongli/go-bmi"
	"learn.go/zuoye/bfr_rely_on/a_bfr"
	"learn.go/zuoye/bfr_rely_on/bfr_suggest"
	"learn.go/zuoye/bfr_rely_on/bmi"
	"learn.go/zuoye/bfr_rely_on/continue_ornot"
	"learn.go/zuoye/bfr_rely_on/hownum"
	"learn.go/zuoye/bfr_rely_on/putin"
)

/*
项目A:
	a文件夹:
		types.go
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
		num := hownum.Hownum()
		totalbfr := 0.0 //初始化总体脂率为0
		for i := 1; i <= num; i++ {
			//录入数据
			name, age, sexval, weight, tall = putin.Putin()

			//计算体脂率
			//自己改造的
			//calcbmi := calcbmi.Calcbmi(weight, tall) //其实不需要这个中间件
			bfr := bfr1.Calcbfr(bmi.Calcbmi(weight, tall), age, sexval) //体脂率
			totalbfr += bfr                                             //总体脂率累加

			//引用老师的
			bmiTeacher, err := gobmi.BMI(weight, tall)
			if err != nil {
				fmt.Println(err)
			}
			bfrTeacher := bfr1.Calcbfr(bmiTeacher, age, sexval)

			//suggest
			sug := bfr1.BfrSuggest(bfr, sexval, age)
			sugTeacher := bfr1.BfrSuggest(bfrTeacher, sexval, age)
			//输出信息
			fmt.Printf("%s的体脂是：%f,给您的建议是：%s。\n", name, bfr, sug)
			fmt.Printf("%s的体脂是：%f,给您的建议是：%s。\n", name, bfrTeacher, sugTeacher)
		}

		//计算平均体脂率
		a_bfr.ABfr(totalbfr, num)

		//是否继续
		if cont := continue_ornot.ContinueOrNot(); !cont {
			break
		}
	}
}
