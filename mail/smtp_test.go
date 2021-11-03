package mail

import "testing"

func TestGetLoginAuth(t *testing.T) {
	mail := []string{"549822881@qq.com"}
	st := NewSMTPWriter("549822881@qq.com", "notpvncntywpbdbj", "smtp.qq.com:587", "LOGIN", "549822881@qq.com", mail)
	auth := st.getLoginAuth()
	err := st.sendMail(auth, []byte("111"))
	if err != nil {
		t.Fatalf("Auth:%v", err)
	}
}

func TestWriteMail(t *testing.T) {
	mail := []string{"549822881@qq.com"}
	st := NewSMTPWriter("549822881@qq.com", "notpvncntywpbdbj", "smtp.qq.com:587", "LOGIN", "549822881@qq.com", mail)
	err := st.WriteMail("测试", []byte("test"), "text")
	if err != nil {
		t.Fatalf("WriteMail:%v", err)
	}
}

func TestCreateMessage(t *testing.T) {
	mail := []string{"549822881@qq.com"}
	st := NewSMTPWriter("549822881@qq.com", "notpvncntywpbdbj", "smtp.qq.com:587", "LOGIN", "549822881@qq.com", mail)
	_, err := st.createMessage("test", []byte("1"), "text")
	if err != nil {
		t.Fatalf("Err:%v", err)
	}
}
