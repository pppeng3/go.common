package mail

import (
	"errors"
	"net/smtp"
)

type SMTPWriter struct {
	Username         string   `mapstructure:"username"`
	Password         string   `mapstructure:"password"`
	Host             string   `mapstructure:"host"`
	AuthType         string   `mapstructure:"auth"`
	FromAddress      string   `mapstructure:"fromAddress"`
	RecipientAddress []string `mapstructure:"sendTos"`
}

type loginAuth struct {
	username, password string
}

func (auth *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (auth *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(auth.username), nil
		case "Password:":
			return []byte(auth.password), nil
		default:
			return nil, errors.New("Unknown fromServer")
		}
	}
	return nil, nil
}
