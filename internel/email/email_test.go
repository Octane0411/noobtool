package email

import "testing"

func TestSendEmail(t *testing.T) {
	SendEmail("", "", "", "<h1>hi<h1>")
}
