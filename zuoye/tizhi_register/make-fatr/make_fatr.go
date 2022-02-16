package make_fatr

import (
	"fmt"
	"learn.go/zuoye/tizhi_register/register"
	"math/rand"
	"time"
)

func makefatr(m *register.Member) (allm *register.Allmember) {
	rand.Seed(time.Now().UnixNano()) //参数类型是int64
	r := rand.Intn(39)
	rToF := (float64(r) + 1) / 100 //加1是避免随机出现0的情况。
	m.FatR = rToF

	fmt.Println("体脂是：", m.FatR)
	fmt.Println(allm.Members[0])
	return allm
}

//func makerank(m *register.Member) (allm *register.Allmember) {
//
//}
