package dbshow

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"learn.go/zuoye/zuoye6_tizhi_dbweb/api"
	"learn.go/zuoye/zuoye6_tizhi_dbweb/showinterface"
	"log"
)

type dbShow struct {
	conn *gorm.DB //把这个当做数据库
}

var _ showinterface.ServeInterface = &dbShow{}

func NewDbShow(conn *gorm.DB) showinterface.ServeInterface {
	if conn == nil {
		log.Fatal("数据库连接为空")
	}
	return &dbShow{
		conn: conn,
	}
}

func (d dbShow) SaveShowInformation(ps *api.PersonalShow) error {
	resp := d.conn.Create(ps)

	if err := resp.Error; err != nil {
		fmt.Printf("创建%s信息失败：%v\n", ps.Name, err)
		return err
	}
	fmt.Printf("创建%s成功！\n", ps.Name)
	return nil
}

func (d dbShow) UpdatePersonalInformation(ps *api.PersonalShow) error {
	resp := d.conn.Updates(ps)
	if err := resp.Error; err != nil {
		fmt.Printf("更新%s时失败：%v\n", ps.Name, err)
		return err //这样原始的数据会成nil么？？？
	}
	fmt.Printf("更新%s成功\n", ps.Name)
	return nil
}

func (d dbShow) GetShow() ([]*api.PersonalShow, error) {
	//idmax := 200
	output := make([]*api.PersonalShow, 0, 20)

	resp := d.conn.Where(&api.PersonalShow{Visiable: true}).Find(&output) // 我预期是找到所有的,可见的。未测试
	//注意！ .Find(&output) 是要带&符号的指针
	if resp.Error != nil {
		fmt.Println("查找失败：", resp.Error)
		return nil, resp.Error
	}

	data, err := json.Marshal(output)
	if err != nil {
		fmt.Println("marshal 出错：", err)
	}
	fmt.Printf("Show圈：%v\n", string(data))

	return output, nil
}

func (d dbShow) GetOneShow(name string) (*api.PersonalShow, error) {
	ps := &api.PersonalShow{}
	resp := d.conn.Where(&api.PersonalShow{Name: name}).Find(&ps)
	if resp.Error != nil {
		fmt.Println("查找失败：", resp.Error)
		return nil, resp.Error
	}

	data, err := json.Marshal(ps)
	if err != nil {
		fmt.Println("marshal 出错：", err)
	}
	fmt.Printf("Show圈：%v\n", string(data))

	return ps, nil
}

func (d dbShow) DeleteTrue(id, personID int64) error { //删除这个人的，第n条状态。
	output := &api.PersonalShow{Id: id} //这一句干嘛的？？不能删了？
	//resp := d.conn.Where(&api.PersonalShow{Name: name}).Find(&output) // 我预期是找到所有的。未测试
	//
	//if resp.Error != nil {
	//	fmt.Println("查找失败：", resp.Error)
	//	return resp.Error
	//}

	resp := d.conn.Delete(output)
	if err := resp.Error; err != nil {
		fmt.Printf("删除%d失败。\n", id)
		return err
	}
	fmt.Printf("%d的第%d条状态删除成功！\n", id, personID)

	return nil
}

//如果有name，在网页上怎么弄name和id两个参数？所以先把name去掉，跑通再说。
//func (d dbShow) DeleteFalse(name string, Id, personID int64) error {
//	upD := &api.PersonalShow{
//		Id:       Id, //是通过Id定位的
//		Visiable: false,
//	}
//
//	resp := d.conn.Model(upD).Select("visiable").Updates(upD) // 预期更新为不可见。其他数据不变
//
//	if resp.Error != nil {
//		fmt.Println("不可视失败：", resp.Error)
//		return resp.Error
//	}
//
//	fmt.Printf("假删除！！！ %s的第%d条状态删除成功！\n", name, personID)
//	return nil
//}
func (d dbShow) DeleteFalse(Id, personID int64) error {
	upD := &api.PersonalShow{
		Id:       Id, //是通过Id定位的
		Visiable: false,
	}

	resp := d.conn.Model(upD).Select("visiable").Updates(upD) // 预期更新为不可见。其他数据不变

	if resp.Error != nil {
		fmt.Println("不可视失败：", resp.Error)
		return resp.Error
	}

	fmt.Printf("假删除！！！ %d的第%d条状态删除成功！\n", Id, personID)
	return nil
}
