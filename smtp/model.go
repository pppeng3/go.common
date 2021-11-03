package smtp

import ()

type SMTPWriter struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	AuthType string `mapstructure:"auth"`
}

type MailAcceptorGroup struct {
	FromAddress      string   `mapstructure:"fromAddress"`
	RecipientAddress []string `mapstructure:"sendTos"`
}

type loginAuth struct {
	username, password string
}
