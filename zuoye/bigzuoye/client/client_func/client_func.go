package client_func

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"learn.go/zuoye/bigzuoye/allinterface"
	"learn.go/zuoye/bigzuoye/api"
	"math/rand"
	"os"
	"sync"
	"time"
)

var _ allinterface.ClientServeInterface = &ClientFunc{}

type ClientFunc struct {
	sync.Mutex
	Person *api.PersonalInformation
}

func NewClientFunc() *ClientFunc {
	return &ClientFunc{
		Person: &api.PersonalInformation{},
	}
}

func (c *ClientFunc) PutinNamePassword() {
	fmt.Println("注册个人信息")
	cmd := &cobra.Command{
		Use:   "personPutin",
		Short: "注册账号",
		Long:  "输入，昵称（name），账号（password），获得账号（account）",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("name:", c.Person.Name)
			fmt.Println("password:", c.Person.Password)
		},
	}
	//录入数据：
	c.Lock()
	cmd.Flags().StringVar(&c.Person.Name, "name", "", "姓名")
	cmd.Flags().Int64Var(&c.Person.Password, "password", 0, "密码")
	c.Unlock()
	cmd.Execute()

	//MakeAccount(c)
	//fmt.Println(c.person)
}

func (c *ClientFunc) Register() error {

	if c.Person.Password > 1000000 { //todo 实现限定6位数密码
		return fmt.Errorf("密码应是6位数字")
	}
	//生成账号
	MakeAccount(c)
	fmt.Println(c.Person)

	return nil
}

//makeAccount生成账号
func MakeAccount(c *ClientFunc) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(1000000)

	//todo 对比账号库
	//如果对比没有重复的，则成功
	c.Person.Account = int64(r)
}

// MakeFile生成密码文件
func (c *ClientFunc) MakeFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法打开文件：", filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()

	bAccount, _ := json.Marshal(c.Person.Account)
	bPassword, _ := json.Marshal(c.Person.Password)
	_, err = file.Write(append(bAccount, '\n'))
	_, err = file.Write(bPassword)
}

func (c *ClientFunc) Login(account int64, password int64) {
	//TODO implement me
	panic("implement me")
}

func (c *ClientFunc) Online() error {
	//TODO implement me
	panic("implement me")
}

func (c *ClientFunc) ChatWith(account1, account2 int64) error {
	//TODO implement me
	panic("implement me")
}

func (c *ClientFunc) ChatHistory() {
	//TODO implement me
	panic("implement me")
}
