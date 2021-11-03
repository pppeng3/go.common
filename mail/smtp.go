package mail

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/smtp"
	"strings"
	"time"
)

const (
	MailAuthPlain  = "PLAIN"
	MailAuthLogin  = "LOGIN"
	MailTypeText   = "text"
	MailTypeHtml   = "html"
	MailTypeXml    = "xml"
	MailTypeGif    = "gif"
	MailTypeJpeg   = "jpeg"
	MailTypePng    = "png"
	MailTypeJson   = "json"
	MailTypePdf    = "pdf"
	MailTypeWord   = "word"
	MailTypeStream = "stream"
)

func NewSMTPWriter(username, password, host, authtype, from string, to []string) *SMTPWriter {
	return &SMTPWriter{
		Username:         username,
		Password:         password,
		Host:             host,
		AuthType:         authtype,
		FromAddress:      from,
		RecipientAddress: to,
	}
}

func (s *SMTPWriter) getLoginAuth() smtp.Auth {
	return &loginAuth{s.Username, s.Password}
}

func (s *SMTPWriter) getPlainAuth() smtp.Auth {
	host := strings.Split(s.Host, ":")
	if len(strings.Trim(s.Username, " ")) == 0 && len(strings.Trim(s.Password, " ")) == 0 {
		return nil
	}
	return smtp.PlainAuth(
		"",
		s.Username,
		s.Password,
		host[0],
	)
}

func (s *SMTPWriter) sendMail(auth smtp.Auth, msgContent []byte) error {
	client, err := smtp.Dial(s.Host)
	if err != nil {
		return err
	}

	host, _, _ := net.SplitHostPort(s.Host)
	tlsConn := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	if err = client.StartTLS(tlsConn); err != nil {
		return err
	}

	if auth != nil {
		if err = client.Auth(auth); err != nil {
			return err
		}
	}

	if err = client.Mail(s.FromAddress); err != nil {
		return err
	}

	for _, rec := range s.RecipientAddress {
		if err = client.Rcpt(rec); err != nil {
			return err
		}
	}

	w, err := client.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msgContent)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return client.Quit()
}

// 如果为附件,msg传路径文件
func (s *SMTPWriter) createMessage(subject string, msg []byte, mailType string) ([]byte, error) {
	header := make(map[string]string)
	header["Subject"] = subject
	header["From"] = s.FromAddress
	header["To"] = strings.Join(s.RecipientAddress, ";")
	var message bytes.Buffer
	for k, v := range header {
		message.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	switch mailType {
	case MailTypeHtml:
		message.WriteString("Content-Type: text/html; charset=UTF-8\r\n\r\n")
		message.Write(msg)
	case MailTypeText:
		message.WriteString("Content-Type: text/plain\r\n\r\n")
		message.WriteString(time.Now().Format("2006-01-02 15:04:05>>>"))
		message.Write(msg)
	case MailTypeXml:
		message.WriteString("Content-Type: text/xml\r\n\r\n")
		message.Write(msg)
	case MailTypeGif:
		message.WriteString("Content-Type: image/gif\r\n\r\n")
	case MailTypeJpeg:
		message.WriteString("Content-Type: image/jpeg\r\n\r\n")
	case MailTypePng:
		message.WriteString("Content-Type: image/png\r\n\r\n")
	case MailTypeJson:
		message.WriteString("Content-Type: application/json\r\n\r\n")
	case MailTypePdf:
		message.WriteString("Content-Type: application/pdf\r\n\r\n")
	case MailTypeWord:
		message.WriteString("Content-Type: application/msword\r\n\r\n")
	case MailTypeStream:
		file := string(msg)
		message.WriteString("Content-Type: application/octet-stream\r\n\r\n")
		message.WriteString("Content-Transfer-Encoding: base64\r\n")
		message.WriteString("Content-Disposition: attachment; filename=\"" + file + "\"")
		message.WriteString("\r\n\r\n")
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		b := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
		base64.StdEncoding.Encode(b, data)
		message.Write(b)
		message.WriteString("\r\n")
	default:
		return nil, fmt.Errorf("invalid mailType%v", mailType)
	}
	return message.Bytes(), nil
}

func (s *SMTPWriter) WriteMail(subject string, msg []byte, mailType string) error {
	message, err := s.createMessage(subject, msg, mailType)
	if err != nil {
		return errors.New("Create message failed")
	}
	if len(message) <= 0 {
		return errors.New("empty message")
	}
	var auth smtp.Auth
	switch s.AuthType {
	case MailAuthPlain:
		auth = s.getPlainAuth()
	case MailAuthLogin:
		auth = s.getLoginAuth()
	}
	return s.sendMail(auth, message)
}
