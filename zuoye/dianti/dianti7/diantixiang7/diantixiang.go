package diantixiang7

import (
	"fmt"
	ren2 "learn.go/zuoye/dianti/dianti7/ren"
	"time"
)

type DianTiXiang struct {
	MaxFloor         int    // 最高几层 5
	MinFloor         int    // 最低几层 1
	AtFloor          int    // 电梯在哪层 3
	FangXiang        string //电梯运行方向 由第一个req决定
	SayFloor         string // "到 3 楼了" 电梯运行经过哪层，就汇报电梯在哪层
	NewReqFloorSlice []int  //把 ren.ReqFloorSlice 排序
	//	ToUpRequestFloorSlice   []int     // 当电梯向上运行，对接这个参数。拆ren.ReqSlice为两个slice
	//	ToDownRequestFloorSlice []int     //当电梯向下运行，对接这个参数。拆ren.ReqSlice为两个slice
	//	ToUpTargetFloorSlice    []int     // 当电梯向上运行，对接这个参数。 [UpTarget]  <-- 拆ren.TargetSlice为两个slice
	//	ToDownTargetFloorSlice  []int     // 当电梯向下运行，对接这个参数。
	ToUpFloorSlice   []int     // 当电梯向上运行，对接这个参数。 [UpTarget]+[UpRequest]  <-- 拆ren.TargetSlice 和 ren.ReqestSlice
	ToDownFloorSlice []int     // 当电梯向下运行，对接这个参数。[DownTarget]+[DownRequest]
	Ren              *ren2.Ren //是哪个包的盯清楚哦！
}

//调用可以知道电梯在哪（像显示电梯在哪的提示灯）
func (d *DianTiXiang) DianTiAtWhere() (atWhere int, err error) {
	if d.AtFloor > d.MaxFloor || d.AtFloor < d.MinFloor {
		return 0, fmt.Errorf("楼层有误！")
	}
	fmt.Printf("已在 %d 楼\n", d.AtFloor)
	return d.AtFloor, nil
}

//电梯运行方向（根据需要生成电梯的运行方向）
func (d *DianTiXiang) DianTiYunXingFangXiang() (fangxiang string, newReqSlice []int) {
	d.NewReqFloorSlice = append(d.NewReqFloorSlice, d.Ren.ReqFloorSlice...)
	d.NewReqFloorSlice = SliceSmallToBig(d.NewReqFloorSlice)

	if d.AtFloor > d.Ren.ReqFloorSlice[0] {
		fmt.Println(d.Ren.ReqFloorSlice)
		d.FangXiang = "ToDown"
	} else if d.AtFloor < d.Ren.ReqFloorSlice[0] {
		d.FangXiang = "ToUp"
	} else {
		d.FangXiang = "不动"
		fmt.Printf("已在 %d 楼\n", d.AtFloor)
		d.DianTiDoor()
		if d.AtFloor > d.Ren.WantFloorSlice[0] {
			d.FangXiang = "ToDown"
		} else if d.AtFloor < d.Ren.WantFloorSlice[0] {
			d.FangXiang = "ToUp"
		}
		//d.Ren.ReqFloorSlice = d.Ren.ReqFloorSlice[1:]
		//d.DianTiYunXingFangXiang()
	}
	return d.FangXiang, d.NewReqFloorSlice
}

