package ren

import (
	"fmt"
	"testing"
)

func TestReqFloorSlice(t *testing.T) {
	var ren Ren
	//1，2，3层需要电梯
	req1 := 1
	ren.ReqFloorSlice = append(ren.ReqFloorSlice, req1)
	if ren.ReqFloorSlice[0] != 1 {
		t.Fatalf("预期 ReqFloorSlice[0] = 1，但得到的是:%d", ren.ReqFloorSlice[0])
	}
	req2 := 2
	ren.ReqFloorSlice = append(ren.ReqFloorSlice, req2)
	req3 := 3
	ren.ReqFloorSlice = append(ren.ReqFloorSlice, req3)
	fmt.Println(ren)

}
