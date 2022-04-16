package db

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"learn.go/zuoye/bigzuoye/allinterface"
	"learn.go/zuoye/bigzuoye/api"
	"log"
)

type dbPersonInformation struct {
	connDb *gorm.DB //把这个当做数据库
}

var _ allinterface.DbPersonalInformationInterface = &dbPersonInformation{}

//初始化 该接口以及对应的结构体
func NewDbPersonInterface(connDb *gorm.DB) allinterface.DbPersonalInformationInterface {
	if connDb == nil {
		log.Fatal("数据库连接为空")
	}
	return &dbPersonInformation{
		connDb: connDb,
	}
}

// 实现 添加个人信息
func (d dbPersonInformation) SavePersonalInformation(pi *api.PersonalInformation) error {
	resp := d.connDb.Create(pi)

	if err := resp.Error; err != nil {
		fmt.Printf("创建%s失败：%v\n", pi.Name, err)
		return err
	}
	fmt.Printf("创建%s成功！\n", pi.Name)
	return nil
}

// 实现 更新个人信息
func (d dbPersonInformation) UpdatePersonalInformation(pi *api.PersonalInformation) error {
	resp := d.connDb.Updates(pi)
	if err := resp.Error; err != nil {
		fmt.Printf("更新%s时失败：%v\n", pi.Name)
		return err
	}
	fmt.Printf("更新%s成功！\n", pi.Name)
	return nil
}

// 实现 获得在线的个人信息
func (d dbPersonInformation) GetPersons() ([]*api.PersonalInformation, error) {
	output := make([]*api.PersonalInformation, 0, 20)

	resp := d.connDb.Where(&api.PersonalInformation{Onlinevisiable: true}).Find(&output)
	//注意！ .Find(&output) 是要带&符号的指针
	if resp.Error != nil {
		fmt.Println("查找失败：", resp.Error)
		return nil, resp.Error
	}

	data, err := json.Marshal(output)
	if err != nil {
		fmt.Println("marshal 出错：", err)
	}
	fmt.Printf("Online的有：%v\n", string(data))
	return output, nil
}

// 实现 删除个人信息
// 真实删除
func (d dbPersonInformation) DeletePersonalInformation(account int64) error {
	output := &api.PersonalInformation{
		Account: account,
	}

	resp := d.connDb.Delete(output)
	if err := resp.Error; err != nil {
		fmt.Printf("删除%d失败。\n", account)
		return err
	}
	fmt.Printf("该%d账户删除成功！", account)

	return nil
}
