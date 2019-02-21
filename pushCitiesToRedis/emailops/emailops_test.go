package emailops

import "testing"

func TestSendMail(t *testing.T) {
	SendMail("test mail body")
}
