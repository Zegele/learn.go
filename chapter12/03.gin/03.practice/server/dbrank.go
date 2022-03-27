package main

import (
	"fmt"
	"gorm.io/gorm"
	"learn.go/chapter12/apiss"
	"learn.go/chapter12/frinterface"
	"log"
)

var _ frinterface.ServeInterface = &dbRank{} //记得带下划线

var _ frinterface.RankInitInterface = &dbRank{} //初始化接口中的init函数，用于初始化内嵌的数据。把数据库中有需要的数据，写入该Init()函数

func NewDbRank(conn *gorm.DB, embedRank frinterface.ServeInterface) frinterface.ServeInterface {
	if conn == nil {
		log.Fatal("数据库连接为空")
	}
	if embedRank == nil {
		log.Fatal("内存排行榜为空")
	}
	return &dbRank{ //实例化
		conn:      conn,
		embedRank: embedRank,
	}
}

// 对比这两个NewDbRank 上面的是返回一个dbRank结构体，并且该结构体已经实现了一个接口。
// 下面的NewDbRank函数是，直接返回一个dbRank结构体，没有实现接口。

//func NewDbRank(conn *gorm.DB, embedRank frinterface.ServeInterface) *dbRank {
//	if conn == nil {
//		log.Fatal("数据库连接为空")
//	}
//	if embedRank == nil {
//		log.Fatal("内存排行榜为空")
//	}
//	return &dbRank{ //实例化
//		conn:      conn,
//		embedRank: embedRank,
//	}
//}

type dbRank struct {
	conn      *gorm.DB
	embedRank frinterface.ServeInterface // 内嵌了一个接口
}

func (d *dbRank) Init() error {
	output := make([]*apiss.PersonalInformation, 0)
	resp := d.conn.Find(&output) //什么意思？？？

	if resp.Error != nil {
		fmt.Println("查找失败：", resp.Error)
		return resp.Error
	}

	for _, item := range output {
		if _, err := d.embedRank.UpdatePersonalInformation(item); err != nil {
			log.Fatalf("初始化%s时失败:%v\n", item.Name, err)
		}
	}
	return nil
}

func (d dbRank) RegisterPersonalInformation(pi *apiss.PersonalInformation) error {
	resp := d.conn.Create(pi)
	//(&PersonalInformation{
	//	Tall:   1.80,
	//	Name:   "xiao",
	//	Sex:    "男",
	//	Weight: 70,
	//	Age:    33,
	//})
	if err := resp.Error; err != nil {
		// 注意：不同企业对log有要求，比如：必须带上某个ID。log时使用公司各自的log框架。
		// e.g. https://github.com/sirupsen/logrus
		fmt.Printf("创建%s时失败：%v\n", pi.Name, err)
		return err
	}
	fmt.Printf("创建%s成功\n", pi.Name)
	_ = d.embedRank.RegisterPersonalInformation(pi) // todo handle error if there are other implementation. here we have the in-memory one
	return nil
}

func (d dbRank) UpdatePersonalInformation(pi *apiss.PersonalInformation) (*apiss.PersonalInformationFatRate, error) {
	resp := d.conn.Updates(pi)
	if err := resp.Error; err != nil {
		fmt.Printf("更新%s时失败：%v\n", pi.Name, err)
		return nil, err
	}
	fmt.Printf("更新%s成功\n", pi.Name)
	return d.embedRank.UpdatePersonalInformation(pi)

	/*
		piFr := &apiss.PersonalInformationFatRate{
			Name:    pi.Name,
			FatRate: 0,
		}

		bmi, err := gobmi.BMI(float64(pi.Weight), float64(pi.Tall))
		if err != nil {
			fmt.Println("为%s计算BMI时出错：%s\n", pi.Name, err)
			return nil, err
		}
		fr := gobmi.CalcFatRate(bmi*100, int(pi.Age), pi.Sex)
		piFr.FatRate = fr
		return piFr, nil

	*/

}

func (d dbRank) GetFatRate(name string) (*apiss.PersonalRank, error) {
	return d.embedRank.GetFatRate(name)
}

func (d dbRank) GetTop() ([]*apiss.PersonalRank, error) {
	return d.embedRank.GetTop()
}
