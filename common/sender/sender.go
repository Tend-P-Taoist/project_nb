package sender

import (
	"net/smtp"
)

func SendEmail(addr string,msg interface{}) error{
	auth := smtp.PlainAuth("","","","")
	err := smtp.SendMail("",auth,"",[]string{""},[]byte{})
	return err
}
