package main

import (
	"fmt"
	"learn.go/zuoye/dianti/dianti6/diantixiang6"
	"learn.go/zuoye/dianti/dianti6/ren"
)

func main() {
	var DianTi diantixiang6.DianTiXiang
	DianTi.MaxFloor = 5 //电梯最高5层
	DianTi.MinFloor = 1 //电梯最低1层
	DianTi.AtFloor = 1  //初始电梯在1楼
	Ren := &ren.Ren{}   //&实例化，初始值，有的自动零值，有的类型不行。 //*指针类型， *ren.Ren显示其真实的值。
	//同上：Ren := &ren.Ren{ WantFangXiang: "" , ReqFloorSlice: []int{}, WantFloorSlice: []int{} }//内部元素不需要全手动实例化

	DianTi.Ren = Ren // DianTi.Ren 是指针类型的结构体，所以给它赋值一个对应的指针变量即可。注意：要实例化过
	/*效果同上
	  Ren := ren.Ren{} 不同点在于，这个Ren不是指针类型
	  	DianTi.Ren = &Ren
	*/
	/*
	   思考 做一做
	      	//var r ren.Ren
	      	//var Ren *ren.Ren 这种定义方法怎么实例化结构体？？？
	*/
	{
		Ren.RequestAtWhere()   //人在几楼 //输入3 电梯从1楼到3楼
		Ren.WantUpOrDown()     //人想上还是想下 电梯向上走//人要向上，如果人要向下，则忽略
		Ren.WantGoWhereAnNiu() //人想去几楼 人想去4，5，2楼，电梯先4，再5，再向下去2。
	}
	/* 效果和上面是一样的
	{
		DianTi.Ren.RequestAtWhere()
		DianTi.Ren.WantUpOrDown()
		DianTi.Ren.WantGoWhereAnNiu()
	}
	*/
	fmt.Println(Ren)
	fmt.Println(*Ren)
	fmt.Println(DianTi)
}
