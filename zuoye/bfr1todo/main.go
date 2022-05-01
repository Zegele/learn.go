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
	"github.com/spf13/cobra"
	bfr1 "learn.go/zuoye/bfr_rely_on/bfr_suggest"
	"learn.go/zuoye/bfr_rely_on/bmi"
	//"learn.go/zuoye/bfr_rely_on/calcbmi"
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
		sex    string
		sexval float64
		tall   float64
		weight float64
		age    int
	)
	fmt.Println("体脂计算器")
	cmd := &cobra.Command{
		Use:   "healthcheck",
		Short: "体脂计算，根据身高，体重，年龄，性别计算体脂率，并给出健康建议。",
		Long:  "...",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("name:", name)
			fmt.Println("sex:", sex)
			fmt.Println("tall:", tall)
			fmt.Println("weight:", weight)
			fmt.Println("age:", age)
			if sex == "男" || sex == "man" {
				sexval = 1
			} else {
				sexval = 0
			}
			//计算bmi bfr

			bmi := bmi.Calcbmi(weight, tall)
			bfr := bfr1.Calcbfr(bmi, age, sexval)
			//totalbfr += bfr
			//suggest
			sug := bfr1.BfrSuggest(bfr, sexval, age)
			//输出信息
			fmt.Printf("%s的体脂是：%f,给您的建议是：%s。\n", name, bfr, sug)
		},
	}
	//录入数据：
	cmd.Flags().StringVar(&name, "name", "", "姓名")
	cmd.Flags().StringVar(&sex, "sex", "", "性别")
	cmd.Flags().Float64Var(&tall, "tall", 0, "身高")
	cmd.Flags().Float64Var(&weight, "weight", 0, "体重")
	cmd.Flags().IntVar(&age, "age", 0, "年龄")

	cmd.Execute()

	//是否继续
	//使用：
}
