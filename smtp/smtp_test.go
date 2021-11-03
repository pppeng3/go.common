package smtp

import (
	"testing"
	"time"
)

func TestWriteMsg(t *testing.T) {
	s := NewSMTPWriter("549822881@qq.com", "notpvncntywpbdbj", "smtp.qq.com:587", "LOGIN")
	a := &MailAcceptorGroup{
		FromAddress:      "549822881@qq.com",
		RecipientAddress: []string{"549822881@q.com"},
	}

	s.WriteMail(time.Now(), "text", "主题", []byte{}, a)
}
