package verification_code

import (
	"Rope_Net/internal"
	"Rope_Net/pkg/logger"
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"
)

func SendVerificationCode(email string, code string) error {
	qqConfig, err := internal.ReadQQEmailConfig()
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	//邮箱配置信息
	from := qqConfig.QQEmail
	password := qqConfig.QQEmailAuthCode
	smtpServer := qqConfig.SMTPServer
	smtpPort := qqConfig.SMTPPort

	to := []string{email}
	subject := "登录验证码"
	body := fmt.Sprintf("您的登录验证码是：%s", code)
	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", from, strings.Join(to, ","), subject, body)
	auth := smtp.PlainAuth("", from, password, smtpServer)
	//配置TLS
	logger.Info("配置TLS")
	tlsConfig := &tls.Config{
		ServerName: smtpServer,
	}
	conn, err := tls.Dial("tcp", smtpServer+":"+smtpPort, tlsConfig)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	defer conn.Close()
	c, err := smtp.NewClient(conn, smtpServer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	defer c.Quit()
	//进行验证
	if err = c.Auth(auth); err != nil {
		logger.Error(err.Error())
		return err
	}

	//设置收件人和发件人
	logger.Info("设置邮件")
	if err = c.Mail(from); err != nil {
		logger.Error(err.Error())
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			logger.Error(err.Error())
			return err
		}
	}
	//发送邮件内容
	w, err := c.Data()
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	_, err = w.Write([]byte(message))
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = w.Close()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil

}
