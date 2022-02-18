package testdianti1

import "testing"

func TestCase1(t *testing.T) {

	var dianti DianTiXiang //电梯箱是个结构体

	var noPeople int = 0
	noPeopleRequest, err := dianti.requestDianTi(noPeople)
	if noPeopleRequest != 0 {
		t.Fatalf("预期的结果是:noPeopleRequest == 0, 但得到的结果是：nopeople == %d", noPeopleRequest)
	}
	if err != nil {
		t.Fatalf("预期的结果是:err == nil, 但得到的结果是：err == %v", err)
	}
	moveWhere, err := dianti.moveDianTi(noPeopleRequest)
	if moveWhere != "没人来，电梯不动。" {
		t.Fatalf("预期的结果是:moveWhere == 没人来，电梯不动。, 但得到的结果是：moveWhere == %s", moveWhere)
	}
	if err != nil {
		t.Fatalf("预期的结果是:err == nil, 但得到的结果是：err == %v", err)
	}

}

/*
func TestCase2(teacher testing.T) {
	onepeople, err := requestDianTi()
}
*/
