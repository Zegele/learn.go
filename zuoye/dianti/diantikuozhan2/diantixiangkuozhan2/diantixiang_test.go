package diantixiangkuozhan2

import (
	"fmt"
	"learn.go/zuoye/dianti/diantikuozhan2/ren"
	"testing"
)

func TestToUpOrDown(t *testing.T) {
	var DianTi DianTiXiang
	DianTi.MaxFloor = 6 //电梯最高5层
	DianTi.MinFloor = 1 //电梯最低1层
	DianTi.AtFloor = 1  //初始电梯在1楼
	Ren := &ren.Ren{}   //&实例化，初始值，有的自动零值，有的类型不行。 //*指针类型， *ren.Ren显示其真实的值。
	//同上：Ren := &ren.Ren{ WantFangXiang: "" , ReqFloorSlice: []int{}, WantFloorSlice: []int{} }//内部元素不需要全手动实例化

	DianTi.Ren = Ren                                                         // DianTi.Ren 是指针类型的结构体，所以给它赋值一个对应的指针变量即可。注意：要实例化过
	Ren.ReqFloorSlice = []int{3, 4}                                          // 无序的，但第一个数可以决定电梯初始的运行方向。
	DianTi.Ren.WantFloorSlice = []int{5, 2, 1, 1, 1, 6, 6, 6, 6, 2, 1, 1, 6} //无序的，且要去1，4楼的人数是一样多的。先去哪个人数多的楼层，电梯要做出判断。
	//1-3-6-5-4-2-1
	DianTi.DianTiYunXingFangXiang() //得到电梯运行方向 这个很重要
	fmt.Println(DianTi.FangXiang)

	DianTi.MostWant() //获得目标楼层最多的，例如 [1,4]，通过电梯运行方向，选择出[4]，优先级最高。
	fmt.Println("mostwanTfloors", DianTi.MostWantFloors)
	fmt.Println("MostWantF", DianTi.MostWantF)

	DianTi.JoinUpAndDown() //把需求和目标楼层合并，然后做出上的楼层和下的楼层，两个数组，结合电梯的方向使用。
	fmt.Println("ToUpFloorSlice:", DianTi.ToUpFloorSlice)
	fmt.Println("ToDownFloorSlice:", DianTi.ToDownFloorSlice)
	fmt.Println(DianTi.NewReqFloorSlice)
	DianTi.NewMoveTo()
	//DianTi.DianTiAtWhere()
}

func TestQuChu(t *testing.T) {
	var DianTi DianTiXiang
	DianTi.MaxFloor = 6 //电梯最高5层
	DianTi.MinFloor = 1 //电梯最低1层
	DianTi.AtFloor = 1  //初始电梯在1楼
	Ren := &ren.Ren{}
	DianTi.Ren = Ren
	DianTi.Ren.WantFloorSlice = []int{5, 2, 1, 1, 1, 6, 6, 6, 6, 2, 1, 1}
	DianTi.MostWantF = 1
	fmt.Println(DianTi.Ren.WantFloorSlice)
	DianTi.QuChuMostFloor()
	fmt.Println(DianTi.Ren.WantFloorSlice)
}
