package main

//支持任意楼request电梯，去任意层电梯。
import (
	"fmt"
	"learn.go/zuoye/dianti/diantikuozhan3/diantixiangkuozhan3"

	ren2 "learn.go/zuoye/dianti/diantikuozhan3/ren"
)

func main() {
	//做一些初始设置
	var DianTi diantixiangkuozhan3.DianTiXiang
	DianTi.MaxFloor = 6 //电梯最高5层
	DianTi.MinFloor = 1 //电梯最低1层
	DianTi.AtFloor = 1  //初始电梯在1楼
	Ren := &ren2.Ren{}  //&实例化，初始值，有的自动零值，有的类型不行。 //*指针类型， *ren.Ren显示其真实的值。
	DianTi.Ren = Ren    // DianTi.Ren 是指针类型的结构体，所以给它赋值一个对应的指针变量即可。注意：要实例化过

	//设置需求和目标
	var howManyRequest int
	fmt.Printf("有几个楼层在请求电梯？（请输入数字%d-%d）", DianTi.MinFloor, DianTi.MaxFloor)
	fmt.Scanln(&howManyRequest)
	for i := howManyRequest; i > 0; i-- {
		Ren.RequestAtWhere() //[4,3,5]楼有需求
	}

	Ren.WantGoWhereAnNiu() //人想去几楼 例如：想去5, 2, 1, 1, 1, 6, 6, 6, 6, 2, 1, 1, 6, 1楼
	//预期电梯从1楼上到3楼，再去最多人的目标层，然后按照电梯运行方向，把大家送依次送达。
	DianTi.DianTiYunXingFangXiang() //第一个请求，确定电梯运行方向
	DianTi.MostWant()               //获得人数最多的目标楼层。由于这个的影响，可能会改变电梯原本的运行方向。

	DianTi.JoinUpAndDown() //根据电梯运行方向，mostFloor等限定，整合好后续要运行的Slice。
	DianTi.NewMoveTo()     //按照整合好的Slice，进行Move。
}
