package api

func (*PersonalShow) TableName() string { // 手动维护 要连接到数据库的哪个table，所以是用来找到对应的表
	return "show"
}
