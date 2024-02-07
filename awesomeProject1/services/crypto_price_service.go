package services

import (
	"awesomeProject1/apistructs"
	"awesomeProject1/client"
	"awesomeProject1/types"
)

type PriceService struct {
	CryptoPriceClient client.ClientInterface
}
type PriceServiceInterface interface {
	CoinDeskPriceService() (types.CurrentPricesResponse, error)
}

func (ps *PriceService) CoinDeskPriceService() (apistructs.CoinDeskStruct, error) {
	return ps.CryptoPriceClient.GetCryptoPrice()
}
