package telnet

import (
	"fmt"
	"strings"
)

// telnet是一种协议，在操作系统中可以在命令行使用telnet命令发起TCP连接。
// 我们一般用Telnet来连接TCP服务器，键盘输入一行字符回车后，即被发送到服务器上。
// 在下例中，定义了以下两个特殊控制命令，用以实现一些功能：
// 1. 输入“@close” 退出当前连接会话
// 2. 输入“@shutdown” 终止服务器运行

func ProcessTelnetCommand(str string, exitChan chan int) bool {
	// @close 指令表示终止本次对话
	if strings.HasPrefix(str, "@close") {
		fmt.Println("Session closed")

		// 告诉外部需要断开连接
		return false

		//@shutdown指令表示终止服务进程
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("Server shutdown")

		// 往通道中写入0， 阻塞等待接受方处理
		exitChan <- 0

		// 告诉外部需要断开连接
		return false
	}

	// 打印输入的字符串
	fmt.Println(str)

	return true
}
