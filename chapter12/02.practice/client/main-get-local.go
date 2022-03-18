package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	_interface "learn.go/chapter12/02.practice/client/interface"
	"log"
	"net/http"
	"time"
)

type frClient struct {
	handRing _interface.Interface
}

func (f frClient) register() {
	pi, _ := f.handRing.ReadPersonalInformation() // 创建了person信息
	data, _ := json.Marshal(pi)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://localhost:8080/register", "application/json", r)
	if err != nil {
		log.Println("WARNING: register fails:", err)
		return
	}
	if resp.Body != nil {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("返回：", string(data))
	}
}

func (f frClient) update() {
	pi, _ := f.handRing.ReadPersonalInformation() //自动生成了一个person信息
	data, _ := json.Marshal(pi)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://localhost:8080/personalinfo", "application/json", r)
	if err != nil {
		log.Println("WARNING: register fails:", err)
		return
	}
	if resp.Body != nil {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("返回：", string(data))
	}
}

func main() {
	frCli := &frClient{handRing: &_interface.FakeInterface{
		//Name:       fmt.Sprintf("夏普%s", time.Now()),
		Name:       fmt.Sprintf("夏普%d", time.Now().UnixNano()),
		BaseWeight: 65.0,
		BaseTall:   1.8,
		BaseAge:    33,
		Sex:        "男",
	}}

	frCli.register()
	for {
		frCli.update()
		time.Sleep(1 * time.Second)
	}
}
