package controller

import (
	"awesomeProject1/services"
	"awesomeProject1/store/cachestore"
	"awesomeProject1/store/sqlstore"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type PriceServiceController struct {
	Store *SingleStore
	// use single store and that can have both cache and db instances
}
type SingleStore struct {
	PriceService *services.PriceService
	DB           *sqlstore.DB
	Cache        *cachestore.Cache
}
type PriceServiceControllerInterface interface {
	FetchCurrentPrice(c *gin.Context)
}

type Services interface {
	GetDataFromDB() bool
	FetchCurrentPrice(c *gin.Context)
}

func (psc *PriceServiceController) FetchCurrentPrice(c *gin.Context) {
	if GetDataFromDB() {
		psc.Store.DB.DBData(c)
		return
	}
	psc.Store.Cache.CacheData(c)
}
func GetDataFromDB() bool {
	viper.SetConfigName("config")
	viper.AddConfigPath("../.")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		return false
	}
	return viper.GetBool("DBoverCache")
}
