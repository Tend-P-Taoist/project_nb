package sender

import (
	"net/smtp"
)

const (
	SMTPHOST = "smtp.163.com"
	USERNAME = "pansifan0525@163.com"
	PASSWORD = "psmortal0525"
	IDENTITY = ""

)

func SendEmail(addr,msg string) error{
	auth := smtp.PlainAuth(IDENTITY,USERNAME,PASSWORD,SMTPHOST)
	var m string = "To: " + addr + "\r\nFrom: " + "pansifan0525@163.com" + "\r\nSubject: " + "账号激活" + "\r\n" + "Content-Type: text/html; charset=UTF-8" + "\r\n\r\n" + msg
	err := smtp.SendMail(SMTPHOST + ":25",auth,"pansifan0525@163.com",[]string{addr},[]byte(m))

	return err
}


func SendMessage(phone,msg string ) error{
	return  nil
}