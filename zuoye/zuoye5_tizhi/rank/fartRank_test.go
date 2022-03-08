package rank

import (
	"fmt"
	"learn.go/zuoye/zuoye5_tizhi/pkg/apii"
	"testing"
)

func TestRank(t *testing.T) {
	person := &apii.Person{
		Name:   "a", //注意这里是逗号
		Age:    18,
		Sex:    "man",
		Tall:   1.8,
		Weight: 66,
		Fatr:   0.3,
	}
	r2 := &apii.RankItem{
		Name: "b", //注意这里是逗号
		Fatr: 0.2,
	}
	fr := &FR{}
	ps := &apii.Persons{}
	fr.appendSliceFromPerson(person, ps)
	fmt.Println("测试写入Persons:", ps) // 测试写入rankItem

	r := &apii.RankItem{}
	r = fr.fromPersonsToRankItem(person)
	fmt.Println("测试写入rankItem:", r) // 测试写入rankItem

	rs := &apii.Rank{}
	fr.appendSliceFromRank(r, rs)
	fr.appendSliceFromRank(r2, rs)
	fmt.Println("测试添加到slice", rs) // 测试添加到slice

	fr.updataRecord(rs, "a", 0.21) // 小于0.3 才会被更新 通过
	fmt.Println("测试更新：", person)
	fmt.Println(ps)
	fmt.Println(r)
	fmt.Println("rs:", rs)

	//fmt.Println("冒泡排序前：", rs.ItemsS)
	//fr.paiXuBubble(rs)
	//fmt.Println("冒泡排序后：", rs.ItemsS) //冒泡通过

	fmt.Println("快排排序前：", rs.ItemsS)
	fr.PaiXuQuick(rs, 0, len(rs.ItemsS)-1)
	fmt.Println("快排排序后：", rs.ItemsS) //快排通过

}

func TestReadFile(t *testing.T) {
	fr := &FR{}
	//fr.ReadJSONFile("E:/Geek/src/learn.go/zuoye/zuoye5_tizhi/client/huanhangdata.json")//todo  json.Unmarshal 失败
	//fr.ReadJSONFile("E:/Geek/src/learn.go/zuoye/zuoye5_tizhi/client/data.json")

	ps := fr.ReadProtoFile("E:/Geek/src/learn.go/zuoye/zuoye5_tizhi/client/data.protobuf")
	rs := fr.MakeRank(ps)
	//fr.PaiXuBubble(rs) //排序通过 test打印缺少元素

	fr.PaiXuQuick(rs, 0, len(rs.ItemsS)-1) //通过
	fmt.Println(rs)

	fr.WriteRankHuanHang("E:/Geek/src/learn.go/zuoye/zuoye5_tizhi/rankhuanhang.json", rs)
}
