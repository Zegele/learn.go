package frinterface

import (
	"learn.go/chapter12/apiss"
)

type ServeInterface interface {
	//嫁接要求
	RegisterPersonalInformation(pi *apiss.PersonalInformation) error

	UpdatePersonalInformation(pi *apiss.PersonalInformation) (*apiss.PersonalInformationFatRate, error)

	GetFatRate(name string) (*apiss.PersonalRank, error)

	GetTop() ([]*apiss.PersonalRank, error)
}

type RankInitInterface interface {
	Init() error
}
