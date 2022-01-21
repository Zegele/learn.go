package main

type Assets struct {
	assets []Asset //能这么用？？？Assets是一个结构体
	// 里面有个元素assets 是Assets结构体类型的切片类型。
}

func (a *Assets) DoStartWork() {
	for _, item := range a.assets {
		if d, ok := item.(Door); ok { // 判断item是不是Door类型？ Door 是接口类型
			d.Unlock()
			d.Open()
		}
	}
}

func (a *Assets) DoStopWork() {
	for _, item := range a.assets {
		if d, ok := item.(Door); ok { //断言item是否被Door接口实现？？？ 还有例子：i.(&struct) 意思是struct是不是实现了这个接口？
			//注意这种 item.(Door) 的用法
			// 所以适用于是双向的？ 可以断言这个接口是否实现了该对象（结构体，或其他变量类型）。也可以断言这个对象（结构体）是否被这个接口实现了。
			d.Close()
			d.Lock()
		}
	}
}
