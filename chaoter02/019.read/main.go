package main

import (
	"fmt"
	"github.com/spf13/cobra"
	learn_go_tools "learn.go.tools"
	calculator "learn.go/chaoter02/015.fatrate.refactor/calc"
)

func main() {
	//录入
	var (
		name   string
		sex    string
		tall   float64
		weight float64
		age    int
	)

	//arguments := os.Args() //go中的读取参数 参数传入[]string（字符串切片）。第1个参数[0]是程序本身，第二个参数[1]或之后的，就是传入的参数。
	//局限性是，传入的参数都是string类型的。

	cmd := &cobra.Command{
		Use:   "healthcheck",
		Short: "体脂计算器，根据身高、体重、性别、年龄计算体脂比，并给出健康建议",
		Long:  "该体脂计算器是基于BMI的体脂计算器，等详细描述。。。",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("name: ", name)
			fmt.Println("sex:", sex)
			fmt.Println("tall:", tall)
			fmt.Println("weight:", weight)
			fmt.Println("age:", age)
			//计算
			bmi := calculator.CalcBMI(tall, weight)
			fatRate := calculator.CalcFatRate(bmi, age, sex)
			fmt.Println("fatRate: ", fatRate)

			// 评价结果  //vendor测试?
			fmt.Println(learn_go_tools.Max(3, 5))
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "姓名")
	cmd.Flags().StringVar(&sex, "sex", "", "性别")
	cmd.Flags().Float64Var(&tall, "tall", 0, "身高")
	cmd.Flags().Float64Var(&weight, "weight", 0, "体重")
	cmd.Flags().IntVar(&age, "age", 0, "年龄")

	cmd.Execute()
	//计算

	//评估结果

	// terminal
	//第7课时间50：08

}
