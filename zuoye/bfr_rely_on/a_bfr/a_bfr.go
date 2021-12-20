package a_bfr

import "fmt"

func ABfr(tBfr float64, num int) {
	aBfr := tBfr / float64(num)
	fmt.Printf("这%d人的平均体脂率为：%.2f\n", num, aBfr)
}
