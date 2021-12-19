package a_bfr

import "fmt"

func aBfr(list []Person) {
	tBfr := 0.0
	for i := 0; i < len(list); i++ {
		tBfr += list[i].bfr
	}
	aBfr := tBfr / float64(len(list))
	fmt.Printf("总人数为%d人，这%d人的平均体脂率为：%.2f\n", len(list), len(list), aBfr)
}
