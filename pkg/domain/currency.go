package domain

import "context"

type Currency struct {
	BaseCurrency  string `json:"base_currency"`
	QuoteCurrency string `json:"quote_currency"`
}

func GetBitcoinToUAH() *Currency {
	return &Currency{
		BaseCurrency:  "BTC",
		QuoteCurrency: "UAH",
	}
}

type ICurrencyService interface {
	GetCurrency(ctx context.Context, c *Currency) (float64, error)
}
