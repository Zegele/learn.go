package main

// 一般接口统一在接口文件（frinterface.go）里？
import "learn.go/zuoye/zuoye5_tizhi/pkg/apii"

// 使用接口，正确录入用户信息。
type LuRu interface {
	Input(person *apii.Person) *apii.Person
}

//var LuRus = []LuRu{
//	&luru1{},
//	&luru2{},
//	&luru3{},
//}

//func ClientZhuCe() error {
//	for _, v := range LuRus {
//		if err := v(); err != nil {
//			return nil
//		}
//	}
//}
