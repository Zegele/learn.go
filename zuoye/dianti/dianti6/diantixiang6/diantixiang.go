package diantixiang6

import (
	"fmt"
	"learn.go/zuoye/dianti/dianti6/ren"
	"time"
)

type DianTiXiang struct {
	MaxFloor         int    // 最高几层 5
	MinFloor         int    // 最低几层 1
	AtFloor          int    // 电梯在哪层 3
	FangXiang        string //电梯运行方向 由第一个req决定
	SayFloor         string // "到 3 楼了" 电梯运行经过哪层，就汇报电梯在哪层
	ToUpFloorSlice   []int  // 当电梯向上运行，对接这个参数 [UpTarget] + [UpRequest] <-- 拆ren.TargetSlice 和ren.RequestSlice
	ToDownFloorSlice []int  // 当电梯向下运行，对接这个参数 [DownTarget] + [DownRequest]
	Ren              *ren.Ren
}

//参数是结构体，要不要用指针类型？？？ 如果引用的值有变动，就用指针类型的

func (d *DianTiXiang) DianTiAtWhere() (atWhere int, err error) {
	if d.AtFloor > d.MaxFloor || d.AtFloor < d.MinFloor {
		return 0, fmt.Errorf("楼层有误！")
	}
	fmt.Printf("已在 %d 楼\n", d.AtFloor)
	return d.AtFloor, nil
}

//电梯运行方向
func (d *DianTiXiang) DianTiYunXingFangXiang() {
	if d.AtFloor > d.Ren.ReqFloorSlice[0] {
		d.FangXiang = "ToDown"
	} else if d.AtFloor < d.Ren.ReqFloorSlice[0] {
		d.FangXiang = "ToUp"
	} else {
		d.FangXiang = "不动"
	}
}

//整合requestSlice和TargetSlice
func (d *DianTiXiang) JoinUpAndDown() {
	for _, valu := range d.Ren.ReqFloorSlice {
		if valu > d.AtFloor {
			d.ToUpFloorSlice = append(d.ToUpFloorSlice, valu)
		} else if valu < d.AtFloor {
			d.ToDownFloorSlice = append(d.ToDownFloorSlice, valu)
		}
	}
	for _, valu := range d.Ren.WantFloorSlice {
		if valu > d.AtFloor {
			d.ToUpFloorSlice = append(d.ToUpFloorSlice, valu)
		} else if valu < d.AtFloor {
			d.ToDownFloorSlice = append(d.ToDownFloorSlice, valu)
		}
	}

	//去重Up
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

/* 这个去重有问题，改良的在上面
func SliceQuChong(yuanShiSlice []int) (quChongSlice []int) {
	for i := 0; i < len(yuanShiSlice); i++ {
		for j := i + 1; j < len(yuanShiSlice); j++ {
			if yuanShiSlice[i] == yuanShiSlice[j] {
				yuanShiSlice = append(yuanShiSlice[:j], yuanShiSlice[j+1:]...)
			}
		}
	}
	quChongSlice = yuanShiSlice
	return quChongSlice
}

*/

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
//直接简化为Move一个动作
func (d *DianTiXiang) MoveTo() {
	if d.FangXiang == "ToUp" {
		d.MoveToUp()
		d.FangXiang = "ToDown"
		d.MoveToDown()

	} else if d.FangXiang == "ToDown" {
		d.MoveToDown()
		d.FangXiang = "ToUp"
		d.MoveToUp()
	}
	d.DianTiDoor()
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

/*
func (d *DianTiXiang) MoveTo() {
	if d.FangXiang == "ToUp" {
		for _, upFloor := range d.ToUpTargetFloorSlice {
			chaJiLou := upFloor - d.AtFloor
			for i := 0; i < chaJiLou; i++ {
				d.AtFloor++
				d.SayJiFloor()
			}
		}
		d.ToUpTargetFloorSlice = []int{}
	} else if d.FangXiang == "ToDown" {
		for _, downFloor := range d.ToDownTargetFloorSlice {
			chaJiLou := d.AtFloor - downFloor
			for i := 0; i < chaJiLou; i++ {
				d.AtFloor--
				d.SayJiFloor()
			}
		}
		d.ToDownTargetFloorSlice = []int{}
	}
	d.DianTiDoor()
}
*/

/*
//移动到req地方
func (d *DianTiXiang) MoveToRequest() (reqFloor int, err error) {
	if d.FangXiang == "ToDown" {
		for chaJiLou := d.AtFloor - d.ReqFloorSlice[0]; chaJiLou > 0; chaJiLou-- {
			d.AtFloor--
			d.SayJiFloor()
		}
	} else if d.FangXiang == "ToUp" {
		for chaJiLou := d.ReqFloorSlice[0] - d.AtFloor; chaJiLou > 0; chaJiLou-- {
			d.AtFloor++
			d.SayJiFloor()
		}
	}
	d.DianTiDoor() //d.AtFloor == d.ReqFloor
	d.ReqFloorSlice = d.ReqFloorSlice[1:]
	return d.AtFloor, nil
}

*/

/*
func (d *DianTiXiang) MoveToRequest1() {
	if d.AtFloor > d.ReqFloor {
		d.MoveDownToRequest()
	} else if d.AtFloor < d.ReqFloor {
		d.MoveUpToRequest()
	}
	d.DianTiDoor() //d.atFloor == d.reqFloor 就直接开门
}

//需求 —>向上
func (d *DianTiXiang) MoveUpToRequest() (toRequest int, err error) {
	for d.AtFloor < d.ReqFloor {
		time.Sleep(1 * time.Second) //上一层需要1秒
		d.AtFloor++
		d.SayJiFloor()
	}
	toRequest = d.AtFloor
	d.DianTiDoor()
	return toRequest, nil
}

//需求 —>向下
func (d *DianTiXiang) MoveDownToRequest() (toRequest int, err error) {
	for d.AtFloor > d.ReqFloor {
		time.Sleep(1 * time.Second) //上一层需要1秒
		d.AtFloor--
		d.SayJiFloor()
	}
	toRequest = d.AtFloor
	d.DianTiDoor()
	return toRequest, nil
}

*/

//移动到目标(总)
//移动到目标
/*
func (d *DianTiXiang) MoveToTarget() (err error) { //slice要排好顺序，是按照slice的顺序运行的。
	for _, targ := range d.TargFloorSlice {
		if d.AtFloor < targ {
			UpJiCeng := targ - d.AtFloor
			for j := 1; j <= UpJiCeng; j++ {
				d.AtFloor++
				d.SayJiFloor()
			}
		} else if d.AtFloor > targ {
			UpJiCeng := d.AtFloor - targ
			for j := 1; j <= UpJiCeng; j++ {
				d.AtFloor--
				d.SayJiFloor()
			}
		}
		d.DianTiDoor() //d.atFloor == targ 就执行开门，关门一套
	}
	return nil
}

*/

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

//at: 3 方向：Up， want[4,5,2,1]，wantUp[4,5],wantDown[2,1]//先从小到大,然后从大到小， req[1，5，4，2] req1[4，5] req2[2,1]
//at: 3 方向：Up， want[1，2]，want1[4,5],want2[2,1]//先从小到大,然后从大到小， req[1，5，4，2] req1[4，5] req2[2,1]
//at: 3 方向：Down， want[4,5,2,1]，want1[2,1],want2[4,5]//先从大到小，然后从小到大 req[1，5，4，2] req1[4，5] req2[2,1]