/*
//整合ren.TargetSlice 成为 ToUpTargetFloorSlice 和 ToDownTargetFloorSlice
func (d *DianTiXiang) JoinTargetUpAndDown() {
	for _, val := range d.Ren.WantFloorSlice {
		if val > d.AtFloor {
			d.ToUpTargetFloorSlice = append(d.ToUpTargetFloorSlice, val)
		} else if val < d.AtFloor {
			d.ToDownTargetFloorSlice = append(d.ToDownTargetFloorSlice, val)
		}
	}
	//去重TargetUp
	d.ToUpTargetFloorSlice = SliceQuChong(d.ToUpTargetFloorSlice)
	//Up 从小到大排序
	d.ToUpTargetFloorSlice = SliceSmallToBig(d.ToUpTargetFloorSlice)

	//去重TargetDown
	d.ToDownTargetFloorSlice = SliceQuChong(d.ToDownTargetFloorSlice)
	//Down 从大到小排序
	d.ToDownTargetFloorSlice = SliceBigToSmall(d.ToDownTargetFloorSlice)
}

//整合ren.requestSlice，成为 ToUpRequestFloorSlice 和 ToDownRequestFloorSlice
func (d *DianTiXiang) JoinRequestUpAndDown() {
	for _, val := range d.Ren.ReqFloorSlice {
		if val > d.AtFloor {
			d.ToUpRequestFloorSlice = append(d.ToUpRequestFloorSlice, val)
		} else if val < d.AtFloor {
			d.ToDownRequestFloorSlice = append(d.ToDownRequestFloorSlice, val)
		}
	}
	//去重RequestUp
	d.ToUpRequestFloorSlice = SliceQuChong(d.ToUpRequestFloorSlice)
	//Up 从小到大排序
	d.ToUpRequestFloorSlice = SliceSmallToBig(d.ToUpRequestFloorSlice)

	//去重RequestDown
	d.ToDownRequestFloorSlice = SliceQuChong(d.ToDownRequestFloorSlice)
	//Down 从大到小排序
	d.ToDownRequestFloorSlice = SliceBigToSmall(d.ToDownRequestFloorSlice)
}
*/

func (d *DianTiXiang) JoinUpAndDown() {
	fmt.Println(d.NewReqFloorSlice)
	for _, valu := range d.NewReqFloorSlice { //[4,3,5]-> [3,4,5]  [5,1,2,4] --> [1,2,4,5]
		//todo 怎么做？ 有ren.Req ,有ren.Want ,有newReq排好序的
		if valu > d.NewReqFloorSlice[0] {
			d.ToUpFloorSlice = append(d.ToUpFloorSlice, valu)
		} else if valu < d.NewReqFloorSlice[0] {
			d.ToDownFloorSlice = append(d.ToDownFloorSlice, valu)
		} else if d.FangXiang == "ToUp" && valu == d.NewReqFloorSlice[0] { //解决valu == d.NewReqFloorSlice[0]相等时，把value放在 Up 的情况。
			d.ToUpFloorSlice = append(d.ToUpFloorSlice, valu)
		} else if d.FangXiang == "ToDown" && valu == d.NewReqFloorSlice[0] { //解决valu == d.NewReqFloorSlice[0]相等时，把value放在 Down 的情况。
			d.ToDownFloorSlice = append(d.ToDownFloorSlice, valu)
		}
	}
	fmt.Println(d.ToUpFloorSlice, "5555", d.ToDownFloorSlice)
	fmt.Println(d.NewReqFloorSlice)
	for _, valu := range d.Ren.WantFloorSlice {
		if valu > d.Ren.ReqFloorSlice[0] {
			d.ToUpFloorSlice = append(d.ToUpFloorSlice, valu)
		} else if valu < d.Ren.ReqFloorSlice[0] {
			d.ToDownFloorSlice = append(d.ToDownFloorSlice, valu)
		}
	}
	fmt.Println(d.ToUpFloorSlice, "6666", d.ToDownFloorSlice)

	//去重Ups
	d.ToUpFloorSlice = SliceQuChong(d.ToUpFloorSlice)
	//Up从小到大排序
	d.ToUpFloorSlice = SliceSmallToBig(d.ToUpFloorSlice)

	//去重Down
	d.ToDownFloorSlice = SliceQuChong(d.ToDownFloorSlice)
	//Down从大到小排序
	d.ToDownFloorSlice = SliceBigToSmall(d.ToDownFloorSlice)
}

//去重函数
func SliceQuChong(yuanShiSlice []int) (quChongSlice []int) {
	for i := 0; i < len(yuanShiSlice); i++ {
		for j := i + 1; j < len(yuanShiSlice); {
			if yuanShiSlice[i] == yuanShiSlice[j] {
				yuanShiSlice = append(yuanShiSlice[:j], yuanShiSlice[j+1:]...)
			} else {
				j++
			}
		}
	}
	quChongSlice = yuanShiSlice
	return quChongSlice
}

