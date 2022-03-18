package frinterface

import "learn.go/chapter12/02.practice/apiss"

type ServeInterface interface {
	//嫁接要求
	RegisterPersonalInformation(pi *apiss.PersonalInfomation) error

	UpdatePersonalInformation(pi *apiss.PersonalInfomation) (*apiss.PersonalInfomationFatRate, error)

	GetFatRate(name string) (*apiss.PersonalRank, error)

	GetTop() ([]*apiss.PersonalRank, error)
}
