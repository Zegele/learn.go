package apis

//项目-pkg-apiss-types.go 里边装各种定义的标准对象

//type PersonalInfomation struct { //大写 不加json注解
//	Name   string //大写
//	Sex    string
//	Tall   float64
//	Weight float64
//	Age    int
//}
//type PersonalInfomation struct { //大写 加json注解
//
//	Name string `json:"name"` // 加了注解，marshal后 是使用注解的字段 这样可以保证精确性。
//	//name string `json:"name"` // 这里的Name 之类的必须大写 如果是 name （小写） 该字段转json时，就会丢失
//	// 注意，私有成员变量在序列化，反序列化时会被忽略
//	Sex    string  `json:"sex"`
//	Tall   float64 `json:"tall"`
//	Weight float64 `json:"weight"`
//	Age    int     `json:"age"`
//}
