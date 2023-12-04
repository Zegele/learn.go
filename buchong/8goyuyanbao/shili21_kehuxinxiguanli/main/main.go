package main

import (
	"learn.go/buchong/8goyuyanbao/shili21_kehuxinxiguanli/service"
)

func main() {
	//在main函数中，创建一个customerView，并运行显示主菜单。
	customerView := customerView{
		key:  "",
		loop: true,
	}
	// 这里完成对customerView结构体的customerService字段的初始化
	customerView.customerService = service.NewCustomerService()
	//显示主菜单
	customerView.mainMenu()
}
