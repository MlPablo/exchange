package mail

import (
	"context"
	"encoding/json"
	"fmt"
	"net/smtp"
)

type EmailSender struct {
	cfg *Config
}

func NewMailService(cfg *Config) *EmailSender {
	return &EmailSender{
		cfg: cfg,
	}
}

func (e *EmailSender) SendEmail(ctx context.Context, data any, recievers ...string) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return e.sendEmail(bytes, recievers...)
}

func (s *EmailSender) sendEmail(message []byte, receiversEmail ...string) error {
	// Authentication.
	// TODO: add auth into struct
	auth := smtp.PlainAuth("", s.cfg.user, s.cfg.password, s.cfg.smtpHost)

	// Sending email.
	return smtp.SendMail(
		fmt.Sprintf("%s:%s", s.cfg.smtpHost, s.cfg.smtpPort),
		auth,
		s.cfg.user,
		receiversEmail,
		message,
	)
}
