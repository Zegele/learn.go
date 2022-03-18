package _interface

import (
	"crypto/rand"
	"learn.go/chapter12/02.practice/apiss"
	"math/big"
)

type Interface interface {
	ReadPersonalInformation() (apiss.PersonalInfomation, error)
}

var _ Interface = &FakeInterface{}

type FakeInterface struct {
	Name       string
	BaseWeight float64
	BaseTall   float64
	BaseAge    int
	Sex        string
}

func (f *FakeInterface) ReadPersonalInformation() (apiss.PersonalInfomation, error) {
	r, _ := rand.Int(rand.Reader, big.NewInt(1000))
	out := float64(r.Int64()) / 1000

	if r.Int64()%2 == 0 {
		out = 0 - out
	}

	pi := apiss.PersonalInfomation{
		Name:   f.Name,
		Sex:    f.Sex,
		Tall:   f.BaseTall,
		Weight: f.BaseWeight, //f.BaseWeight都是指针内存的值
		Age:    f.BaseAge,
	}
	f.BaseWeight += out // 修改了指针内存的值。
	return pi, nil
}
