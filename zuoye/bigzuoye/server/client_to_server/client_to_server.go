package client_to_server

import (
	"fmt"
	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learn.go/zuoye/bigzuoye/api"
	"log"
	"sync"
)

var _ api.Client_ServiceServer = &CToSServer{}

type CToSServer struct {
	sync.Mutex
	Persons map[string]*api.PersonalInformation
}

func (c *CToSServer) Register(ctx context.Context, information *api.PersonalInformation) (*api.PersonalInformation, error) {
	c.Lock()
	defer c.Unlock()
	c.Persons[information.Name] = information
	log.Printf("收到新注册人：%s\n", information.String())

	//gorm
	connDB := connectDb()
	if err := creatNewPersonToDb(connDB, information); err != nil {
		fmt.Println("存储到数据库出错：", err)
	}
	fmt.Println("存储到数据库成功！")

	return information, nil
}

//gorm
func connectDb() *gorm.DB { //返回一个数据库
	connDb, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/testdb"))
	if err != nil {
		log.Fatal("数据库链接失败：", err)
	}
	fmt.Println("连接数据库成功")
	return connDb
}

func creatNewPersonToDb(conn *gorm.DB, pi *api.PersonalInformation) error {
	resp := conn.Create(pi)
	if err := resp.Error; err != nil {
		fmt.Printf("创建%s时失败：%v\n", pi.Name, err)
		return err
	}
	fmt.Printf("创建%s成功！", pi.Name)
	return nil
}
