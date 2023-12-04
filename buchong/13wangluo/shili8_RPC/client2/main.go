// 客户端进行RPC调用
package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

// 算数运算请求结构体
type ArithRequest struct {
	A int
	B int
}

// 算数运算响应结构体
type ArithResponse struct {
	Pro int // 乘积
	Que int // 商
	Rem int // 余数
}

func main() {
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln("连接错误：", err)
	}
	req := ArithRequest{11, 3}
	var res ArithResponse
	err = conn.Call("Arith.Multiply", req, &res)
	//乘法运算
	if err != nil {
		log.Fatalln("算数误差：", err)
	}
	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)
	err = conn.Call("Arith.Divide", req, &res)
	if err != nil {
		log.Fatalln("算数误差：", err)
	}
	fmt.Printf("%d / %d = %d ... %d\n", req.A, req.B, res.Que, res.Rem)

}
