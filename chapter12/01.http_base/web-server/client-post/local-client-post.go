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
	//r := strings.NewReader("http://localhost:8088/?name=xiao&&sex=男,aglajhf") //这样就可以传如任何数据
	resp, err := http.Post("http://localhost:8088", "*/*", r) //contentType种类很多。 */*表示不管是什么都接受。
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println("resp:", string(data))
}
