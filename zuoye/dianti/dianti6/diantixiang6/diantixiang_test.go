package diantixiang6

import (
	"fmt"
	"learn.go/zuoye/dianti/dianti6/ren"
	"testing"
)

func TestToUpOrDown(t *testing.T) {
	var DianTi DianTiXiang
	DianTi.MaxFloor = 5 //电梯最高5层
	DianTi.MinFloor = 1 //电梯最低1层
	DianTi.AtFloor = 3  //初始电梯在1楼
	Ren := &ren.Ren{}   //&实例化，初始值，有的自动零值，有的类型不行。 //*指针类型， *ren.Ren显示其真实的值。
	//同上：Ren := &ren.Ren{ WantFangXiang: "" , ReqFloorSlice: []int{}, WantFloorSlice: []int{} }//内部元素不需要全手动实例化

	DianTi.Ren = Ren // DianTi.Ren 是指针类型的结构体，所以给它赋值一个对应的指针变量即可。注意：要实例化过
	Ren.ReqFloorSlice = []int{1, 4, 3, 2, 5, 1, 4}
	DianTi.Ren.WantFloorSlice = []int{5, 2, 1, 3, 4, 4, 2, 1} //上面一行和这样两种写法效果相同？ 貌似是的
	fmt.Println(DianTi)
	fmt.Println(DianTi.Ren.ReqFloorSlice)
	fmt.Println(DianTi.Ren.WantFloorSlice)
	DianTi.JoinUpAndDown()
	fmt.Println(DianTi.ToUpFloorSlice)
	fmt.Println(DianTi.ToDownFloorSlice)
	a := SliceQuChong([]int{4, 5, 4})
	fmt.Println(a)
}
