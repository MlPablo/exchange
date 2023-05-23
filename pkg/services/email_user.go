package services

import (
	"context"
	"exchange/pkg/domain"
)

type emailUserService struct {
	ctx           context.Context
	emailUserRepo domain.EmailRepository
}

func NewEmailUserService(
	ctx context.Context,
	emailRepo domain.EmailRepository,
) domain.IEmailService {
	return &emailUserService{
		ctx:           ctx,
		emailUserRepo: emailRepo,
	}
}

func (e *emailUserService) NewEmailUser(ctx context.Context, emailUser *domain.EmailUser) error {
	return e.emailUserRepo.SaveEmail(ctx, emailUser)
}
