package v3

import "time"

type (
	Meta struct {
		LastUpdatedAt time.Time `json:"last_updated_at"`
	}
	CurrencyExchange struct {
		Code  string  `json:"code"`
		Value float64 `json:"value"`
	}
	CurrencyInfo struct {
		Symbol        string `json:"symbol"`
		Name          string `json:"name"`
		SymbolNative  string `json:"symbol_native"`
		DecimalDigits int    `json:"decimal_digits"`
		Rounding      int    `json:"rounding"`
		Code          string `json:"code"`
		NamePlural    string `json:"name_plural"`
	}

	StatusRequest  struct{}
	StatusResponse struct {
		Quotas struct {
			Month struct {
				Total     int `json:"total"`
				Used      int `json:"used"`
				Remaining int `json:"remaining"`
			} `json:"month"`
		} `json:"quotas"`
	}

	LatestRequest struct {
		From string   `json:"base_currency"`
		To   []string `json:"currencies"`
	}
	LatestResponse struct {
		Meta Meta                        `json:"meta"`
		Data map[string]CurrencyExchange `json:"data"`
	}
)
