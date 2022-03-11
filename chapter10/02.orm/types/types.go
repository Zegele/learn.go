package types

type PersonalInformation struct {
	ID int64 `json:"id,omitempty" gorm:"primaryKey; column:id"` // 添加设置primaryKey
	// 测试，没有添加 primaryKey，但是结构体中加ID（primaryKey）就依然可正确执行。
	Name string `json:"name,omitempty" gorm:"column:name"`
	//有json标记时没有omitempty标记，序列化时将使用配置的json名称(字段大写时)
	//有json标记时有omitempty标记，序列化时将忽略有omitempty并且没有赋值的字段，当具有值时则显示。
	Sex    string  `json:"sex,omitempty" gorm:"column:sex"`
	Tall   float32 `json:"tall,omitempty" gorm:"column:tall"`
	Weight float32 `json:"weight,omitempty" gorm:"column:weight"`
	Age    int64   `json:"age,omitempty" gorm:"column:weight"`
}

func (*PersonalInformation) TableName() string { // 手动维护
	return "personal_information"
}
