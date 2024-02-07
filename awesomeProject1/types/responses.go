package types

type CoinDeskCurrentPriceResponse struct {
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}

type CurrentPricesResponse struct {
	Rates []CoinDeskCurrentPriceResponse `json:"rates"`
}
