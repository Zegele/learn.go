package main

type Assets struct {
	assets []Asset //能这么用？？？Assets是一个结构体
	// 里面有个元素assets 是Assets结构体类型的切片类型。
}

func (a *Assets) DoStartWork() {
	for _, item := range a.assets {
		if d, ok := item.(Door); ok {
			d.Unlock()
			d.Open()
		}
	}
}

func (a *Assets) DoStopWork() {
	for _, item := range a.assets {
		if d, ok := item.(Door); ok {
			d.Close()
			d.Lock()
		}
	}
}
