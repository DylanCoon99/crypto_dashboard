package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	//"github.com/joho/godotenv"
)

/*
func UnixTimeToDate(unixTime int32) time.Time {

	return time.Unix(int64(unixTime), 0)
}
*/

func InsightServiceAPI(coin_name string) *InsightResponse {
	/*
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	*/

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
	/*
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	*/

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

func HistoricPriceAPI(coin_name string) *MarketChartFormatted {

	/*
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	*/

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

	marketChart := MarketChartFormatted{
		Prices:       make([]Price, 0),
		MarketCaps:   make([]MarketCap, 0),
		TotalVolumes: make([]TotalVolume, 0),
	}

	n := len(res.Prices)

	for i := 0; i < n; i++ {

		unix_time_stamp := res.Prices[i][0]
		price := res.Prices[i][1]

		marketChart.Prices = append(marketChart.Prices, Price{UnixTimeStamp: unix_time_stamp, Price: price})

		market_cap := res.MarketCaps[i][1]
		marketChart.MarketCaps = append(marketChart.MarketCaps, MarketCap{UnixTimeStamp: unix_time_stamp, MarketCap: market_cap})

		total_volume := res.TotalVolumes[i][1]
		marketChart.TotalVolumes = append(marketChart.TotalVolumes, TotalVolume{UnixTimeStamp: unix_time_stamp, TotalVolume: total_volume})

	}

	return &marketChart

}

func RealTimePriceAPI(coin_name string) *RealTimePrice {
	// /simple/price?vs_currencies=usd&ids=bitcoin

	/*
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	*/

	resp, err := http.Get(os.Getenv("COIN_GECKO_API_ENDPOINT") + "/simple/price?vs_currencies=usd&ids=" + coin_name)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: %v", err)
	}

	//fmt.Printf("Response: %s\n", body)

	switch coin_name {
	case "bitcoin":
		var coin Bitcoin

		err = json.Unmarshal(body, &coin)

		if err != nil {
			log.Printf("Error unmarshaling: %v", err)
			return nil
		}

		res := RealTimePrice{
			CoinName: coin_name,
			PriceUSD: coin.Coin.USD,
			Date:     time.Now(),
		}

		return &res

	case "ethereum":
		var coin Ethereum

		err = json.Unmarshal(body, &coin)

		if err != nil {
			log.Printf("Error unmarshaling: %v", err)
			return nil
		}

		res := RealTimePrice{
			CoinName: coin_name,
			PriceUSD: coin.Coin.USD,
			Date:     time.Now(),
		}

		return &res
	case "solana":
		var coin Solana

		err = json.Unmarshal(body, &coin)

		if err != nil {
			log.Printf("Error unmarshaling: %v", err)
			return nil
		}

		res := RealTimePrice{
			CoinName: coin_name,
			PriceUSD: coin.Coin.USD,
			Date:     time.Now(),
		}

		return &res

	}

	return nil

}
