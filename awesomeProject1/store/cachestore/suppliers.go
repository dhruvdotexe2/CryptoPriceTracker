package cachestore

import (
	"awesomeProject1/services"
	"awesomeProject1/types"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"net/http"
	"time"
)

type Cache struct {
	PriceService *services.PriceService
}

var myCache = cache.New(5*time.Minute, 10*time.Minute)

func (cache *Cache) CacheData(c *gin.Context) {
	resp, err := cache.GetData("prices")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Error in getting response",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Data": resp,
	})
}
func (cachestruct *Cache) GetData(key string) (types.CurrentPricesResponse, error) {
	cachedData, found := myCache.Get(key)
	if found {
		fmt.Println("Data found in cache:", cachedData)
		return types.CurrentPricesResponse{}, nil
	}

	// If not found in cache, fetch from the source
	fetchedData, err := cachestruct.FetchAPI()
	if err != nil {
		return types.CurrentPricesResponse{}, err
	}

	// Store fetched data in the cache with a 10-minute expiration
	myCache.Set(key, fetchedData, 10*time.Minute)

	fmt.Println("Fetched data:", fetchedData)
	return fetchedData, nil
}
func (cache *Cache) FetchAPI() (types.CurrentPricesResponse, error) {
	resp, err := cache.PriceService.CoinDeskPriceService()
	if err != nil {
		return types.CurrentPricesResponse{}, err
	}

	resp.Fetch_time = time.Now()
	var result types.CurrentPricesResponse
	for key, val := range resp.Bpi {
		result.Rates = append(result.Rates, types.CoinDeskCurrentPriceResponse{Price: val.RateFloat, Currency: key})
	}
	return result, nil
}
