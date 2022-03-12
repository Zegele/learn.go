package main

import (
	"encoding/json"
	"fmt"
	"learn.go/chapter11/02.practice/client/interface"
	"log"
	"net"
	"time"
)

func main() {

	var input interFace.Interface = &interFace.FakeInterface{
		Name:       "xiao",
		BaseWeight: 71.0,
		BaseTall:   1.8,
		BaseAge:    30,
		Sex:        "男",
	}

	for {
		func() {
			conn, err := net.Dial("tcp", "localhost:8080")
			if err != nil {
				log.Fatal("拨号失败：", err)
			}
			defer conn.Close()
			fmt.Println("连接成功，开始发送数据：")
			pi, err := input.ReadPersonalInformation()
			fmt.Println(pi.Weight)
			if err != nil { //说明没读取到
				log.Println("WRANING:读取失败，请重新录入：", err)
				return
			}
			data, err := json.Marshal(pi)
			if err != nil {
				log.Println("无法编码个人信息：", err)
				return
			}
			log.Println("读取到的数据：", string(data))
			talk(conn, string(data))
		}()
		time.Sleep(1 * time.Second)
	}

}

func talk(conn net.Conn, message string) { //发送message，并获得返回信息。
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
