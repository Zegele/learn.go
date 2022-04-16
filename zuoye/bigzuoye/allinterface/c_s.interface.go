package allinterface

type ClientServeInterface interface {
	//注册
	Register() error
	//登录
	Login(account int64, password int64)
	// 查看在线人
	Online() error
	// 发起聊天
	ChatWith(account1, account2 int64) error
	//查看聊天历史
	ChatHistory()
}
