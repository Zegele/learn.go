package showinterface

import (
	"learn.go/zuoye/zuoye6_tizhi_dbweb/api"
)

type ServeInterface interface {
	//保存show information
	SaveShowInformation(ps *api.PersonalShow) error

	//更新show information
	UpdatePersonalInformation(ps *api.PersonalShow) error

	//查询show information
	GetShow() ([]*api.PersonalShow, error)

	GetOneShow(name string) (*api.PersonalShow, error)

	//真删除show information
	DeleteTrue(id, personID int64) error

	//假删除show information
	//DeleteFalse(name string, Id, personID int64) error
	DeleteFalse(Id, personID int64) error
}

type RankInitInterface interface {
	Init() error
}
