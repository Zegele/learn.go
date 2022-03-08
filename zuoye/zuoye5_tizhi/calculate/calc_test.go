package calculate

import (
	"fmt"
	"learn.go/zuoye/zuoye5_tizhi/pkg/apii"
	"testing"
)

func TestCalc(t *testing.T) {
	cal := &Calc{}
	p := &apii.Person{
		Name:   "1",
		Age:    18,
		Sex:    "man",
		Tall:   1.8,
		Weight: 70,
	}
	p = cal.CalcFatRate(p)
	fmt.Println(p)
	bmi := 70 / (1.8 * 1.8) / 100
	fatr := (bmi*100*1.2 + 0.23*18 - 5.4 - 10.8) / 100
	fmt.Println(bmi)
	fmt.Println(fatr)
}
