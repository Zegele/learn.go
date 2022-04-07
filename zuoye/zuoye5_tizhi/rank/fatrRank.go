package rank

import (
	"encoding/json"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"learn.go/zuoye/zuoye5_tizhi/pkg/apii"
	"log"
)

//排序

type FR struct {
	num int
}

//生成Prosons
func (fr *FR) appendSliceFromPerson(p *apii.Person, ps *apii.Persons) {
	ps.Items = append(ps.Items, p)
}

//从person中生成rankItem数据
func (fr *FR) fromPersonsToRankItem(p *apii.Person) (ritem *apii.RankItem) {
	ritem = &apii.RankItem{} // 需要实例化下
	ritem.Name = p.Name
	ritem.Fatr = p.Fatr
	return
}

// 从rankItem生成切片（准备排列前的切片数据）
func (fr *FR) appendSliceFromRank(p *apii.RankItem, rs *apii.Rank) {
	rs.ItemsS = append(rs.ItemsS, p)
}

// 更新体脂数据
func (fr *FR) updataRecord(rs *apii.Rank, name string, fatr float32) (p *apii.Person, ps *apii.Persons, b bool) {
	for i, item := range rs.ItemsS {
		if item.Name == name {
			if item.Fatr >= fatr {
				item.Fatr = fatr
			}
			rs.ItemsS[i] = item
			b = true
			break
		}
	}
	return p, ps, b //bool 默认是false， 如果找到name，则bool为true。否则（没有找到）则为false
}

//读取文件，排序后，存入新文件。
//func (fr *FR) ReadJSONFile(filePath string) { //json.Unmarshal()没通过
//	persons := make([]apii.Persons, 0, 1000)
//
//	data, err := ioutil.ReadFile(filePath)
//
//	if err != nil {
//		fmt.Println("读取文件失败：", err)
//		return
//	}
//	//fmt.Println("______", string(data))
//	if err := json.Unmarshal(data, &persons); err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("fgfha", persons)
//}

func (fr *FR) ReadProtoFile(filePath string) (ps *apii.Persons) {
	ps = &apii.Persons{}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("读取文件失败：", err)

	}
	if err = proto.Unmarshal(data, ps); err != nil { // 转换回结构体 &apii.Persons{} 成功
		log.Fatal(err)
	}
	return ps
}

// 批量生成rankItem --> ranks
func (fr *FR) MakeRank(ps *apii.Persons) (rs *apii.Rank) {
	rs = &apii.Rank{}
	for _, v := range ps.Items {
		fr.appendSliceFromRank(fr.fromPersonsToRankItem(v), rs)
	}
	return rs
}

//排序
//体脂率排序
//1. 冒泡排序
func (fr *FR) PaiXuBubble(rs *apii.Rank) {
	for i := 0; i < len(rs.ItemsS)-1; i++ {
		for j := 0; j < len(rs.ItemsS)-i-1; j++ {
			if (rs.ItemsS)[j].Fatr > rs.ItemsS[j+1].Fatr {
				(rs.ItemsS)[j], (rs.ItemsS)[j+1] = (rs.ItemsS)[j+1], (rs.ItemsS)[j]
			}
		}
	}
	for i, v := range rs.ItemsS {
		v.Rank = int64(i) + 1
	}
}

//2. 快排
func (fr *FR) PaiXuQuick(rs *apii.Rank, startIndex, endIdex int) {
	pivotIdx := (startIndex + endIdex) / 2
	pivotV := (rs.ItemsS)[pivotIdx].Fatr
	l, r := startIndex, endIdex
	for l <= r {
		for (rs.ItemsS)[l].Fatr < pivotV {
			l++
		}
		for (rs.ItemsS)[r].Fatr > pivotV {
			r--
		}
		if l >= r {
			break
		}
		(rs.ItemsS)[l], (rs.ItemsS)[r] = (rs.ItemsS)[r], (rs.ItemsS)[l]
		l++
		r--
	}
	if l == r {
		l++
		r--
	}

	if r > startIndex {
		fr.PaiXuQuick(rs, startIndex, r) //交错形成2个新的slice， 递归
	}
	if l < endIdex {
		fr.PaiXuQuick(rs, l, endIdex)
	}

	for i, v := range rs.ItemsS {
		v.Rank = int64(i) + 1
	}
}

//将排好序的数据写入文件
// proto 写入 不换行
func (fr *FR) WriteRank(filePath string, rs *apii.Rank) {
	data, err := proto.Marshal(rs) //proto.Marshal参数是要message类型
	if err != nil {
		log.Fatal(err)
	}
	if err = ioutil.WriteFile(filePath, data, 0777); err != nil {
		log.Fatal(err)
	}
}

// proto 写入 换行
func (fr *FR) WriteRankHuanHang(filePath string, rs *apii.Rank) {
	ranks := make([]byte, 0, 100000)

	for _, v := range rs.ItemsS {
		rank, err := json.Marshal(v)
		if err != nil {
			log.Fatal(err)
		}
		for _, vv := range rank {
			ranks = append(ranks, vv)
		}
		ranks = append(ranks, '\n')
	}

	//data, err := json.Marshal(ranks) //ranks已经成为换行好的数据。
	//if err != nil {
	//	log.Fatal(err)
	//}

	if err := ioutil.WriteFile(filePath, ranks, 0777); err != nil {
		log.Fatal(err)
	}
}
