package calcbft

import (
	"fmt"
	"strconv"
)

func CalcBFT(bmi float64, age int, sex string) (bft float64, err error) {
	//排除非法bmi
	if bmi <= 0 {
		return 0, fmt.Errorf("bft不能为0，或负数")
	}
	//排除非法age
	if age <= 0 && age >= 150 {
		return -1, fmt.Errorf("age不能是0， 或负数，且不能大于150")
	}
	//排除非法性别
	if sex != "男" && sex != "man" && sex != "woman" && sex != "女" { //注意 使用 || ，还是使用&&。
		return -2, fmt.Errorf("传入的性别不是 男（man），或 女（woman） ")
	}

	var sexval float64
	if sex == "man" || sex == "男" {
		sexval = 1
	} else {
		sexval = 0
	}

	bft = (1.2*(bmi*100) + 0.23*float64(age) - 5.4 - 10.8*(sexval)) / 100
	bft, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", bft), 64)
	return bft, nil
}
