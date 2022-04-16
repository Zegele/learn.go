package api

func (*PersonalInformation) TableName() string { // 手动维护 看对接数据库中的哪个table
	return "client_personalinformation"
}
