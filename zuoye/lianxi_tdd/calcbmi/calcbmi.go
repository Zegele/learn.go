package calcbmi

import (
	"fmt"
	"strconv"
)

func CalcBMI(tall float64, weight float64) (bmi float64, err error) {
	if tall <= 0 {
		return -1, fmt.Errorf("身高不能是0或负数。")
	}
	if weight <= 0 {
		return -1, fmt.Errorf("体重不能是0或负数。")
	}

	bmi = weight / tall / tall / 100
	bmi, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", bmi), 64) //查到了这个转换格式的包
	return bmi, nil
}
