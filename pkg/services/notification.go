package services

import (
	"context"
	"exchange/pkg/domain"
)

type MailService interface {
	SendEmail(ctx context.Context, data any, recievers ...string) error
}

type notificationService struct {
	ctx             context.Context
	emailUserRepo   domain.EmailRepository
	currencyService domain.ICurrencyService
	mailService     MailService
}

func NewNotificationService(
	ctx context.Context,
	emailRepo domain.EmailRepository,
	currencyService domain.ICurrencyService,
	mailService MailService,
) domain.INotificationService {
	return &notificationService{
		ctx:             ctx,
		emailUserRepo:   emailRepo,
		mailService:     mailService,
		currencyService: currencyService,
	}
}

func (n *notificationService) Notify(ctx context.Context, not *domain.Notification) error {
	btcUsd := domain.GetBitcoinToUAH()

	currency, err := n.currencyService.GetPrice(ctx, btcUsd)
	if err != nil {
		return err
	}

	emails, err := n.emailUserRepo.GetAllEmails(ctx)
	if err != nil {
		return err
	}

	return n.mailService.SendEmail(ctx, currency, emails...)
}
