package services

import (
	"context"
	"github.com/mailgun/mailgun-go/v4"
	"os"
)

func SendEmail(ctx context.Context, email string) (string, error) {
	mg := mailgun.NewMailgun(os.Getenv("domain"), os.Getenv("apiKey"))

	m := mg.NewMessage(
		"Choreo Task Manager Demo <postmaster@sandboxd7b034e52844480e83f0356a777d1536.mailgun.org>",
		"Task Manager Notification",
		"",
		email,
	)
	m.SetTemplate("task-repopened")
	_, id, err := mg.Send(ctx, m)

	return id, err
}
