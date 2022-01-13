package diantixiangkuozhan

import (
	"fmt"
	ren2 "learn.go/zuoye/dianti/diantikuozhan/ren"
	"time"
)

type DianTiXiang struct {
	MaxFloor         int    // 最高几层 5
	MinFloor         int    // 最低几层 1
	AtFloor          int    // 电梯在哪层 3
	FangXiang        string //电梯运行方向 由第一个req决定
	MostWantFloors   []int  //目标楼层中人数最多的那层
	MostWantF        int
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
		d.FangXiang = "ToDown"
	} else if d.AtFloor < d.Ren.ReqFloorSlice[0] {
		d.FangXiang = "ToUp"
	} else {
		d.FangXiang = "不动"
		d.DianTiDoor()
		if d.AtFloor > d.Ren.WantFloorSlice[0] {
			d.FangXiang = "ToDown"
		} else if d.AtFloor < d.Ren.WantFloorSlice[0] {
			d.FangXiang = "ToUp"
		}
	}
	return d.FangXiang, d.NewReqFloorSlice
}

func (d *DianTiXiang) JoinUpAndDown() {
	for _, valu := range d.NewReqFloorSlice { //[4,3,5]-> [3,4,5]  [5,1,2,4] --> [1,2,4,5]
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
	for _, valu := range d.Ren.WantFloorSlice {
		if valu > d.Ren.ReqFloorSlice[0] {
			d.ToUpFloorSlice = append(d.ToUpFloorSlice, valu)
		} else if valu < d.Ren.ReqFloorSlice[0] {
			d.ToDownFloorSlice = append(d.ToDownFloorSlice, valu)
		}
	}

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
		//先运行到第一个请求电梯的地方。
		for chaJiLou := d.ToUpFloorSlice[0] - d.AtFloor; chaJiLou > 0; chaJiLou-- {
			d.AtFloor++
			d.SayJiFloor()
		}
		d.DianTiDoor()
		d.ToUpFloorSlice = d.ToUpFloorSlice[1:] //到这层后，就把这层删掉

		//运行目标最多的楼层
		for chaJiLou := d.MostWantF - d.AtFloor; chaJiLou > 0; chaJiLou-- {
			d.AtFloor++
			d.SayJiFloor()
		}
		d.DianTiDoor()

		//然后运行目标地方
		d.MoveToUp()

		if d.ToDownFloorSlice != nil { //判断是否有向下的必要 如果ToDownFloorSlice是空，就停在那层就行了。
			d.FangXiang = "ToDown"
			fmt.Println("电梯向下运行>>>")
			d.MoveToDown()
		}

	} else if d.FangXiang == "ToDown" {
		fmt.Println("电梯向下运行>>>")
		for chaJiLou := d.AtFloor - d.ToDownFloorSlice[0]; chaJiLou > 0; chaJiLou-- {
			d.AtFloor--
			d.SayJiFloor()
		}
		d.DianTiDoor()
		d.ToDownFloorSlice = d.ToDownFloorSlice[1:] //到这层后，就把这层删掉
		//todo 加运行最多楼层的动作
		for chaJiLou := d.AtFloor - d.MostWantF; chaJiLou > 0; chaJiLou-- {
			d.AtFloor--
			d.SayJiFloor()
		}
		d.DianTiDoor()

		//
		d.MoveToDown()
		if d.ToUpFloorSlice != nil { //判断是否有向下的必要 如果ToDownFloorSlice是空，就停在那层就行了。
			d.FangXiang = "ToUp"
			fmt.Println("电梯向上运行>>>")
			d.MoveToUp()
		}
	}
}

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

func (d *DianTiXiang) MostWant() { //获取目标楼层中人数最多的那层
	count := 0
	countDuibi := 0
	for _, v := range d.Ren.WantFloorSlice {
		if v == 1 {
			count++
		}
	}
	d.MostWantFloors = append([]int{}, 1)

	for i := 2; i <= d.MaxFloor; i++ {
		for _, v := range d.Ren.WantFloorSlice {
			if v == i {
				countDuibi++
			}
		}
		if count < countDuibi {
			count = countDuibi
			countDuibi = 0
			d.MostWantFloors = append([]int{}, i)
		} else if count == countDuibi {
			countDuibi = 0
			d.MostWantFloors = append(d.MostWantFloors, i)
		} else {
			countDuibi = 0
		}
	}
	d.SelectMostWant()
	fmt.Println("mostWant:", d.MostWantF)
	d.QuChuMostFloor()
}

func (d *DianTiXiang) SelectMostWant() {
	if d.FangXiang == "ToUp" {
		d.MostWantFloors = SliceSmallToBig(d.MostWantFloors) //从小到大排序
		for _, v := range d.MostWantFloors {
			if v > d.AtFloor {
				d.MostWantF = v
				fmt.Println("d.MostWantF, v :", d.MostWantF, v)
				break
			}
		}
	} else if d.FangXiang == "ToDown" {
		d.MostWantFloors = SliceBigToSmall(d.MostWantFloors) //从大到小排序
		for _, v := range d.MostWantFloors {
			if v < d.AtFloor {
				d.MostWantF = v
				fmt.Println("d.MostWantF, v :", d.MostWantF, v)
				break
			}
		}
	}
}

//去除mostfloor
func (d *DianTiXiang) QuChuMostFloor() {
	for i, v := range d.ToUpFloorSlice {
		if v == d.MostWantF {
			d.ToUpFloorSlice = append(d.ToUpFloorSlice[:i], d.ToUpFloorSlice[i+1:]...)
			fmt.Println("去除", d.ToUpFloorSlice)
		}
	}
	for i, v := range d.ToDownFloorSlice {
		if v == d.MostWantF {
			d.ToDownFloorSlice = append(d.ToDownFloorSlice[:i], d.ToDownFloorSlice[i+1:]...)
			fmt.Println("去除", d.ToDownFloorSlice)
		}
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
