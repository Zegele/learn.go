package session

import (
	"bufio"
	"fmt"
	"learn.go/buchong/9bingfa/shili19_telnet/telnet"
	"net"
	"strings"
)

// www.likecs.com/show-306390433.html
// 回音服务器的基本逻辑是“收到什么，返回什么”，reader.ReadString可以一直读取Socket连接中的数据，直到等到期望的结尾符。
// 这种期望的结尾符，也叫定界符，一般用于将TCP封包中的逻辑数据拆分开。

// 下例使用定界符是回车换行符（"\r\n"） ，HTTP协议也是使用同样的定界符。
// 使用reader.ReadString() 函数可以将封包简单地拆分开。

// 回音服务器需要将收到的有效数据通过socket发送回去

// 连接的会话逻辑
func HandleSession(conn net.Conn, exitChan chan int) {
	// 会话入口，传入连接和退出用的通道。 HandleSession()函数被服务端并发执行。
	fmt.Println("Session started:")

	// 创建一个网络连接数据的读取器
	reader := bufio.NewReader(conn)
	// 使用bufio包的NewReader()方法，创建一个网络数据读取器，这个Reader将输入数据的读取过程进行封装，方便我们迅速获取到需要的数据。

	// 接受数据的循环
	for {
		// 会话开始
		// 读取字符串，直到碰到回车返回
		str, err := reader.ReadString('\n')
		// 通过reader读取器读取封包，处理封包后，需要继续读取从网络发送过来的下一个封包。因此需要一个会话处理循环。
		// 使用reader.ReadString()方法进行封包读取。内部会自动处理粘包过程，直到下一个回车符到达后返回数据。

		// 数据读取正确
		if err == nil {

			// 这里认为封包来自Telnet，每个指令以回车换行符("\r\n")结尾
			// 去掉字符串尾部的回车和空白符
			str = strings.TrimSpace(str)

			// 处理Telnet指令
			if !telnet.ProcessTelnetCommand(str, exitChan) {
				//将str字符串传入TelNet指令处理函数processsTelnetCommand()中，同时传入退出控制通道exitChan。
				// 当这个函数返回false时，表示需要关闭当前连接。
				conn.Close()
				break
			}

			//Echo逻辑， 发什么数据，原样返回
			//将有效数据通过conn的Write() 方法写入，同时在字符串尾部添加回车换行符（"\r\n"）
			// 数据将被socket发送给连接方。
			conn.Write([]byte(str + "\r\n"))
		} else {
			// 发送错误
			// 处理当reader.ReadString()函数返回错误时，打印错误信息，并关闭连接，退出会话，并结束循环
			fmt.Println("Session closed")
			conn.Close()
			break
		}
	}
}
