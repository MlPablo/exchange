package pkg

import "exchange/pkg/domain"

type Services struct {
	CurrencyService      domain.ICurrencyService
	EmailUserService     domain.IEmailService
	NotificatioinService domain.INotificationService
}

func NewServices(
	currency domain.ICurrencyService,
	email domain.IEmailService,
	notification domain.INotificationService,
) *Services {
	return &Services{
		CurrencyService:      currency,
		EmailUserService:     email,
		NotificatioinService: notification,
	}
}
