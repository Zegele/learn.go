package main

import (
	"context"
	"google.golang.org/grpc"
	"learn.go/zuoye/bigzuoye/api"
	"learn.go/zuoye/bigzuoye/client/client_func"
	"log"
)

func main() {
	cf := client_func.NewClientFunc() //初始化 ClientFunc 结构体
	//cf := &client_func.ClientFunc{
	//	Person: &api.PersonalInformation{},
	//}
	cf.PutinNamePassword()
	cf.Register()
	//PutinNamePassword()

	conn, err := grpc.Dial("0.0.0.0:9090", grpc.WithInsecure())
	// 建立拨号：localhost:9090；  这个grpc.WithInsecure()，是强制使用http的，因为grpc是直接支持https的。如果只用http就好使用这个东西。
	// conn就是链接
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := api.NewClient_ServiceClient(conn) // 建立了客户端和服务器的连接，并是客户端的。
	// NewClient_Service是api里的标准服务器的名字，Client表示是给客户端使用的。
	// 把conn（链接） 传给该函数
	ret, err2 := c.Register(context.TODO(), cf.Person) //c.Register是注册了该服务，运行Register后,就把cf.Person发送给服务端
	// c.Register 是自动生成的
	if err2 != nil {
		log.Fatal("注册失败。", err)
	}
	log.Println("注册成功", ret)

	//生成登录文件
	filepath := "E:/Geek/src/learn.go/zuoye/bigzuoye/client/file.txt"
	cf.MakeFile(filepath)

	// 登录
}

/*
func PutinNamePassword() {
	var personalInformation = &api.PersonalInformation{}
	fmt.Println("注册个人信息")
	cmd := &cobra.Command{
		Use:   "personPutin",
		Short: "注册账号",
		Long:  "输入，昵称（name），账号（password），获得账号（account）",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("name:", personalInformation.Name)
			fmt.Println("password:", personalInformation.Password)
		},
	}
	//录入数据：
	cmd.Flags().StringVar(&personalInformation.Name, "name", "", "姓名")
	cmd.Flags().Int64Var(&personalInformation.Password, "password", 0, "密码")

	cmd.Execute()
}

*/
