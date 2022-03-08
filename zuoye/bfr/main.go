package main

/*
计算多个人的平均体脂
1. 实现完整的体脂计算器
2. 连续输入3人的姓名、性别、身高、体重、年龄信息
3. 计算每个人的BMI，体脂率
输出：
1. 每个人姓名、BMI、体脂率、建议。
2. 总人数，平均体脂率
*/
import (
	"fmt"
)

type Person struct {
	num     int
	name    string
	age     int
	sex     string
	sexval  float64
	tall    float64
	weight  float64
	bmi     float64
	bfr     float64
	suggest string
}

func main() {
	var person Person
	var list []Person
	fmt.Println("体脂计算器")

	for i := 1; i <= 3; i++ {

		person.num = i

		fmt.Print("请输入姓名：")
		fmt.Scanln(&person.name)

		fmt.Print("请输入年龄：")
		fmt.Scanln(&person.age)
		for person.age < 18 || person.age > 150 {
			fmt.Print("抱歉，不在符合的年龄范围，请输入18-150的整数：")
			fmt.Scanln(&person.age)
		}
		/*
			for {
				fmt.Scanln(&client.age)
				if client.age < 18 || client.age > 150 {
					fmt.Print("抱歉，不在符合的年龄范围，请输入18-150的整数：")
				} else {
					break
				}
			}
		*/

		fmt.Print("请输入性别（男/女）：")
		for {
			fmt.Scanln(&person.sex)
			if person.sex == "男" || person.sex == "man" {
				person.sexval = 1.0
				break
			} else if person.sex == "女" || person.sex == "woman" {
				person.sexval = 0.0
				break
			} else {
				fmt.Print("性别输入有误请重新输入（男/女）:")
			}
		}

		fmt.Print("请输入体重（kg）：")
		fmt.Scanln(&person.weight)
		for person.weight < 20 || person.weight > 1000 {
			fmt.Println("抱歉，体重(kg)不在计算区间，请输入20-1000之间的数。")
			fmt.Scanln(&person.weight)
		}
		/*
			for {
				fmt.Scanln(&client.weight)
				if client.weight < 20 || client.weight > 1000 {
					fmt.Println("抱歉，体重不在计算区间，请输入20-1000之间的数。")
				} else {
					break
				}
			}
		*/
		fmt.Print("请输入身高（米）：")
		fmt.Scanln(&person.tall)
		for person.tall < 0.5 || person.tall > 3 {
			fmt.Println("抱歉，身高（米）不在计算区间，请输入0.5-3之间的数。")
			fmt.Scanln(&person.tall)
		}

		person.bmi = person.countbmi() / 100
		person.bfr = person.countbfr() / 100

		//suggest
		var sug string
		var nbfr float64 = person.bfr
		if person.sexval == 1.0 { // 男
			if person.age >= 18 || person.age < 40 {
				switch {
				case nbfr > 0.0 && nbfr <= 0.1:
					sug = "偏瘦，赶紧多吃点!"
				case nbfr > 0.1 && nbfr <= 0.16:
					sug = "标准，继续保持！"
				case nbfr > 0.16 && nbfr <= 0.21:
					sug = "偏重，现在少吃点还来得及。"
				case nbfr > 0.21 && nbfr <= 0.26:
					sug = "肥胖，抓紧运动，或许还来得及。"
				default:
					sug = "算了，放弃吧..."
				}
			} else if person.age >= 40 || person.age < 60 {
				switch {
				case nbfr > 0.0 && nbfr <= 0.11:
					sug = "偏瘦，赶紧多吃点!"
				case nbfr > 0.11 && nbfr <= 0.17:
					sug = "标准，继续保持！"
				case nbfr > 0.17 && nbfr <= 0.22:
					sug = "偏重，现在少吃点还来得及。"
				case nbfr > 0.22 && nbfr <= 0.27:
					sug = "肥胖，抓紧运动，或许还来得及。"
				default:
					sug = "算了，放弃吧..."
				}
			} else {
				switch {
				case nbfr > 0.0 && nbfr <= 0.13:
					sug = "偏瘦，赶紧多吃点!"
				case nbfr > 0.13 && nbfr <= 0.19:
					sug = "标准，继续保持！"
				case nbfr > 0.19 && nbfr <= 0.24:
					sug = "偏重，现在少吃点还来得及。"
				case nbfr > 0.24 && nbfr <= 0.29:
					sug = "肥胖，抓紧运动，或许还来得及。"
				default:
					sug = "算了，放弃吧..."
				}
			}
			person.suggest = sug //给结构体对象suggest赋值
		} else { //client.sexval == 0.0 女
			if person.age >= 18 || person.age < 40 {
				switch {
				case nbfr > 0.0 && nbfr <= 0.2:
					sug = "偏瘦，赶紧多吃点!"
				case nbfr > 0.2 && nbfr <= 0.27:
					sug = "标准，继续保持！"
				case nbfr > 0.27 && nbfr <= 0.34:
					sug = "偏重，现在少吃点还来得及。"
				case nbfr > 0.34 && nbfr <= 0.39:
					sug = "肥胖，抓紧运动，或许还来得及。"
				default:
					sug = "算了，放弃吧..."
				}
			} else if person.age >= 40 || person.age < 60 {
				switch {
				case nbfr > 0.0 && nbfr <= 0.21:
					sug = "偏瘦，赶紧多吃点!"
				case nbfr > 0.21 && nbfr <= 0.28:
					sug = "标准，继续保持！"
				case nbfr > 0.28 && nbfr <= 0.35:
					sug = "偏重，现在少吃点还来得及。"
				case nbfr > 0.35 && nbfr <= 0.40:
					sug = "肥胖，抓紧运动，或许还来得及。"
				default:
					sug = "算了，放弃吧..."
				}
			} else {
				switch {
				case nbfr > 0.0 && nbfr <= 0.22:
					sug = "偏瘦，赶紧多吃点!"
				case nbfr > 0.22 && nbfr <= 0.29:
					sug = "标准，继续保持！"
				case nbfr > 0.29 && nbfr <= 0.36:
					sug = "偏重，现在少吃点还来得及。"
				case nbfr > 0.36 && nbfr <= 0.41:
					sug = "肥胖，抓紧运动，或许还来得及。"
				default:
					sug = "算了，放弃吧..."
				}
			}
			person.suggest = sug //给结构体对象suggest赋值
		}

		list = append(list, person)
	}

	//输出信息
	for _, val := range list {
		fmt.Printf("序号:%d， 姓名：%s， BMI：%.2f， BFR:%.2f， 建议：%s\n", val.num, val.name, val.bmi, val.bfr, val.suggest)
	}

	//总人数和平均体脂率
	aBfr(list)
}

//bmi
func (p Person) countbmi() (bmi float64) {
	bmi = p.weight / (p.tall * p.tall)
	return
}

//bfr
func (p Person) countbfr() (bfr float64) {
	bfr = 1.2*(p.bmi*100) + 0.23*float64(p.age) - 5.4 - 10.8*(p.sexval)
	return
}

//平均bfr
func aBfr(list []Person) {
	tBfr := 0.0
	for i := 0; i < len(list); i++ {
		tBfr += list[i].bfr
	}
	aBfr := tBfr / float64(len(list))
	fmt.Printf("总人数为%d人，这%d人的平均体脂率为：%.2f\n", len(list), len(list), aBfr)
}
