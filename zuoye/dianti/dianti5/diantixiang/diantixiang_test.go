package diantixiang

import (
	"fmt"
	"testing"
)

var DianTi DianTiXiang

//测试 在3楼
func TestDianTiAtWhere(t *testing.T) {
	DianTi.MaxFloor = 5
	DianTi.MinFloor = 1
	DianTi.AtFloor = 1
	atFloor, err := DianTi.DianTiAtWhere()
	if atFloor != 1 {
		t.Fatalf("预期atFloor = 1，但atFloor = %d", atFloor)
	}
	if err != nil {
		t.Fatalf("预期的结果是:err == nil, 但得到的结果是：err == %v", err)
	}
}

/*
	//3楼需求电梯 (应该放在Ren的结构图体中更好？)
	DianTi.ReqFloor = 3                            //3楼需要电梯
	dianTiFangXiang, err := DianTi.RequestDianTi() //对电梯的需求是要放在电梯箱的结构体中，还是要专门做个人的结构体才好弄？？？
	if dianTiFangXiang != "toUp" {
		teacher.Fatalf("预期 DianTi.fangXiang = toUp，但得到的结果是：DianTi.fangXiang = %s ", dianTiFangXiang)
	}
	if err != nil {
		teacher.Fatalf("预期的结果是:err == nil, 但得到的结果是：err == %v", err)
	}
*/

//电梯去3楼
func TestToReqFloor(t *testing.T) {
	DianTi.AtFloor = 1
	DianTi.ReqFloor = 3
	toreqfloor, err := DianTi.MoveToRequest() //需要atfloor, reqfloor, 需要“toUp”么？
	if toreqfloor != 3 {
		t.Fatalf("预期 target = 3，但得到的 target = %d", toreqfloor)
	}
	if err != nil {
		t.Fatalf("预期的结果是:err == nil, 但得到的结果是：err == %v", err)
	}
}

//测试 目标4，5，2 楼
func TestYunXing(t *testing.T) {
	DianTi.AtFloor = 3
	fmt.Println(DianTi)
	DianTi.TargFloorSlice = []int{4, 2, 5}
	err := DianTi.MoveToTarget() //输入目标数组 MoveToTarget--MoveUp/MoveDown
	if err != nil {
		t.Fatalf("预期的结果是:err == nil, 但得到的结果是：err == %v", err)
	}
}
