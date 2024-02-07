package client

import (
	"awesomeProject1/apistructs"
	"encoding/json"
	"net/http"
)

type CoinDeskClient struct {
	CoinDeskURL string
}

type ClientInterface interface {
	GetCryptoPrice() (apistructs.CoinDeskStruct, error)
}

func (cdc *CoinDeskClient) GetCryptoPrice() (apistructs.CoinDeskStruct, error) {
	fetchUrl := cdc.CoinDeskURL
	res, err := http.Get(fetchUrl)
	if err != nil {
		return apistructs.CoinDeskStruct{}, err
	}
	defer res.Body.Close()
	var response apistructs.CoinDeskStruct
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return apistructs.CoinDeskStruct{}, err
	}

	return response, nil
}
