package make_fatr

import (
	"fmt"
	"learn.go/zuoye/tizhi_register/register"
	"testing"
)

func TestMakeFatR(t *testing.T) {
	//var m *register.Member
	//m.Name = "a"
	m := &register.Member{Name: "a1"}

	allm := &register.Allmember{Members: []*register.Member{m}}
	fmt.Println(allm.Members[0])

	allm = makefatr("a1")

	fmt.Println(allm.Members[0])
	//for i := 0; i < 10; i++ {
	//for i := 0; i < 10; i++ {
	//	makefatr(m)
	//	time.Sleep(1 * time.Millisecond)
	//}

}
