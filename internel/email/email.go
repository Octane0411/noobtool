package email

import (
	"fmt"
	"github.com/go-gomail/gomail"
	"log"
	"net/mail"
	"noobtool/global"
	"strings"
)


// 全局变量，因为发件人账号、密码，需要在发送时才指定
var serverHost, fromEmail, fromPasswd string
var serverPort int
var m *gomail.Message

// SendEmail 发送邮件
// Toers: 接受者邮件，如有多个，用逗号分割，不能为空
// Ccers: 抄送者邮件，如有多个，用逗号分割，可以为空
// subject: 邮件的主题
// body: 邮件的内容
func SendEmail(Toers, Ccers, subject, body string) error {
	toers := []string{}

	serverHost = global.EmailSetting.ServerHost
	serverPort = global.EmailSetting.ServerPort
	fromEmail = global.EmailSetting.FromEmail
	fromPasswd = global.EmailSetting.FromPasswd

	m = gomail.NewMessage()

	if len(Toers) == 0 {
		toers = append(toers, global.EmailSetting.FromEmail)
	} else {
		for _, tmp := range strings.Split(Toers, ",") {
			if !ValidEmail(tmp) {
				return fmt.Errorf("邮箱格式不正确：%v", tmp)
			}
			toers = append(toers, strings.TrimSpace(tmp))
		}
	}
	// 收件人可以有多个，故用此方式
	m.SetHeader("To", toers...)
	//抄送列表
	if len(Ccers) != 0 {
		for _, tmp := range strings.Split(Ccers, ",") {
			if !ValidEmail(tmp) {
				return fmt.Errorf("邮箱格式不正确：%v", tmp)
			}
			toers = append(toers, strings.TrimSpace(tmp))
		}
		m.SetHeader("Cc", toers...)
	}
	// 发件人
	// 第三个参数为发件人别名，如"李大锤"，可以为空（此时则为邮箱名称）
	m.SetAddressHeader("From", fromEmail, "")
	// 主题
	m.SetHeader("Subject", subject)
	// 正文
	m.SetBody("text/html", body)
	d := gomail.NewPlainDialer(serverHost, serverPort, fromEmail, fromPasswd)
	// 发送
	err := d.DialAndSend(m)
	if err != nil {
		log.Printf("sendEmail DialAndSend error: %v", err)
		return err
	}
	return nil
}

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}


