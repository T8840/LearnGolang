// come from : https://www.ctolib.com/topics-143735.html
// QQ Email Setting To Use SMTP: https://www.cnblogs.com/kimsbo/p/10671851.html
package main

import (
	"log"
	"mime"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

func main() {
	sendEmail("测试第三方 email 库", "receiver@qq.com")
}

func sendEmail(subject string, tos ...string) error {
	e := email.NewEmail()

	smtpUsername := "sender@qq.com"
	e.From = mime.QEncoding.Encode("UTF-8", "Go语言中文网") + "<sender@qq.com>"
	e.To = tos
	e.Subject = subject
	e.HTML = []byte("<h1>HTML 正文</h1>")
	e.AttachFile("zap.log")

	auth := smtp.PlainAuth("", smtpUsername, "password", "smtp.qq.com")
	err := e.Send("smtp.qq.com:25", auth)
	if err != nil {
		log.Println("Send Mail to", strings.Join(tos, ","), "error:", err)
		return err
	}
	log.Println("Send Mail to", strings.Join(tos, ","), "Successfully")
	return nil
}
