package diantixiang7

import (
	"fmt"
	"learn.go/zuoye/dianti/dianti7/ren"
	"testing"
)

func TestToUpOrDown(t *testing.T) {
	var DianTi DianTiXiang
	DianTi.MaxFloor = 5 //电梯最高5层
	DianTi.MinFloor = 1 //电梯最低1层
	DianTi.AtFloor = 1  //初始电梯在1楼
	Ren := &ren.Ren{}   //&实例化，初始值，有的自动零值，有的类型不行。 //*指针类型， *ren.Ren显示其真实的值。
	//同上：Ren := &ren.Ren{ WantFangXiang: "" , ReqFloorSlice: []int{}, WantFloorSlice: []int{} }//内部元素不需要全手动实例化

	DianTi.Ren = Ren                                       // DianTi.Ren 是指针类型的结构体，所以给它赋值一个对应的指针变量即可。注意：要实例化过
	Ren.ReqFloorSlice = []int{4, 3, 5}                     // 345
	DianTi.Ren.WantFloorSlice = []int{5, 2, 1, 4, 4, 2, 1} //上面一行和这样两种写法效果相同？ 貌似是的//1245 3 4 5 2 1

	DianTi.DianTiYunXingFangXiang()

	DianTi.JoinUpAndDown() //把需求和目标楼层合并，然后做出上的楼层和下的楼层，两个数组，结合电梯的方向使用。
	fmt.Println("ToUpFloorSlice:", DianTi.ToUpFloorSlice)
	fmt.Println("ToDownFloorSlice:", DianTi.ToDownFloorSlice)
	fmt.Println(DianTi)
	fmt.Println(*DianTi.Ren)
	//DianTi.MoveTo()
	DianTi.NewMoveTo()
	DianTi.DianTiAtWhere()
}