//从小到大排序函数
func SliceSmallToBig(quChongSlice []int) (smallToBigSlice []int) {
	for i := 0; i < len(quChongSlice)-1; i++ {
		for j := 0; j < len(quChongSlice)-1-i; j++ {
			if quChongSlice[j] > quChongSlice[j+1] {
				quChongSlice[j], quChongSlice[j+1] = quChongSlice[j+1], quChongSlice[j]
			}
		}
	}
	smallToBigSlice = quChongSlice
	return
}

//从大到小排序函数
func SliceBigToSmall(quChongSlice []int) (bigToSmallSlice []int) {
	for i := 0; i < len(quChongSlice)-1; i++ {
		for j := 0; j < len(quChongSlice)-1-i; j++ {
			if quChongSlice[j] < quChongSlice[j+1] {
				quChongSlice[j], quChongSlice[j+1] = quChongSlice[j+1], quChongSlice[j]
			}
		}
	}
	bigToSmallSlice = quChongSlice
	return
}

//电梯走到需求的地方
func (d *DianTiXiang) NewMoveTo() {
	if d.FangXiang == "ToUp" {
		fmt.Println("电梯向上运行>>>")
		for chaJiLou := d.ToUpFloorSlice[0] - d.AtFloor; chaJiLou > 0; chaJiLou-- {
			d.AtFloor++
			d.SayJiFloor()
		}
		d.DianTiDoor()
		d.ToUpFloorSlice = d.ToUpFloorSlice[1:] //到这层后，就把这层删掉
		d.MoveToUp()
		d.FangXiang = "ToDown"
		fmt.Println("电梯向下运行>>>")
		d.MoveToDown()
	} else if d.FangXiang == "ToDown" {
		fmt.Println("电梯向下运行>>>")
		for chaJiLou := d.AtFloor - d.ToDownFloorSlice[0]; chaJiLou > 0; chaJiLou-- {
			d.AtFloor--
			d.SayJiFloor()
		}
		d.DianTiDoor()
		d.ToDownFloorSlice = d.ToDownFloorSlice[1:] //到这层后，就把这层删掉
		d.MoveToDown()
		d.FangXiang = "ToUp"
		fmt.Println("电梯向上运行>>>")
		d.MoveToUp()
	}
}

/*
func (d *DianTiXiang) MoveTo() {
	if d.FangXiang == "ToUp" {
		fmt.Println("电梯向上运行>>>")
		d.MoveToUp()
		d.FangXiang = "ToDown"
		fmt.Println("电梯向下运行>>>")
		d.MoveToDown()

	} else if d.FangXiang == "ToDown" {
		fmt.Println("电梯向下运行>>>")
		d.MoveToDown()
		d.FangXiang = "ToUp"
		fmt.Println("电梯向上运行>>>")
		d.MoveToUp()
	}
}

*/

func (d *DianTiXiang) MoveToUp() {
	for _, floor := range d.ToUpFloorSlice {
		chaJiLou := floor - d.AtFloor
		for i := 0; i < chaJiLou; i++ {
			d.AtFloor++
			d.SayJiFloor()
		}
		d.ToUpFloorSlice = d.ToUpFloorSlice[1:]
		d.DianTiDoor()
	}
}

func (d *DianTiXiang) MoveToDown() {
	for _, floor := range d.ToDownFloorSlice {
		chaJiLou := d.AtFloor - floor
		for i := 0; i < chaJiLou; i++ {
			d.AtFloor--
			d.SayJiFloor()
		}
		d.ToDownFloorSlice = d.ToDownFloorSlice[1:]
		d.DianTiDoor()
	}
}

//电梯门开关
func (d *DianTiXiang) DianTiDoor() {
	fmt.Println("叮（或待停稳） --> 1s")
	time.Sleep(1 * time.Second)
	fmt.Println("开门 --> 3s")
	time.Sleep(3 * time.Second)
	fmt.Println("关门 --> 1s")
	time.Sleep(1 * time.Second)
}

//电梯报楼层
func (d *DianTiXiang) SayJiFloor() (sayFloor string) {
	time.Sleep(1 * time.Second) //1秒过一层
	d.SayFloor = "到" + fmt.Sprintf(" %d ", d.AtFloor) + "楼了............."
	fmt.Println(d.SayFloor)
	return d.SayFloor
}
