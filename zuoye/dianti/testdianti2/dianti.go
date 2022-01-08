package testdianti2

import (
	"fmt"
	"time"
)

type Lou struct {
	maxfloor int //最大楼层
	minfloor int //最小楼层
	DianTiXiang
}
type DianTiXiang struct {
	atFloor   int    // 在哪层
	reqFloor  int    //哪层请求电梯
	fangXiang string //电梯是向上，或向下
}

func (l *Lou) DianTiAtWhere(floornum int) (atWhere int, err error) {
	if floornum > l.maxfloor || floornum < l.minfloor {
		return 0, fmt.Errorf("楼层有误！")
	}
	l.atFloor = floornum
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

func (d *DianTiXiang) MoveUpToRequest(DianTiDoor func()) (toRequest int, err error) {
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

//怎样把int转成字符串？
func (d *DianTiXiang) DianTiDoor() {
	fmt.Println("待停稳 --> 1s")
	time.Sleep(1 * time.Second)
	fmt.Println("开门 --> 3s")
	time.Sleep(3 * time.Second)
	fmt.Println("关门 --> 1s")
	time.Sleep(1 * time.Second)
}
