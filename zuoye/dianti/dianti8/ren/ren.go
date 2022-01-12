package ren

import (
	"fmt"
)

type Ren struct {
	WantFangXiang  string //自己想上，还是想下 //这个应该不需要了。电梯运行里已经把方向整合好了。
	ReqFloorSlice  []int  //哪层有需求 [3] 3楼需要电梯 是乱序的，第一个决定电梯方向。决定方向后，生成一个有顺序的slice
	WantFloorSlice []int  //想去哪层 [5] 要去5楼
}

//嵌套不同包的结构体： 包名.结构体名  //跨包使用，得首字母大写

//确定人在几楼
func (r *Ren) RequestAtWhere() { //几楼需要电梯
	var renAtFloor int
	fmt.Print("你在几楼？(几楼需要电梯)")
	fmt.Scanln(&renAtFloor)
	r.ReqFloorSlice = append(r.ReqFloorSlice, renAtFloor)
}

//确定人是想上还是想下
func (r *Ren) WantUpOrDown() {
	var wantUpOrDown string
	fmt.Print("想上还是想下？（ToUp/ToDown）")
	fmt.Scanln(&wantUpOrDown)
	if wantUpOrDown == "ToUp" {
		//todo 把想UP但按低层的按钮过滤掉。
		r.WantFangXiang = wantUpOrDown
	} else if wantUpOrDown == "ToDown" {
		r.WantFangXiang = wantUpOrDown
	}
}

//确定进电梯的人，想去哪些楼层
func (r *Ren) WantGoWhereAnNiu() {
	var renWantFloor int
	for {
		fmt.Print("想去几楼？")
		fmt.Scanln(&renWantFloor)
		r.WantFloorSlice = append(r.WantFloorSlice, renWantFloor)
		if cont := ContinueOrNot(); !cont {
			break
		}
	}
}

//是否继续输入
func ContinueOrNot() bool {
	var continueOrNot string
	fmt.Print("是否录入下一个？（y/n）")
	fmt.Scanln(&continueOrNot)
	if continueOrNot == "y" || continueOrNot == "Y" {
		return true
	}
	return false
}

// at 3; want:452; req:143
