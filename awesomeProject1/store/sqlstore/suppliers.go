package sqlstore

import (
	"awesomeProject1/apistructs"
	"awesomeProject1/services"
	"awesomeProject1/types"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	"sync"
	"time"
)

type DB struct {
	SqlDB        *sql.DB
	PriceService *services.PriceService
}

var once sync.Once

type APIQueryResponse struct {
	ID int `json:"id"`

	Time struct {
		Updated    string `json:"updated`
		UpdatedISO string `json:"updatedISO`
		UpdatedUK  string `json:"updateduk`
	} `json:"time"`
	Fetch_time  time.Time `json:"fetched_time"`
	Disclaimer  string    `json:"disclaimer"`
	ChartName   string    `json:"chartName"`
	Code        string    `json:"code"`
	Symbol      string    `json:"symbol"`
	Rate        string    `json:"rate"`
	Description string    `json:"description"`
	Rate_float  float64   `json:"rate_float"`
}

func CreateTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS api_data (
			id SERIAL PRIMARY KEY,
			updated timestamp,
			updated_iso timestamp,
			updated_uk timestamp,
			fetched_time timestamp,
			disclaimer text,
			chart_name text,
			currency_code text,
			currency_symbol text,
			currency_rate text,
			currency_description text,
			currency_rate_float numeric
		)
	`
	_, err := db.Exec(query)
	return err
}
func InsertData(db *sql.DB, data apistructs.CoinDeskStruct) error {
	for currencyCode, currencyDetail := range data.Bpi {
		query := `
			INSERT INTO api_data (
				updated, updated_iso, updated_uk,
				fetched_time,
				disclaimer, chart_name,
				currency_code, currency_symbol, currency_rate, currency_description, currency_rate_float
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,$11)
		`
		_, err := db.Exec(
			query,
			data.Time.Updated, data.Time.UpdatedISO, data.Time.UpdatedUK,
			time.Now(),
			data.Disclaimer, data.ChartName,
			currencyCode, currencyDetail.Symbol, currencyDetail.Rate, currencyDetail.Description, currencyDetail.RateFloat,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dbstruct *DB) DBData(c *gin.Context) {
	result, err := dbstruct.GetData()
	if err != nil {
		fmt.Println("Error in getting data :", err)
	}
	expired, data := dbstruct.CheckForExpiry(result)
	if expired || err != nil {
		data, err = dbstruct.FetchAPI()
		if err != nil {
			fmt.Println("Error in fetchingAPI:", err)
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"Data:": data,
	})
}

func (dbstruct *DB) GetData() ([]APIQueryResponse, error) {
	data, err := dbstruct.SqlDB.Query(`Select * from api_data`)
	if err != nil {
		fmt.Println("Error", err)
		return []APIQueryResponse{}, err
	}
	var result []APIQueryResponse
	for data.Next() {
		var temp APIQueryResponse
		err = data.Scan(
			&temp.ID,
			&temp.Time.Updated,
			&temp.Time.UpdatedISO,
			&temp.Time.UpdatedUK,
			&temp.Fetch_time,
			&temp.Disclaimer,
			&temp.ChartName,
			&temp.Code,
			&temp.Symbol,
			&temp.Rate,
			&temp.Description,
			&temp.Rate_float,
		)
		if err != nil {
			fmt.Println("Error in scanning", err)
			return []APIQueryResponse{}, err
		}

		result = append(result, temp)

	}
	return result, nil
}

func (dbstruct *DB) CheckForExpiry(data []APIQueryResponse) (bool, types.CurrentPricesResponse) {
	var result types.CurrentPricesResponse
	for _, val := range data {
		duration := time.Now().Sub(val.Fetch_time)
		duration = duration + time.Hour*5 + time.Minute*30

		if duration.Minutes() > 10 {
			return true, types.CurrentPricesResponse{}
		}
		result.Rates = append(result.Rates, types.CoinDeskCurrentPriceResponse{Price: val.Rate_float, Currency: val.Code})
	}

	return false, result

}
func (dbstruct *DB) FetchAPI() (types.CurrentPricesResponse, error) {
	resp, err := dbstruct.PriceService.CoinDeskPriceService()
	if err != nil {
		return types.CurrentPricesResponse{}, err
	}
	once.Do(func() { err = CreateTable(dbstruct.SqlDB) })
	if err != nil {
		fmt.Println("Error in Creating Table")
		return types.CurrentPricesResponse{}, err
	}
	resp.Fetch_time = time.Now()
	err = InsertData(dbstruct.SqlDB, resp)
	if err != nil {
		fmt.Println("Error in Inserting Data", err)
		return types.CurrentPricesResponse{}, err
	}
	var result types.CurrentPricesResponse
	for key, val := range resp.Bpi {
		result.Rates = append(result.Rates, types.CoinDeskCurrentPriceResponse{Price: val.RateFloat, Currency: key})
	}
	return result, nil
}
