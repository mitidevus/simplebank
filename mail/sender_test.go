package mail

import (
	"testing"

	"github.com/mitidevus/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	// Skip this test when run "make test"
	// If it's true, then it means -test.short flag has been set
	// Note: Remember to add -short flag into test in Makefile
	if testing.Short() {
		t.Skip()
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	<p>This is a text message from <a href="https://portfolio-mitidevus.netlify.app/">Simple Bank</a></p>
	`
	to := []string{"mitidevus@gmail.com", "minhtri.do2410@gmail.com"}

	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
