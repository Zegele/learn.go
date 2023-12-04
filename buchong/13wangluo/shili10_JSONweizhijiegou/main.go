// go语言解码未知结构的JSON数据
// www.kancloud.cn/imdszxs/golang/1509681
package main

import (
	"encoding/json"
	"log"
	"os"
)

//go语言内置的encoding/json标准库提供了对JSON数据进行编解码的功能
//在实际开发过程中，有时候我们可能并不知道要解码的JSON数据结构是什么样子的
//这个时候应该怎么处理呢？
//如果要解码一段未知结构的JSON，只需将这段JSON数据解码输出到一个空接口即可。
//关于JSON的数据的编码和解码参考：11wenjianchuli\shili2_JSONduxie

//类型转换规则
//在前面介绍接口的时候，我们提到基于GO语言的面向对象特性
//可以通过空接口来表示任何类型
//这同样也适用于对未知结构JSON数据进行解码
//只需要将这段JSON数据解码输出到一个空接口即可
//再实际解码过程中，JSON结构里边的数据元素将如下类型转换：
//1. 布尔值将会转换为Go语言的bool类型；
//2. 数值将会转换为Go语言的float64类型；
//3. 字符串转换后还是string类型
//4. JSON数组会转换为[]interface{}类型
//5. JSON对象会转换为map[string]interface{}类型
//6. null值会转换为nil

//在Go语言标准库encoding/json中，可以使用map[string]interface{}和[]interface{}
//类型的值来分别存放未知结构的JSON对象或数组

/*
//示例：解析JSON数据，并将结果映射到空接口对象
func main() {
	u3 := []byte(`{"name":"golang", "website":"http://c.biancheng.net", "course":["GOLANG","PHP","JAVA","C"]}`)
	var user4 interface{}
	err := json.Unmarshal(u3, &user4)
	if err != nil {
		fmt.Printf("JSON解码失败：%v\n", err)
		return
	}
	fmt.Printf("JSON解码结果：%#v\n", user4)
	fmt.Printf("JSON解码结果：%v\n", user4)
}

*/

//上述代码中，user4被定义为一个空接口
//json.Unmarshal()函数将一个JSON对象u3解码到空接口user4中
//最终user4将会是一个键值对的map[string]interface{}结构
//因为u3整体上是一个JSON对象，内部属性也会遵循上述类型的转化规则进行转换

//访问解码后的数据
//要访问解码后的数据结构
//需要先判断目标结构是否为预期的数据类型
//然后我们可以通过for循环搭配range语句访问解码后的目标数据
/*
func main() {
	u3 := []byte(`{"name":"golang", "website":"http://c.biancheng.net", "course":["GOLANG","PHP","JAVA","C"]}`)
	var user4 interface{}
	err := json.Unmarshal(u3, &user4)
	if err != nil {
		fmt.Printf("JSON解码失败：%v\n", err)
		return
	}
	if user5, ok := user4.(map[string]interface{}); ok {
		//user4.(map[string]interface{}) interface{}类型判断
		for k, v := range user5 {
			switch v2 := v.(type) {
			case string:
				fmt.Println(k, "is string：", v2)
			case float64:
				fmt.Println(k, "is float64：", v2)
			case bool:
				fmt.Println(k, "is boll：", v2)
			case []interface{}:
				fmt.Println(k, "is an slice:")
				for i, iv := range v2 {
					fmt.Println(i, iv)
				}
			default:
				fmt.Println(k, "类型未知")
			}
		}
	}
}

*/

// JSON的流式读写
// go语言内置的encoding/json包还提供了Decoder和Encoder两个类型，用于支持JSON数据的流式读写
// 并提供了NewDecoder()和NewEncoder()两个函数用于具体实现：
// func NewDecoder(r io.Reader)*Decoder
// func NewEncoder(w io.Writer)*Encoder
// 从标准输入流中读取JSON数据，然后将其解码，最后再写入到标准输出流中：
func main() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)

	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}

// 执行上面的代码，我们需要先输入JSON结构数据供标准输入流os.Stdin读取
//读到数据后，会通过json.NewDecoder返回的解码器对其进行解码
//最后再通过json.NewEncoder返回的编码器将数据编码后写入标准输出流os.Stdout并打印出来：
//go run main.go
//{"name": "C语言中文网", "website": "http://c.biancheng.net/", "course": ["Golang", "PHP", "JAVA", "C"]}

//使用Decoder和Encoder对数据流进行处理可以应用得更为广泛些，
//比如读写HTTP连接，WebSocket或文件等，
//Go语言标准库中的net/rpc/jsonrpc就是一个应用了Decoder和Encoder的实际例子：
//NewServerCodec returns a new rep.ServerCodec using JSON-RPC on conn.
//func NewServerCodec(conn io.ReadWriteCloseer)rpc.ServerCodec{
//		return &serverCodec{
//			dec: json.NewDecoder(conn),
//			enc: json.NewEncoder(conn),
//			pending: make(map[uint64]*json.RawMessage),
//		}
//}
