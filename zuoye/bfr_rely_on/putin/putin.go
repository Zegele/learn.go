package putin

import "fmt"

func Putin() (name string, age int, sexval float64, weight float64, tall float64) {
	var sex string
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)

	fmt.Print("请输入年龄：")
	fmt.Scanln(&age)
	for age < 18 || age > 150 {
		fmt.Print("抱歉，不在符合的年龄范围，请输入18-150的整数：")
		fmt.Scanln(&age)
	}

	fmt.Print("请输入性别（男/女）：")
	for {
		fmt.Scanln(&sex)
		if sex == "男" || sex == "man" {
			sexval = 1.0
			break
		} else if sex == "女" || sex == "woman" {
			sexval = 0.0
			break
		} else {
			fmt.Print("性别输入有误请重新输入（男/女）:")
		}
	}

	fmt.Print("请输入体重（kg）：")
	fmt.Scanln(&weight)
	for weight < 20 || weight > 1000 {
		fmt.Println("抱歉，体重(kg)不在计算区间，请输入20-1000之间的数。")
		fmt.Scanln(&weight)
	}

	fmt.Print("请输入身高（米）：")
	fmt.Scanln(&tall)
	for tall < 0.5 || tall > 3 {
		fmt.Println("抱歉，身高（米）不在计算区间，请输入0.5-3之间的数。")
		fmt.Scanln(&tall)
	}
	fmt.Println("基本信息：", name, sex, age, weight, tall)
	return
}
