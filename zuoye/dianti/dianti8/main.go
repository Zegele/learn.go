package main

//支持任意楼request电梯，去任意层电梯。
import (
	"fmt"
	"learn.go/zuoye/dianti/dianti8/diantixiang8"
	ren2 "learn.go/zuoye/dianti/dianti8/ren"
)

func main() {
	//做一些初始设置
	var DianTi diantixiang8.DianTiXiang
	DianTi.MaxFloor = 5 //电梯最高5层
	DianTi.MinFloor = 1 //电梯最低1层
	DianTi.AtFloor = 1  //初始电梯在1楼
	Ren := &ren2.Ren{}  //&实例化，初始值，有的自动零值，有的类型不行。 //*指针类型， *ren.Ren显示其真实的值。
	DianTi.Ren = Ren    // DianTi.Ren 是指针类型的结构体，所以给它赋值一个对应的指针变量即可。注意：要实例化过

	//设置需求和目标
	var howManyRequest int
	fmt.Print("有多少楼层要求电梯？（请输入数字1-5）")
	fmt.Scanln(&howManyRequest)
	for i := howManyRequest; i > 0; i-- {
		Ren.RequestAtWhere() //[4,3,5]楼有需求
	}

	Ren.WantGoWhereAnNiu() //人想去几楼 人想去4，5，2，1楼
	//预期电梯从1楼上到3楼，再4，5，向下，2，1。
	fmt.Println("请求电梯的楼层：", DianTi.Ren.ReqFloorSlice)
	fmt.Println("目标楼层：", DianTi.Ren.WantFloorSlice)
	DianTi.DianTiYunXingFangXiang() //第一个请求，确定电梯运行方向
	DianTi.JoinUpAndDown()          //根据电梯运行方向，整合好运行规则。即如果向上，就向上执行完毕，再转向，反之亦然。
	//例如：[4,3,5] 和 [4,5,2,1]--> [3,4,5] [2,1]
	DianTi.NewMoveTo() //按照整合好的Slice，进行Move。
}
