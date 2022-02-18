package testdianti3

import (
	"testing"
)

func TestCase3(t *testing.T) {

	var lou Lou //Lou是个结构体,有最高层，有最低层，有电梯箱结构体？有人？
	var DianTi DianTiXiang

	lou.maxFloor = 5
	lou.minFloor = 1
	lou.DianTiXiang = DianTi

	DianTi.atFloor = 3
	DianTi.targFloor = []int{4, 5, 2}
	//看看目前电梯在几楼
	lookDianTi, err := lou.DianTiAtWhere(DianTi)
	if lookDianTi != 3 {
		t.Fatalf("预期 lookDianTi = 3。 但得到的 lookDianti = %d", lookDianTi)
	}
	if err != nil {
		t.Fatalf("预期的结果是:err == nil, 但得到的结果是：err == %v", err)
	}

	//关门后，依次去目标4，5，2楼
	err = DianTi.MoveToTarget() //输入目标数组 MoveToTarget--MoveUp/MoveDown
	if err != nil {
		t.Fatalf("预期的结果是:err == nil, 但得到的结果是：err == %v", err)
	}

	/* case2中的内容，case3暂时不需要
	//3楼需求电梯
	DianTi.reqFloor = 3                            //3楼需要电梯
	dianTiFangXiang, err := DianTi.RequestDianTi() //对电梯的需求是要放在电梯箱的结构体中，还是要专门做个人的结构体才好弄？？？
	if dianTiFangXiang != "toUp" {
		teacher.Fatalf("预期 DianTi.fangXiang = toUp，但得到的结果是：DianTi.fangXiang = %s ", dianTiFangXiang)
	}
	if err != nil {
		teacher.Fatalf("预期的结果是:err == nil, 但得到的结果是：err == %v", err)
	}

	//电梯去3楼
	toreqfloor, err := DianTi.MoveUpToRequest(DianTi.DianTiDoor) //需要atfloor, reqfloor, 需要“toUp”么？
	if toreqfloor != 3 {
		teacher.Fatalf("预期 target = 3，但得到的 target = %d", toreqfloor)
	}
	if err != nil {
		teacher.Fatalf("预期的结果是:err == nil, 但得到的结果是：err == %v", err)
	}

	*/
}
