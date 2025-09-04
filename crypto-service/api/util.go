package api

import "time"

type InsightResponse struct {
	CoinName string    `json:"coin_name"`
	Time     time.Time `json:"time"`
	Insight  string    `json:"insight"`
}

type Sentiment struct {
	Neg            float32   `json:"neg"`
	Neu            float32   `json:"neu"`
	Pos            float32   `json:"pos"`
	Compound       float32   `json:"compound"`
	SentimentLabel string    `json:"sentiment_label"`
	Time           time.Time `json:"time"`
}

type SentimentResponse struct {
	CoinName       string    `json:"coin_name"`
	SentimentScore Sentiment `json:"sentiment_score"`
}

type MarketChartResponse struct {
	Prices       [][]float32 `json:"prices"`
	MarketCaps   [][]float32 `json:"market_caps"`
	TotalVolumes [][]float32 `json:"total_volumes"`
}

type Price struct {
	UnixTimeStamp float32 `json:"unix_timestamp"`
	Price         float32 `json:"price"`
}

type MarketCap struct {
	UnixTimeStamp float32 `json:"unix_timestamp"`
	MarketCap     float32 `json:"market_cap"`
}

type TotalVolume struct {
	UnixTimeStamp float32 `json:"unix_timestamp"`
	TotalVolume   float32 `json:"total_volume"`
}

type MarketChartFormatted struct {
	Prices       []Price       `json:"prices"`
	MarketCaps   []MarketCap   `json:"market_caps"`
	TotalVolumes []TotalVolume `json:"total_volumes"`
}

type Bitcoin struct {
	Coin struct {
		USD float64 `json:"usd"`
	} `json:"bitcoin"`
}

type Ethereum struct {
	Coin struct {
		USD float64 `json:"usd"`
	} `json:"ethereum"`
}

type Solana struct {
	Coin struct {
		USD float64 `json:"usd"`
	} `json:"solana"`
}

type RealTimePrice struct {
	CoinName string    `json:"coin_name"`
	PriceUSD float64   `json:"price_usd"`
	Date     time.Time `json:"date"`
}
