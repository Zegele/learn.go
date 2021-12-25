package putin

import "fmt"

func putinName() (name string, err error) {
	fmt.Scanln(&name)
	if name == "" {
		return "0", fmt.Errorf("名字不能为空，请输入名字")
	}
	return name, nil
}

func putinAge() (age int, err error) {
	fmt.Scanln(&age)
	if age < 18 || age > 150 {
		return -1, fmt.Errorf("请重新输入年龄（整数），并确保年龄大于等于18，且小于等于150")
	}
	return age, nil
}

func putinSex() (sex string, err error) {
	fmt.Scanln(&sex)
	if sex != "男" && sex != "man" && sex != "女" && sex != "woman" {
		return "-2", fmt.Errorf("请重新输入年龄（整数），并确保年龄大于等于18，且小于等于150")
	}
	return sex, nil
}

func putinTall() (tallM float64, err error) {
	fmt.Scanln(&tallM)
	if tallM < 0.5 || tallM > 3 {
		return -3, fmt.Errorf("请重新输入身高（m），并确保身高大于0.5，且小于3")
	}
	return tallM, nil
}

func putinWeight() (weightKG float64, err error) {
	fmt.Scanln(&weightKG)
	if weightKG < 20 || weightKG > 1000 {
		return -4, fmt.Errorf("请重新输入体重（kg），并确保体重大于20，且小于1000")
	}
	return weightKG, nil
}

func PutinAll() (name string, age int, sex string, tall float64, weight float64, err error) {

	name, err0 := putinName()
	for err0 != nil {
		name, err0 = putinName()
	}

	age, err1 := putinAge()
	for err1 != nil {
		age, err1 = putinAge()
	}

	sex, err2 := putinSex()
	for err2 != nil {
		sex, err2 = putinSex()
	}

	tall, err3 := putinTall()
	for err3 != nil {
		tall, err3 = putinTall()
	}
	weight, err4 := putinWeight()
	for err4 != nil {
		weight, err3 = putinWeight()
	}

	return name, age, sex, tall, weight, nil

	return PutinAll()
}
