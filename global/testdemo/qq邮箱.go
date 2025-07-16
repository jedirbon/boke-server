package main

import (
	"boke-server/core"
	"boke-server/flags"
	"boke-server/global"
	"boke-server/service/email_service"
)

// 解析请求参数
type request struct {
	From    string   `json:"from" binding:"required"`    // 发件人
	To      []string `json:"to" binding:"required"`      // 收件人列表
	Subject string   `json:"subject" binding:"required"` // 邮件主题
	Text    string   `json:"text"`                       // 纯文本内容（可选）
	HTML    string   `json:"html"`                       // HTML 内容（可选）
	SMTP    struct {
		Host     string `json:"host" binding:"required"`     // SMTP 服务器地址
		Port     int    `json:"port" binding:"required"`     // SMTP 端口
		Username string `json:"username" binding:"required"` // SMTP 用户名
		Password string `json:"password" binding:"required"` // SMTP 密码
	} `json:"smtp" binding:"required"`
}

func main() {
	flags.Parse()
	//读取配置信息
	global.Config = core.ReadConf()
	core.InitLogrus()
	email_service.SendRegisterCode("1145480106@qq.com", "10020")
}
