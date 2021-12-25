package putin

//todo 录入测试怎么做？
import (
	"fmt"
	"testing"
)

func TestPutinAll(t *testing.T) {
	//输入合法姓名，年龄，性别，身高，体重
	{
		name, age, sex, tall, weight, err := PutinAll()
		if name != "a" {
			t.Fatalf("预期是：a，但测试结果为：%s", name)
		}
		if age != 30 {
			t.Fatalf("预期是：30，但测试结果为：%d", age)
		}
		if sex != "男" && sex != "man" && sex != "女" && sex != "woman" {
			t.Fatalf("预期是：男（man），或女（woman），但测试结果为：%s", sex)
		}
		if tall != 1.8 {
			t.Fatalf("预期是：1.8，但测试结果为：%d", age)
		}
		if weight != 65 {
			t.Fatalf("预期是：65，但测试结果为：%d", age)
		}
		if err != nil {
			t.Fatalf("预期是：nil，但测试结果为：%v", err)
		}
	}

	//输入非法年龄
	{
		age, err := putinAge()
		if age == 17 {
			t.Errorf("预期是：age>=18，但测试结果为：%d", age)
		}
		if err != fmt.Errorf("请重新输入年龄（整数），并确保年龄大于等于18，且小于等于150") {
			t.Fatalf("预期是：请重新输入年龄（整数），并确保年龄大于等于18，且小于等于150，但测试结果为：%v", err)
		}
	}

	{
		//输入非法tall
		tall, err := putinTall()
		if tall == 0.4 {
			t.Errorf("预期是：tall>=0.5 && tall<=3，但测试结果为：%f", tall)
		}
		if err != fmt.Errorf("请重新输入身高（m），并确保身高大于0.5，且小于3") {
			t.Fatalf("预期是：请重新输入tall值，并确保tall>=0.5，且tall<=3，但测试结果为：%v", err)
		}
	}

	//输入非法weight
	{
		weight, err := putinWeight()
		if weight == 17 {
			t.Errorf("预期是：weight>=20 && weight <=1000，但测试结果为：%f", weight)
		}
		if err != fmt.Errorf("请重新输入体重（kg），并确保体重大于20，且小于1000") {
			t.Fatalf("预期是：请重新输入体重（kg），并确保体重大于20，且小于1000，但测试结果为：%v", err)
		}
	}

	//输入非法sex
	{
		sex, err := putinSex()
		if sex != "男" && sex != "man" && sex != "女" && sex != "woman" {
			t.Errorf("预期是：性别是男（man）或女（woman），但测试结果为：%s", sex)
		}
		if err != fmt.Errorf("请重新输入性别，男（man）或女（woman）") {
			t.Fatalf("预期是：请重新输入性别，男（man）或女（woman），但测试结果为：%v", err)
		}
	}

}
