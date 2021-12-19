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
	"learn.go/zuoye/bfr_rely_on/bfr_suggest"
	"learn.go/zuoye/bfr_rely_on/bmi"
	"learn.go/zuoye/bfr_rely_on/putin"
)

func main() {
	var (
		name   string
		sexval float64
		tall   float64
		weight float64
		age    int
	)
	fmt.Println("体脂计算器")
	i := 0
	for {
		//录入数据
		name, age, sexval, weight, tall = putin.Putin()

		//计算体脂率
		bmi := bmi.Calcbmi(weight, tall) / 100
		bfr := bfr1.Calcbfr(bmi, age, sexval) / 100 //体脂率

		//suggest
		sug := bfr1.BfrSuggest(bfr, sexval, age)

		//输出信息
		fmt.Printf("%s的体脂是：%f,给您的建议是：%s。", name, bfr, sug)

	}
	//总人数和平均体脂率

}

//平均bfr
