package currencyapi

import (
	"context"
	"encoding/json"
	"exchange/pkg/domain"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	apikey       = "apikey"
	baseCurrency = "base_currency"
	currencies   = "currencies"
)

type CurrencyAPI struct {
	baseURL string
	cfg     *Config
}

// This is the implementation of logic that can get currency.
// So service doesn't need to know about how we do this, and we can implement any currency api and interfaces we want
// I'm not sure about putting this into infrastructure folder.
func NewCurrencyAPI(cfg *Config, link string) *CurrencyAPI {
	return &CurrencyAPI{
		cfg:     cfg,
		baseURL: link,
	}
}

func (api *CurrencyAPI) GetCurrency(ctx context.Context, cur *domain.Currency) (float64, error) {
	resp, err := api.makeLatestCurrencyRequest(ctx, cur.BaseCurrency, cur.QuoteCurrency)
	if err != nil {
		return 0, err
	}

	return resp.Data.Uah.Value, nil
}

func (api *CurrencyAPI) makeLatestCurrencyRequest(
	ctx context.Context,
	base, quote string,
) (*apiResponse, error) {
	const latest = "latest"

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/%s", api.baseURL, latest),
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()
	q.Add(apikey, api.cfg.apiKey)
	q.Add(baseCurrency, strings.ToUpper(base))
	q.Add(currencies, strings.ToUpper(quote))
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, domain.ErrInvalidStatus
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	apiResp := new(apiResponse)
	if err = json.Unmarshal(respBody, apiResp); err != nil {
		return nil, err
	}

	return apiResp, nil
}
