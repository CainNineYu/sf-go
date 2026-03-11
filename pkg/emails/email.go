package emails

import (
	"crypto/tls"
	gomail "gopkg.in/mail.v2"
)

func Send(subject string, text string, email string) (err error) {
	m := gomail.NewMessage()                   // 声明一封邮件对象
	m.SetHeader("From", "bitbotx@outlook.com") // 发件人
	m.SetHeader("To", email)                   // 收件人
	m.SetHeader("Subject", subject)            // 邮件主题
	m.SetBody("text/plain", text)              // 邮件内容

	// host 是提供邮件的服务器，port是服务器端口，username 是发送邮件的账号, password是发送邮件的密码
	d := gomail.NewDialer("smtp.office365.com", 587, "bitbotx@outlook.com", "Yu123456")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true} // 配置tls，跳过验证

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
