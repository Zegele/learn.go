package diantixiang

import (
	"fmt"
	"learn.go/zuoye/dianti/dianti5/ren"
	"time"
)

type DianTiXiang struct {
	MaxFloor  int    // 最高几层 5
	MinFloor  int    // 最低几层 1
	AtFloor   int    // 电梯在哪层 3
	FangXiang string //电梯运行方向 由第一个req决定
	//ReqFloor       int    // 哪层请求电梯 3 这个应该放在Ren结构体中？
	SayFloor       string // "到 3 楼了" 电梯运行经过哪层，就汇报电梯在哪层
	TargFloorSlice []int  // [4,5,2] 目标是去几楼，按照顺序运行
	ReqFloorSlice  []int  // [3]几楼有需求，按照顺序运行
	Ren            *ren.Ren
}

//参数是结构体，要不要用指针类型？？？ 如果引用的值有变动，就用指针类型的

func (d *DianTiXiang) DianTiAtWhere() (atWhere int, err error) {
	if d.AtFloor > d.MaxFloor || d.AtFloor < d.MinFloor {
		return 0, fmt.Errorf("楼层有误！")
	}
	fmt.Printf("电梯目前停在 %d 楼\n", d.AtFloor)
	return d.AtFloor, nil
}
func (d *DianTiXiang) GetTatget() {
	d.TargFloorSlice = d.Ren.WantFloorSlice
}

func (d *DianTiXiang) GetRequest() {
	d.ReqFloorSlice = d.Ren.ReqFloorSlice
}

func (d *DianTiXiang) MoveToRequest() (shengYuReqFloor []int, err error) {
	for d.AtFloor > d.ReqFloorSlice[0] {
		d.AtFloor--
		d.SayJiFloor()
	}
	for d.AtFloor < d.ReqFloorSlice[0] {
		d.AtFloor++
		d.SayJiFloor()
	}

	d.DianTiDoor() //d.AtFloor == d.ReqFloor
	return d.ReqFloorSlice[1:], nil
}

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

//电梯门开关
func (d *DianTiXiang) DianTiDoor() {
	fmt.Println("叮（待停稳） --> 1s")
	time.Sleep(1 * time.Second)
	fmt.Println("开门 --> 3s")
	time.Sleep(3 * time.Second)
	fmt.Println("关门 --> 1s")
	time.Sleep(1 * time.Second)
}

//电梯报楼层
func (d *DianTiXiang) SayJiFloor() (sayFloor string) {
	time.Sleep(1 * time.Second) //1秒过一层
	d.SayFloor = "到" + fmt.Sprintf(" %d ", d.AtFloor) + "楼了"
	fmt.Println(d.SayFloor)
	return d.SayFloor
}
