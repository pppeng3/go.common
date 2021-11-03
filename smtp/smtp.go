package smtp

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"strings"
	"time"
)

const (
	MailAuthPlain = "PLAIN"
	MailAuthLogin = "LOGIN"
	MailTypeText  = "text"
	MailTypeHtml  = "html"
)

func (auth *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (auth *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username":
			return []byte(auth.username), nil
		case "Password":
			return []byte(auth.password), nil
		default:
			return nil, errors.New("Unknown fromServer")
		}
	}
	return nil, nil
}

func NewSMTPWriter(username, password, host, authtype string) *SMTPWriter {
	return &SMTPWriter{
		Username: username,
		Password: password,
		Host:     host,
		AuthType: authtype,
	}
}

// Init smtp writer with json config.
// config like:
//	{
//		"username":"example@gmail.com",
//		"password:"password",
//		"host":"smtp.gmail.com:465",
//		"subject":"email title",
//		"fromAddress":"from@example.com",
//		"sendTos":["email1","email2"],
//	}
func (s *SMTPWriter) Init(jsonConfig string) error {
	return json.Unmarshal([]byte(jsonConfig), s)
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

func (s *SMTPWriter) sendMail(hostAddressWithPort string, auth smtp.Auth, fromAddress string, recipients []string, msgContent []byte) error {
	client, err := smtp.Dial(hostAddressWithPort)
	if err != nil {
		return err
	}

	host, _, _ := net.SplitHostPort(hostAddressWithPort)
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

	if err = client.Mail(fromAddress); err != nil {
		return err
	}

	for _, rec := range recipients {
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

func (s *SMTPWriter) createMessage(when time.Time, mailtype string, subject string, msg []byte, acceptor *MailAcceptorGroup) ([]byte, error) {
	header := make(map[string]string)
	header["Subject"] = subject
	header["Content-Type"] = "text/plain; charset=UTF-8"
	header["From"] = acceptor.FromAddress
	header["To"] = strings.Join(acceptor.RecipientAddress, ";")
	var message bytes.Buffer
	for k, v := range header {
		message.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	switch mailtype {
	case MailTypeHtml:
		message.WriteString("Content-Type: text/html,charset=UTF-8\r\n\r\n")
		message.Write(msg)
	case MailTypeText:
		message.WriteString("Content-Type: text/plain; charset=UTF-8\r\n\r\n")
		message.WriteString(when.Format("2006-01-02 15:04:05>>>"))
		message.Write(msg)
	default:
		return nil, fmt.Errorf("invalid mailtype%v", mailtype)
	}
	return message.Bytes(), nil
}

func (s *SMTPWriter) WriteMail(when time.Time, mailtype string, subject string, msg []byte, acceptor *MailAcceptorGroup) error {
	message, err := s.createMessage(when, mailtype, subject, msg, acceptor)
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
	return s.sendMail(s.Host, auth, acceptor.FromAddress, acceptor.RecipientAddress, message)
}

func (s *SMTPWriter) WriteMsg(when time.Time, subject string, msg []byte, acceptor *MailAcceptorGroup) error {
	return s.WriteMail(when, MailTypeText, subject, msg, acceptor)
}
