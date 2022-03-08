package main

import (
	"learn.go/zuoye/zuoye5_tizhi/calculate"
	"learn.go/zuoye/zuoye5_tizhi/client"
	"learn.go/zuoye/zuoye5_tizhi/rank"
	"time"

	"learn.go/zuoye/zuoye5_tizhi/pkg/apii"
)

func main() {
	//各种初始化
	input := &client.InputData{} //记得实例化
	cal := &calculate.Calc{}
	registerChan := make(chan *apii.Person, 100)
	allPerson := 100
	ps := &apii.Persons{}
	filePathHuanHang := "E:/Geek/src/learn.go/zuoye/zuoye5_tizhi/itemhuanhang.json"
	filePathBuHuanHang := "E:/Geek/src/learn.go/zuoye/zuoye5_tizhi/itembuhuanhang.json"
	filePathProtobuf := "E:/Geek/src/learn.go/zuoye/zuoye5_tizhi/itembuhuanhang.protobuf"

	fr := &rank.FR{}

	// 注册信息和计算体脂率
	for i := 0; i < allPerson; i++ { //client生成 apii.Person
		go cal.CalcFatRate(input.Register(i, registerChan)) //input.Register 注入基本信息， cal.CalcFatRate计算fatr（体脂率），完善apii.Person整体信息。
	}
	time.Sleep(1 * time.Second)
	close(registerChan)
	for luRuPerson := range registerChan { //生成 apii.Persons
		ps.Items = append(ps.Items, luRuPerson)
	}

	// 保存client的基本信息
	input.WriteClientByJSONHuanHang(filePathHuanHang, ps) // 把client的基本信息写入（保存）到文件。(换行)
	input.WriteClientByJSON(filePathBuHuanHang, ps)       // 把client的基本信息写入（保存）到文件。
	input.WriteClientByProtobuf(filePathProtobuf, ps)     // 把client的基本信息写入（保存）到文件。
	//input.WriteClientByJSON(filePathBuHuanHang, ps)
	// 排序

	unmarshalPs := fr.ReadProtoFile(filePathProtobuf)
	// 排序阶段
	rs := fr.MakeRank(unmarshalPs)
	//fr.PaiXuBubble(rs)
	fr.PaiXuQuick(rs, 0, len(rs.ItemsS)-1)

	RankFileWritePathByJSON := "E:/Geek/src/learn.go/zuoye/zuoye5_tizhi/rankhuanhang.json"
	RankFileWritePathByPROTOBUF := "E:/Geek/src/learn.go/zuoye/zuoye5_tizhi/rank.protobuf"
	fr.WriteRank(RankFileWritePathByPROTOBUF, rs)
	fr.WriteRankHuanHang(RankFileWritePathByJSON, rs)

}
