package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	r := strings.NewReader("foooooooo") //strings包要看看
	//要Reader类型的数据  把字符串，通过strings.NewReader变成了Reader类型
	//bytes.NewReader([]byte(`你好`)) // 把[]byte变成Reader类型

	resp, err := http.Post("https://www.baidu.com", "*/*", r) //contentType种类很多。 */*表示不管是什么都接受。
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}
