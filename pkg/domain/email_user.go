package domain

import "context"

type EmailUser struct {
	Email string `json:"email"`
}

func NewEmailUser(email string) *EmailUser {
	return &EmailUser{
		Email: email,
	}
}

type IEmailService interface {
	NewEmailUser(ctx context.Context, eu *EmailUser) error
}

type EmailRepository interface {
	SaveEmail(ctx context.Context, eu *EmailUser) error
	GetByEmail(ctx context.Context, email string) (*EmailUser, error)
	GetAllEmails(ctx context.Context) ([]string, error)
}
