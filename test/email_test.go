package test

import (
	"net/smtp"
	"testing"

	"oj/define"

	"github.com/jordan-wright/email"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Wxf <68725032@qq.com>"
	e.To = []string{"68725032@qq.com"}
	e.Subject = "验证码发送"
	e.HTML = []byte("您的验证码：<b>123456</b>")
	// 返回 EOF 时，关闭SSL重试
	err := e.Send("smtp.qq.com:587",
		smtp.PlainAuth("", "68725032@qq.com", define.MailPassword, "smtp.qq.com"),
	)
	if err != nil {
		t.Fatal(err)
	}
}
