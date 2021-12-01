package binance

import (
    "context"
    "encoding/json"
    "log"

    "go-hexagonal/internal/port.adapter/dependency/http"
    "go-hexagonal/util/logger"
)

/**
 * @author Rancho
 * @date 2021/12/1
 */

type TickerPrice struct {
	Symbol       string  `json:"symbol"`
	Price        float64 `json:"price,string"`
	TickerPrices []*TickerPrice
}

type Ticker24Hour struct {
	Symbol             string  `json:"symbol"`
	PriceChange        float64 `json:"priceChange,string"`
	PriceChangePercent float64 `json:"priceChangePercent,string"`
	WeightedAvgPrice   float64 `json:"weightedAvgPrice,string"`
	PrevClosePrice     float64 `json:"prevClosePrice,string"`
	LastPrice          float64 `json:"lastPrice,string"`
	LastQty            float64 `json:"lastQty,string"`
	BidPrice           float64 `json:"bidPrice,string"`
	BidQty             float64 `json:"bidQty,string"`
	AskPrice           float64 `json:"askPrice,string"`
	AskQty             float64 `json:"askQty,string"`
	OpenPrice          float64 `json:"openPrice,string"`
	HighPrice          float64 `json:"highPrice,string"`
	LowPrice           float64 `json:"lowPrice,string"`
	Volume             float64 `json:"volume,string"`
	QuoteVolume        float64 `json:"quoteVolume,string"`
	OpenTime           int64   `json:"openTime"`
	CloseTime          int64   `json:"closeTime"`
	FirstId            int     `json:"firstId"`
	LastId             int     `json:"lastId"`
	Count              int     `json:"count"`
}

type TickerKLine struct {
	Symbol    string
	Interval  string // e.g. 1m 3m 5m 15m 30m 1h 2h 4h 6h 8h 12h 1d 3d 1w 1M
	StartTime int64
	EndTime   int64
}

func (tp *TickerPrice) GetTickerPrice(ctx context.Context) *TickerPrice {
	if tp.Symbol == "" {
		log.Fatalln("need set symbol first")
	}

	body := http.GetTickerPrice(tp.Symbol)
	err := json.Unmarshal(body, tp)
	if err != nil {
		logger.Log.Errorf(ctx, "unmarshal fail when GetTickerPrice, err: %v", err)
	}

	return tp
}

func (tp *TickerPrice) GetAllTickerPrice(ctx context.Context) []*TickerPrice {
	body := http.GetTickerPrice("")
	err := json.Unmarshal(body, &tp.TickerPrices)
	if err != nil {
		logger.Log.Errorf(ctx, "unmarshal fail when GetAllTickerPrice, err: %v", err)
	}

	return tp.TickerPrices
}

func (th *Ticker24Hour) GetTicker24Hour(ctx context.Context) *Ticker24Hour {
	body := http.GetTicker24Hour(th.Symbol)
	err := json.Unmarshal(body, th)
	if err != nil {
		logger.Log.Errorf(ctx, "unmarshal fail when GetTicker24Hour, err: %v", err)
	}
	return th
}

func (tkl *TickerKLine) GetTickerKLine(ctx context.Context) *TickerKLine {
	body := http.GetTickerKLine(tkl.Symbol, tkl.Interval, tkl.StartTime, tkl.EndTime)
	err := json.Unmarshal(body, &tkl)
	if err != nil {
		logger.Log.Errorf(ctx, "unmarshal fail when GetAllTickerPrice, err: %v", err)
	}

	return tkl
}
