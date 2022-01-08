package testdianti2

//注意：文件名必须_test.go结尾  我之前dianti_test2.go是有问题的。
//注意：文件名不要有数字。
import (
	"testing"
)

func TestCase2(t *testing.T) {

	var lou Lou //Lou是个结构体,有最高层，有最低层，有电梯箱结构体？有人？
	var dianTi DianTiXiang

	lou.maxfloor = 5
	lou.minfloor = 1
	lou.DianTiXiang = dianTi

	//看看目前电梯在几楼
	lookDianTi, err := lou.DianTiAtWhere(1)
	if lookDianTi != 1 {
		t.Fatalf("预期 lookDianTi = 1。 但得到的 lookDianti = %d", lookDianTi)
	}
	if err != nil {
		t.Fatalf("预期的结果是:err == nil, 但得到的结果是：err == %v", err)
	}

	//3楼需求电梯
	dianTi.reqFloor = 3                            //3楼需要电梯
	dianTiFangXiang, err := dianTi.RequestDianTi() //对电梯的需求是要放在电梯箱的结构体中，还是要专门做个人的结构体才好弄？？？
	if dianTiFangXiang != "toUp" {
		t.Fatalf("预期 dianTi.fangXiang = toUp，但得到的结果是：dianTi.fangXiang = %s ", dianTiFangXiang)
	}
	if err != nil {
		t.Fatalf("预期的结果是:err == nil, 但得到的结果是：err == %v", err)
	}

	//电梯去3楼
	toreqfloor, err := dianTi.MoveUpToRequest(dianTi.DianTiDoor) //需要atfloor, reqfloor, 需要“toUp”么？
	if toreqfloor != 3 {
		t.Fatalf("预期 target = 3，但得到的 target = %d", toreqfloor)
	}
	if err != nil {
		t.Fatalf("预期的结果是:err == nil, 但得到的结果是：err == %v", err)
	}
}
