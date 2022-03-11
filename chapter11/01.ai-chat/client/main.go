package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080") //Dial 拨号  //表示给localhost:8080拨号
	if err != nil {
		log.Fatal("拨号失败：", err) //拨号失败：dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it.
	}
	defer conn.Close() // 把连接关掉
	fmt.Println("连接成功，开始聊天吧。")
	//交互
	for {
		r := bufio.NewReader(os.Stdin)
		//input, _ := r.ReadLine()//todo handle error
		input, _, _ := r.ReadLine()
		if len(input) != 0 { //防止特殊字符或换行，而让程序识别为没有完成输入。
			talk(conn, string(input))
		}
	}
	//talk(conn, "你好") //测试通过
	//talk(conn, "你是谁？")
	//talk(conn, "你是男是女？")
	//talk(conn, "今天天气怎样？")
	//talk(conn, "再见")

}

func talk(conn net.Conn, message string) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		log.Println("发送消息失败：", err)
	} else {
		data := make([]byte, 1024)
		validLen, err := conn.Read(data) // 收到（读取）服务端的回应
		if err != nil {
			log.Println("WATNING:读取服务器返回数据时出错：", err)
		} else {
			validData := data[:validLen]
			log.Println("发送：", message, "---", "服务器回复：", string(validData))
		}
	}
}
