package mail

import (
	"fmt"
	"github.com/Xurliman/auth-service/internal/config/config"
	"github.com/Xurliman/auth-service/pkg/log"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

func VerifyEmail(email string, verificationURL string) {
	cfg := config.GetMailSettings()

	m := gomail.NewMessage()
	m.SetHeader("From", cfg.From)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Welcome!")
	m.SetBody("text/html", fmt.Sprintf(`
		<p>Hello,%s</p>
		<p>Please verify your email by clicking the link below:</p>
		<a href="%s">Verify Email</a>
		<p>This link will expire in 24 hours.</p>
	`,
		email,
		verificationURL))

	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.User, cfg.Password)
	if err := d.DialAndSend(m); err != nil {
		log.Warn("error sending email: ", zap.Error(err))
	}
}
