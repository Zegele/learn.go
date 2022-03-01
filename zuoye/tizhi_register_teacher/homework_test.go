package tizhi_register_teacher

import "testing"

type Rank interface {
	UpdateFR(name string, fr float64)
	GetRank(name string) int
}

type Client interface {
	UpdateFR(name string, fr float64)
	GetRank(name string) int
}

func TestHomework(t *testing.T) {
	var rank Rank //实例化自己完成

	var clients []Client //实例化自己完成

	for i := 0; i < len(clients); i++ {
		go func(idx int) { //todo 提取函数
			// todo add context to control exit
			go func() { //todo 提取函数
				for { // 不停更新小强的体脂
					clients[idx].UpdateFR("小强", 0.23) //0.23 to be replaced with base +- delta
				}
			}()
			go func() { //todo 提取函数
				for { // 不停获取小强的排名
					clients[idx].GetRank("小强")
				}
			}()
		}(i)
	}

}
