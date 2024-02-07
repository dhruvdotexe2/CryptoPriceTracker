package test

import (
	"awesomeProject1/apistructs"
	"awesomeProject1/client"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type MockCoinDeskClient struct {
	MockCoinDeskURL string
}

func (cdc *MockCoinDeskClient) GetCryptoPrice() (apistructs.CoinDeskStruct, error) {
	return apistructs.CoinDeskStruct{Time: struct {
		Updated    string `json:"updated`
		UpdatedISO string `json:"updatedISO`
		UpdatedUK  string `json:"updateduk`
	}(struct {
		Updated    string `json:"updated"`
		UpdatedISO string `json:"updatedISO"`
		UpdatedUK  string `json:"updateduk"`
	}{
		Updated:    "Mock updated",
		UpdatedISO: "Mock updatedISO",
		UpdatedUK:  "Mock updatedUK",
	}), Disclaimer: "Mock disclaimer", ChartName: "Mock chartName", Bpi: map[string]apistructs.CurrencyBPI{
		"USD": {
			Code:        "USD",
			Symbol:      "$",
			Rate:        "50000",
			Description: "Mock description",
			RateFloat:   50000.00,
		},
		"GBP": {
			Code:        "USD",
			Symbol:      "$",
			Rate:        "50000",
			Description: "Mock description",
			RateFloat:   50000.00,
		},
		"EUR": {
			Code:        "USD",
			Symbol:      "$",
			Rate:        "50000",
			Description: "Mock description",
			RateFloat:   50000.00,
		},
		// Add more mock data for other currencies if needed
	}, Fetch_time: time.Now()}, nil
}

func TestGetCryptoPrice(t *testing.T) {
	cdclient := &client.CoinDeskClient{CoinDeskURL: "https://api.coindesk.com/v1/bpi/currentprice.json"}
	mockserver := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Respond with a mock JSON payload
		mockResponse := `{"updated":"Mock updated","updatedISO":"Mock updated","updateduk":"Mock updated","disclaimer":"Mock disclaimer","chartName":"Mock chartName","bpi":{"USD":{"code": "USD","symbol":"$","rate":"50000","description":"Mock Description","rate_float": 50000.0},"EUR":{"code": "EUR","symbol":"%","rate":"60000","description":"Mock Description","rate_float": 60000.0},"GBP":{"code": "GBP","symbol":"#","rate":"70000","description":"Mock Description","rate_float": 70000.0}},"fetched_time":2024-02-04 20:55:31.16597}`
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	}))
	defer mockserver.Close()

	// Use the mock server URL for the CoinDeskClient
	cdclient.CoinDeskURL = mockserver.URL

	// Call the GetCryptoPrice method
	result, err := cdclient.GetCryptoPrice()

	// Check for errors
	assert.NoError(t, err, "Unexpected error during GetCryptoPrice")

	// Check the structure of the response
	expected := apistructs.CoinDeskStruct{Time: struct {
		Updated    string `json:"updated`
		UpdatedISO string `json:"updatedISO`
		UpdatedUK  string `json:"updateduk`
	}(struct {
		Updated    string `json:"updated"`
		UpdatedISO string `json:"updatedISO"`
		UpdatedUK  string `json:"updateduk"`
	}{
		Updated:    "Mock updated",
		UpdatedISO: "Mock updatedISO",
		UpdatedUK:  "Mock updatedUK",
	}), Disclaimer: "Mock disclaimer", ChartName: "Mock chartName", Bpi: map[string]apistructs.CurrencyBPI{
		"USD": {
			Code:        "USD",
			Symbol:      "$",
			Rate:        "50000",
			Description: "Mock description",
			RateFloat:   50000.00,
		},
		"GBP": {
			Code:        "USD",
			Symbol:      "$",
			Rate:        "50000",
			Description: "Mock description",
			RateFloat:   50000.00,
		},
		"EUR": {
			Code:        "USD",
			Symbol:      "$",
			Rate:        "50000",
			Description: "Mock description",
			RateFloat:   50000.00,
		},
		// Add more mock data for other currencies if needed
	}, Fetch_time: time.Now()}

	assert.ElementsMatch(t, expected, result, "Unexpected result from GetCryptoPrice")
}
