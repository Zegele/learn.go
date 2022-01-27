package main

import (
	"fmt"
	"sync"
	"time"
)

type WebServerV1 struct {
	config     Config
	configLock sync.RWMutex
}

func (ws *WebServerV1) reload() { //加载文件或内容
	ws.configLock.Lock()
	defer ws.configLock.Unlock()
	ws.config.Content = fmt.Sprintf("%d", time.Now().UnixNano())
}

func (ws *WebServerV1) ReloadWorker() {
	for {
		time.Sleep(10 * time.Millisecond)
		ws.reload()
	}
}

func (ws *WebServerV1) Visit() string { //访问文件或内容
	ws.configLock.RLock()
	defer ws.configLock.RUnlock()
	return ws.config.Content
}
