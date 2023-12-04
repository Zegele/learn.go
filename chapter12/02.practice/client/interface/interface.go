package _interface

import (
	"crypto/rand"
	"learn.go/chapter12/apiss"
	"math/big"
)

type Interface interface {
	ReadPersonalInformation() (apiss.PersonalInformation, error)
}

//var _ Interface = &FakeInterface{}

type FakeInterface struct {
	Name       string
	BaseWeight float64
	BaseTall   float64
	BaseAge    int
	Sex        string
}

func (f *FakeInterface) ReadPersonalInformation() (apiss.PersonalInformation, error) {
	r, _ := rand.Int(rand.Reader, big.NewInt(1000))
	out := float64(r.Int64()) / 1000

	if r.Int64()%2 == 0 {
		out = 0 - out
	}

	pi := apiss.PersonalInformation{
		Name:   f.Name,
		Sex:    f.Sex,
		Tall:   float32(f.BaseTall),
		Weight: float32(f.BaseWeight), //f.BaseWeight都是指针内存的值
		Age:    int64(f.BaseAge),
	}
	f.BaseWeight += out // 修改了指针内存的值。
	return pi, nil
}
