package main

import "fmt"

var _ Door = &GlassDoor{} //注意： var _ Door 这种方式,让结构体和接口对应起来，如果接口改变，就会提示结构体。
//如果不用这种方式，如果接口改变，结构体不能使用该接口， 但是不会报错，而在运行时才报错，这样不好。所以要注意。

type GlassDoor struct{}

func (d *GlassDoor) Unlock() {
	fmt.Println("GlassDoor Unlock")
}

func (d *GlassDoor) Lock() {
	fmt.Println("GlassDoor Lock")

}

func (*GlassDoor) Open() {
	fmt.Println("GlassDoor Open")
}

func (*GlassDoor) Close() {
	fmt.Println("GlassDoor Close")
}
