//Tcp客户端

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 客户端
//一个tcp客户端进行tcp通信的流程如下：
//1. 建立与服务端的连接
//2. 进行数据收发
//3. 关闭连接
//如下：

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000") //拨号，建立连接
	if err != nil {
		fmt.Println("dail fail err:", err)
		return
	}
	defer conn.Close() //关闭连接

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') //读取用户输入
		inputInfo := strings.Trim(input, "\r\n") // 去掉input前后的回车符换行符
		if strings.ToUpper(inputInfo) == "Q" {   //如果输入q就退出
			_, err = conn.Write([]byte("客户端退出")) // 发送数据
			return
		}
		_, err = conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:]) //接收服务端的数据
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

// 将上面的代码编译成clien.exe可执行文件
//先启动server.exe，再启动client.exe
//再client端输入任意内容回车，之后就能在server端看到client端发送的数据
//从而实现tcp通信
