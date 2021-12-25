package putin

import "fmt"

func putinName() (name string, err error) {
	fmt.Print("请输入名字：")
	fmt.Scanln(&name)
	if name == "" {
		return "0", fmt.Errorf("名字不能为空。\n")
	}
	return name, nil
}

func putinAge() (age int, err error) {
	fmt.Print("请输入年龄：")
	fmt.Scanln(&age)
	if age < 18 || age > 150 {
		return -1, fmt.Errorf("请重新输入年龄（整数），并确保年龄大于等于18，且小于等于150\n")
	}
	return age, nil
}

func putinSex() (sex string, err error) {
	fmt.Print("请输入性别（男（man）或 女（woman））：")
	fmt.Scanln(&sex)
	if sex != "男" && sex != "man" && sex != "女" && sex != "woman" {
		return "-2", fmt.Errorf("请重新输入性别，男（man）或 女（woman）\n")
	}
	return sex, nil
}

func putinTall() (tallM float64, err error) {
	fmt.Print("请输入身高（M）：")
	fmt.Scanln(&tallM)
	if tallM < 0.5 || tallM > 3 {
		return -3, fmt.Errorf("请重新输入身高（m），并确保身高大于0.5，且小于3\n")
	}
	return tallM, nil
}

func putinWeight() (weightKG float64, err error) {
	fmt.Print("请输入体重（kg）：")
	fmt.Scanln(&weightKG)
	if weightKG < 20 || weightKG > 1000 {
		return -4, fmt.Errorf("请重新输入体重（kg），并确保体重大于20，且小于1000\n")
	}
	return weightKG, nil
}

func PutinAll() (name string, age int, sex string, tall float64, weight float64, err error) {

	name, err0 := putinName()
	for err0 != nil {
		fmt.Print(err0)
		name, err0 = putinName()
	}

	age, err1 := putinAge()
	for err1 != nil {
		fmt.Print(err1)
		age, err1 = putinAge()
	}

	sex, err2 := putinSex()
	for err2 != nil {
		fmt.Print(err2)
		sex, err2 = putinSex()
	}

	tall, err3 := putinTall()
	for err3 != nil {
		fmt.Print(err3)
		tall, err3 = putinTall()
	}

	weight, err4 := putinWeight()
	for err4 != nil {
		fmt.Print(err4)
		weight, err4 = putinWeight()
	}
	fmt.Printf("基本信息：\n姓名：%s，年龄：%d，性别：%s，身高：%f，体重：%f\n", name, age, sex, tall, weight)
	return name, age, sex, tall, weight, nil
}

func ContinueOrNot() bool {
	var continueornot string
	fmt.Println("是否继续？(y/n)")
	for {
		fmt.Scanln(&continueornot)
		if continueornot == "n" {
			return false
		} else if continueornot == "y" {
			return true
		} else {
			fmt.Println("输入有误，请输入 y 或 n 。")
		}
	}
}

// os.Exit(1)
