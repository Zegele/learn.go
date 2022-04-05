package frinterface

import (
	"learn.go/chapter13/02.4.grpc_fatr_shizhan/apis"
)

type ServeInterface interface {
	//嫁接要求
	RegisterPersonalInformation(pi *apis.PersonalInformation) error

	UpdatePersonalInformation(pi *apis.PersonalInformation) (*apis.PersonalInformationFatRate, error)

	GetFatRate(name string) (*apis.PersonalRank, error)

	GetTop() ([]*apis.PersonalRank, error)
}

type RankInitInterface interface {
	Init() error
}
