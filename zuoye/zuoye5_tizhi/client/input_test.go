package client

import (
	"learn.go/zuoye/zuoye5_tizhi/pkg/apii"
	"testing"
	"time"
)

func TestInputData(t *testing.T) { //测试里不能测试输入？？
	input := &InputData{} //记得实例化
	registerChan := make(chan *apii.Person, 100)
	allPerson := 100
	ps := &apii.Persons{}
	for i := 0; i < allPerson; i++ {
		go input.Register(i, registerChan)
	}
	time.Sleep(1 * time.Second)
	close(registerChan)
	for luRuPerson := range registerChan {
		ps.Items = append(ps.Items, luRuPerson)
	}
	filePathjson := "E:/Geek/src/learn.go/zuoye/zuoye5_tizhi/client/data.json"
	filePathprotobuf := "E:/Geek/src/learn.go/zuoye/zuoye5_tizhi/client/data.protobuf"
	filePathjsonhuanhang := "E:/Geek/src/learn.go/zuoye/zuoye5_tizhi/client/huanhangdata.json"

	input.WriteClientByJSON(filePathjson, ps)                 //Json 不换行通过
	input.WriteClientByProtobuf(filePathprotobuf, ps)         //ProtoBuf 不换行通过
	input.WriteClientByJSONHuanHang(filePathjsonhuanhang, ps) //Json 换行通过
}
