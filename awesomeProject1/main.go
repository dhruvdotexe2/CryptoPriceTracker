package main

import (
	"awesomeProject1/client"
	"awesomeProject1/controller"
	"awesomeProject1/services"
	"awesomeProject1/store"
	"awesomeProject1/store/cachestore"
	"awesomeProject1/store/sqlstore"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file", err)
		return
	}
	CoinDesk := client.CoinDeskClient{CoinDeskURL: viper.GetString("APILinks.coindeskurl")}
	PriceService := services.PriceService{CryptoPriceClient: &CoinDesk}

	router := gin.Default()
	db, err := store.GetDBInstance()
	if err != nil {
		fmt.Println("Error in connecting DB", err)
	}
	PriceServiceController := controller.PriceServiceController{Store: &controller.SingleStore{PriceService: &PriceService, DB: &sqlstore.DB{db, &PriceService}, Cache: &cachestore.Cache{PriceService: &PriceService}}}
	router.GET("/current-prices", PriceServiceController.FetchCurrentPrice)
	router.Run(":8081")
}
