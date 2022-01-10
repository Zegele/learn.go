package testdianti4

import (
	"fmt"
	"time"
)

// 对结构体的理解：只要这个方法是A结构体的，该方法可以使用A结构体中的所有变量和其他方法（函数）。
type Lou struct {
	maxFloor int //最大楼层
	minFloor int //最小楼层
	DianTiXiang
}
type DianTiXiang struct {
	atFloor   int    // 在哪层
	reqFloor  int    // 哪层请求电梯
	fangXiang string // 电梯是向上，或向下
	sayFloor  string // 汇报电梯在哪层
	targFloor []int  // 目标是去几楼，按照顺序运行
}

func (l *Lou) DianTiAtWhere(d DianTiXiang) (atWhere int, err error) { //需要使用DianTiXiang结构体中的atFloor
	if d.atFloor > l.maxFloor || d.atFloor < l.minFloor {
		return 0, fmt.Errorf("楼层有误！")
	}
	l.atFloor = d.atFloor
	fmt.Println(l.atFloor)
	return l.atFloor, nil //直接继承了DianTiXiang的方法
}

func (d *DianTiXiang) RequestDianTi() (fangXiang string, err error) { //参数是结构体，要不要用指针类型？？？
	if d.reqFloor > d.atFloor { //这里没有用指针类型，合适不？？？
		d.fangXiang = "toUp"
		return d.fangXiang, nil
	} else if d.reqFloor < d.atFloor {
		d.fangXiang = "toDown"
		return d.fangXiang, nil
	} else if d.reqFloor == d.atFloor {
		d.fangXiang = "Here"
		return d.fangXiang, nil
	} else {
		return "", fmt.Errorf("什么情况？")
	}
}

//需求 —>向上
func (d *DianTiXiang) MoveUpToRequest() (toRequest int, err error) {
	if d.atFloor >= d.reqFloor {
		return 0, fmt.Errorf("不是向上的")
	}
	for d.atFloor < d.reqFloor {
		time.Sleep(1 * time.Second) //上一层需要1秒
		d.atFloor++
	}
	toRequest = d.atFloor
	d.DianTiDoor()
	return toRequest, nil
}

//需求 —>向下
func (d *DianTiXiang) MoveDownToRequest() (toRequest int, err error) {
	if d.atFloor <= d.reqFloor {
		return 0, fmt.Errorf("不是向下的")
	}
	for d.atFloor > d.reqFloor {
		time.Sleep(1 * time.Second) //上一层需要1秒
		d.atFloor--
	}
	toRequest = d.atFloor
	d.DianTiDoor()
	return toRequest, nil
}

//移动到目标
/*
func (d *DianTiXiang) MoveToTarget() (err error) {
	if d.atFloor < d.targFloor[0] {
		d.MoveUpToTarget()
	}
*/

/*
		if d.atFloor > d.targFloor[0] {
			if err := d.MoveDownToTarget(); err != nil {
				log.Printf("错了", err)
			}
			if err = d.MoveUpToTarget(); err != nil {
				log.Printf("错了", err)
			}
		} else if d.atFloor < d.targFloor[0] {
			if err = d.MoveUpToTarget(); err != nil {
				log.Printf("错了", err)
			}
			if err = d.MoveDownToTarget(); err != nil {
				log.Printf("错了", err)
			}
		}
		return nil

}
*/

//目标 -> 向上
func (d *DianTiXiang) MoveToTarget() (err error) { //4,5,2  2,1,5呢？
	for _, targ := range d.targFloor {
		if d.atFloor < targ { //3<4  4<5
			UpJiCeng := targ - d.atFloor
			for j := 1; j <= UpJiCeng; j++ {
				d.atFloor++ //4  5
				d.SayJiFloor()
			}
		} else if d.atFloor > targ {
			UpJiCeng := d.atFloor - targ
			for j := 1; j <= UpJiCeng; j++ {
				d.atFloor--
				d.SayJiFloor()
			}
		}
		d.DianTiDoor() //d.atFloor == targ 就开门，关门一套
	}
	return nil
}

//目标 -> 向下
func (d *DianTiXiang) MoveDownToTarget() (err error) {
	for _, targ := range d.targFloor {
		if d.atFloor > targ {
			UpJiCeng := d.atFloor - targ
			for j := 1; j <= UpJiCeng; j++ {
				d.atFloor--
				d.SayJiFloor()
			}
			if d.atFloor == targ {
				d.DianTiDoor()
			}
		}
	}
	return nil
}

//怎样把int转成字符串？ fmt.Sprint() 或 fmt.Sprintf() ，然后就可以和其他string类型join一起。
func (d *DianTiXiang) DianTiDoor() {
	fmt.Println("待停稳 --> 1s")
	time.Sleep(1 * time.Second)
	fmt.Println("开门 --> 3s")
	time.Sleep(3 * time.Second)
	fmt.Println("关门 --> 1s")
	time.Sleep(1 * time.Second)

}

func (d *DianTiXiang) SayJiFloor() (sayFloor string) {
	time.Sleep(1 * time.Second) //1秒过一层
	d.sayFloor = "到" + fmt.Sprintf(" %d ", d.atFloor) + "楼了"
	fmt.Println(d.sayFloor)
	return d.sayFloor
}
