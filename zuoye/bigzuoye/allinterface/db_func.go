package allinterface

import "learn.go/zuoye/bigzuoye/api"

type DbPersonalInformationInterface interface {
	// 添加个人信息
	SavePersonalInformation(pi *api.PersonalInformation) error
	// 更新个人信息
	UpdatePersonalInformation(pi *api.PersonalInformation) error
	//获得在线个人信息
	GetPersons() ([]*api.PersonalInformation, error)
	//删除个人信息
	DeletePersonalInformation(account int64) error
}
