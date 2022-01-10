package ren

import (
	"fmt"
)

type Ren struct {
	WantFangXiang  string
	ReqFloorSlice  []int //哪层有需求 [3] 3楼需要电梯
	WantFloorSlice []int //想去哪层 [5] 要去5楼
}

//嵌套不同包的结构体： 包名.结构体名  //跨包使用，得首字母大写

//确定人在几楼
func (r *Ren) RequestAtWhere() {
	var renAtFloor int
	fmt.Print("你在几楼？")
	fmt.Scanln(&renAtFloor)
	r.ReqFloorSlice = append(r.ReqFloorSlice, renAtFloor)
}

//确定人是想上还是想下
func (r *Ren) WantUpOrDown() {
	var wantUpOrDown string
	fmt.Print("想上还是想下？（ToUp/ToDown）")
	fmt.Scanln(&wantUpOrDown)
	if wantUpOrDown == "ToUp" {
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
		if cont := r.ContinueOrNot(); !cont {
			break
		}
	}
}

//是否继续输入要去几楼
func (r *Ren) ContinueOrNot() bool {
	var continueOrNot string
	fmt.Print("是否继续输入楼层（y/n）:")
	fmt.Scanln(&continueOrNot)
	if continueOrNot == "y" || continueOrNot == "Y" {
		return true
	}
	return false
}

// at 3; want:452; req:143
