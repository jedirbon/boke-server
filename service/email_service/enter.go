package email_service

import (
	"boke-server/global"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"strings"
)

// 注册账号
func SendRegisterCode(to string, code string) error {
	text := fmt.Sprintf("你正在进行账号注册操作，这是你的验证码%s，10分中内有效", code)
	em := global.Config.Email
	subject := fmt.Sprintf("【%s】注册邮箱", em.SendNickname)
	return SendEmail(to, subject, text)
}

// 充值密码
func SendResetPwdCode(to string, code string) error {
	em := global.Config.Email
	text := fmt.Sprintf("你正在进行账号密码重置操作，这是你的验证码%s，10分中内有效", code)
	subject := fmt.Sprintf("【%s】重置密码", em.SendNickname)
	return SendEmail(to, subject, text)
}

func SendEmail(to, subject string, text string) (err error) {
	em := global.Config.Email
	// 创建邮件
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", em.SendNickname, em.SendEmail)
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(text)
	// SMTP 认证
	auth := smtp.PlainAuth("", em.SendEmail, em.AuthCode, em.Domain)
	// 发送邮件
	err = e.Send(fmt.Sprintf("%s:%d", em.Domain, em.Port), auth)
	if err != nil && !strings.Contains(err.Error(), "short response:") {
		fmt.Println(err)
		fmt.Println("发送失败")
		return err
	} else {
		fmt.Println("发送成功")
		return nil
	}
	return nil
}
