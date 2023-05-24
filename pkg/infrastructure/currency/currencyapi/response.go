package currencyapi

import "time"

// this response from currencyapi website.
type apiResponse struct {
	Meta meta `json:"meta"`
	Data data `json:"data"`
}

type data struct {
	Uah currencyInfo `json:"UAH"`
	Usd currencyInfo `json:"usd"`
}
type currencyInfo struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}

type meta struct {
	LastUpdatedAt time.Time `json:"last_updated_at"`
}
