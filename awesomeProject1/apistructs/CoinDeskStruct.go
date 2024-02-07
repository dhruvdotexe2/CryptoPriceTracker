package apistructs

import (
	"time"
)

type CoinDeskStruct struct {
	Time struct {
		Updated    string `json:"updated`
		UpdatedISO string `json:"updatedISO`
		UpdatedUK  string `json:"updateduk`
	} `json:"time"`
	Disclaimer string                 `json:"disclaimer"`
	ChartName  string                 `json:"chartName"`
	Bpi        map[string]CurrencyBPI `json:"bpi"`
	Fetch_time time.Time              `json:"fetched_time"`
}

type CurrencyBPI struct {
	Code        string  `json:"code"`
	Symbol      string  `json:"symbol"`
	Rate        string  `json:"rate"`
	Description string  `json:"description"`
	RateFloat   float64 `json:"rate_float"`
}
