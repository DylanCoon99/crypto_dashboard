package api


import (
	"os"
	"net/http"
	//"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
	"github.com/joho/godotenv"
)


// MAP COIN NAMES TO COIN IDs
var coinMap map[string]string


type InsightResponse struct {
	CoinName string      `json:"coin_name"`
	Time     time.Time   `json:"time"`
	Insight  string      `json:"insight"`
}


type Sentiment struct {
	Neg            float32 `json:"neg"`
	Neu            float32 `json:"neu"`
	Pos            float32 `json:"pos"`
	Compound       float32 `json:"compound"`
	SentimentLabel string      `json:"sentiment_label"`
	Time           time.Time   `json:"time"`
}


type SentimentResponse struct {
	CoinName        string      `json:"coin_name"`
	SentimentScore  Sentiment   `json:"sentiment_score"`
}


type MarketChartResponse struct {
	Prices       [][]float32      `json:"prices"`
	MarketCaps   [][]float32   `json:"market_caps"`
	TotalVolumes [][]float32 `json:"total_volumes"`
}



func InsightServiceAPI(coin_name string) *InsightResponse {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	resp, err := http.Get(os.Getenv("AI_API_ENDPOINT") + "/insight/" + coin_name)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: %v", err)
	}

	//fmt.Printf("Response: %s\n", body)

	var res InsightResponse

	err = json.Unmarshal(body, &res)

	if err != nil {
		log.Printf("Error unmarshaling: %v", err)
		return nil
	}


	return &res

}


func SentimentServiceAPI(coin_name string) *SentimentResponse {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	resp, err := http.Get(os.Getenv("AI_API_ENDPOINT") + "/sentiment/" + coin_name)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: %v", err)
	}

	//fmt.Printf("Response: %s\n", body)

	var res SentimentResponse

	err = json.Unmarshal(body, &res)

	if err != nil {
		log.Printf("Error unmarshaling: %v", err)
		return nil
	}


	return &res

}

// *MarketChartResponse

func HistoricPriceAPI(coin_name string) *MarketChartResponse {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	resp, err := http.Get(os.Getenv("COIN_GECKO_API_ENDPOINT") + "/coins/" + coin_name + "/market_chart?days=7&vs_currency=usd&interval=daily")
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: %v", err)
	}

	//fmt.Printf("Response: %s\n", body)

	
	var res MarketChartResponse

	err = json.Unmarshal(body, &res)

	if err != nil {
		log.Printf("Error unmarshaling: %v", err)
		return nil
	}
	


	return &res


}





