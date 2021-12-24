package calcbft

import (
	"fmt"
	"testing"
)

//测试体脂率计算

func TestCalcBFT(t *testing.T) {
	//录入正常BMI、年龄、性别，确保计算结果符合预期
	{
		t_bft, err := CalcBFT(0.20, 30, "man") //tall 1.8, weight 65
		if err != nil {
			t.Fatalf("预期的err是nil,但测试结果为%v", err)
		}
		if t_bft != 0.147 {
			t.Fatalf("预期结果为0.147，但测试结果为：%f:", t_bft)
		}
		//bfr = (1.2*(bmi*100) + 0.23*float64(age) - 5.4 - 10.8*(sexval)) / 100
	}

	//录入非法BMI，返回错误
	{
		t_bft, err := CalcBFT(0, 30, "man")

		if t_bft != 0 {
			t.Fatalf("预期结果为：0，但测试结果为：%f:", t_bft)
		}
		if err == fmt.Errorf("bft不能为0，或负数") { //err != nil
			t.Fatalf("预期结果为：bft不能为0，或负数，但测试结果为：%s", err)
		}
	}

	//录入非法age，返回错误
	{
		t_bft, err := CalcBFT(0.20, -1, "man") //tall 1.8, weight 65
		if err == fmt.Errorf("age不能是0， 或负数，且不能大于150") {
			t.Fatalf("预期的err是:age不能是0， 或负数，且不能大于150,但测试结果为%v", err)
		}
		if t_bft != -1 {
			t.Fatalf("预期结果为:-1，但测试结果为：%f:", t_bft)
		}
	}

	//录入非法性别，返回错误
	{
		t_bft, err := CalcBFT(0.20, 30, "an") //tall 1.8, weight 65
		if err == fmt.Errorf("传入的性别不是 男（man），或 女（woman）") {
			t.Fatalf("预期的err是：传入的性别不是男（man）或女（woman）,但测试结果为%v", err)
		}
		if t_bft != -2 {
			t.Fatalf("预期结果为-2，但测试结果为：%f:", t_bft)
		}
	}

}
