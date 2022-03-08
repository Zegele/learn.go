package calculate

import (
	"learn.go/zuoye/zuoye5_tizhi/pkg/apii"
)

type Calc struct {
}

func (c *Calc) calcBMI(p *apii.Person) (bmi float32) {
	bmi = p.Weight / (p.Tall * p.Tall) / 100 //除以100，这里改动
	return
}

func (c *Calc) CalcFatRate(p *apii.Person) *apii.Person {
	sexWeight := 0
	if p.Sex == "男" || p.Sex == "man" {
		sexWeight = 1
	}
	p.Fatr = (1.2*(c.calcBMI(p)*100) + 0.23*float32(p.Age) - 5.4 - 10.8*float32(sexWeight)) / 100
	return p
}
