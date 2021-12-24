package calcbmi

import "testing"

func TestCalcBMI1(t *testing.T) {

	//计算测试
	//传入合法数据
	{
		t_bmi1, err := CalcBMI(1.8, 65.0)
		if t_bmi1 != 0.20 {
			t.Fatalf("预期bmi结果为0.20，但是得到：%f", t_bmi1) //t.Fatal 失败，并退出
		}
		if err != nil {
			t.Fatalf("预期err结果为nil，但是得到：%v", err)
		}
	}

	//传入非法数据 tall<=0

	{
		t_bmi2, err := CalcBMI(0, 65)
		if t_bmi2 != -1.0 {
			t.Fatalf("预期bmi结果为-1.0，但是得到：%f", t_bmi2)
		}
		if err == nil {
			t.Fatalf("预期err结果为:身高不能是0或负数，但是得到：%v", err)
		}
	}

	//传入非法数据 weight<=0
	{
		t_bmi2, err := CalcBMI(1.8, 0)
		if t_bmi2 != -1.0 {
			t.Fatalf("预期bmi结果为-1.0，但是得到：%f", t_bmi2)
		}
		if err == nil {
			t.Fatalf("预期err结果为:体重不能是0或负数，但是得到：%v", err)
		}
	}

}
