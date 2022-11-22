package utils

import (
	"fmt"
	"mongogram/config"
	"net/smtp"
)

func SendMail(to []string, from string, message []byte) error {
	host := "smtp.gmail.com"
	port := "587"

	addr := fmt.Sprintf("%v:%v", host, port)
	username := config.Config("SMTP_USERNAME")
	pass := config.Config("SMTP_PASS")

	auth := smtp.PlainAuth("", username, pass, host)
	return smtp.SendMail(addr, auth, from, to, message)

}
