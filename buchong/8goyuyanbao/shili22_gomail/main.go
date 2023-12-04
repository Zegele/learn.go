package main

import _ "github.com/go-gomail/gomail"

type EmailParam struct {
	// ServerHost 邮箱服务器地址，如腾讯邮箱为smtp.qq.com
	ServerHost string
	// ServerPort 邮箱服务器端口，如腾讯邮箱为465
	ServerPort int
	// FromEmail 发件人邮箱地址
	FromEmail string
	// FromPasswd 发件人邮箱密码（注意，这里是明文形式）， TODO: 如何设置成密文？
	FromPasswd string
	// Toers 接收者邮件，如有多个，则以英文逗号“ , ” 隔开，不能为空
	Toers string
	// CCers 抄送者邮件，如有多个，则以英文逗号“  , ” 隔开，可以为空
	CCers string
}

// 全局变量， 因为发件人账号、密码，需要在发送时才指定
// 注意，由于是小写，外面的包无法使用

var serverHost, fromEmail, fromPasswd string
var serverPort int

var m *gomail.Message
