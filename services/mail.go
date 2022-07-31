package services

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/quotedprintable"
	"net/smtp"
	"strings"
)

// โปรดทำการยืนยันอีเมลของคุณภายใน 24 ชั่วโมง
// ในกรณีที่ไม่พบเมลที่กล่องขาเข้า (Inbox) ขอให้ตรวจสอบที่ Spam หรือ Junk mail
/**
For example, for Gmail:
Login into your account Gmail or Google Apps then goto:https://www.google.com/settings/security/lesssecureapps
and select Turn On for Access for less secure apps.

**/

const (
	/**
		Gmail SMTP Server
	**/
	SMTPServer = "smtp.gmail.com"
	callback   = "http://localhost:8080/auth/email/callback"
)

type Sender struct {
	User     string
	Password string
}

func (sv *service) SendMail(mailTo, url, passcode, code string) error {
	sender := newSender("pien.consultant@gmail.com", "djlcyigrudibetsn")

	receiver := []string{mailTo}

	subject := "รับรหัสเพื่อเข้าระบบเพียรนิวส์"

	message, err := mailTemplate(url, passcode, code)
	if err != nil {
		return err
	}

	bodyMessage := writeEmail(sender, receiver, "text/html", subject, message)
	err = sendGMail(sender, receiver, subject, bodyMessage)
	if err != nil {
		return err
	}

	return nil
}

func newSender(Username, Password string) *Sender {
	return &Sender{Username, Password}
}

func sendGMail(sender *Sender, To []string, Subject, bodyMessage string) error {
	msg := "From: " + sender.User + "\n" +
		"To: " + strings.Join(To, ",") + "\n" +
		"Subject: " + Subject + "\n" + bodyMessage

	err := smtp.SendMail(SMTPServer+":587",
		smtp.PlainAuth("", sender.User, sender.Password, SMTPServer),
		sender.User, To, []byte(msg))

	if err != nil {
		return err
	}
	return nil
}

func writeEmail(sender *Sender, dest []string, contentType, subject, bodyMessage string) string {
	header := make(map[string]string)
	header["From"] = sender.User

	receipient := ""

	for _, user := range dest {
		receipient = receipient + user
	}

	header["To"] = receipient
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(bodyMessage))
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message
}

func mailTemplate(url, passcode, code string) (string, error) {
	content, err := ioutil.ReadFile("./pages/mail.html")
	if err != nil {

		return "", err
	}
	html := string(content)
	html = strings.Replace(html, "{{URL}}", fmt.Sprintf("%v/auth/passcode/%v", url, passcode), -1)
	html = strings.Replace(html, "{{CODE}}", code, -1)

	return html, nil
}
