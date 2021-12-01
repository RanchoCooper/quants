package binance

import (
	"context"
	"encoding/json"
	"fmt"

	"go-hexagonal/internal/port.adapter/dependency/http"
	"go-hexagonal/util/logger"
)

/**
 * @author Rancho
 * @date 2021/12/1
 */

type TickerPrice struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
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

func (tp *TickerPrice) GetTickerPrice(ctx context.Context) {
	body := http.GetTickerPrice(tp.Symbol)
	err := json.Unmarshal(body, tp)
	if err != nil {
		logger.Log.Errorf(ctx, "unmarshal fail when GetTickerPrice, err: %v", err)
	}
}

func (th *Ticker24Hour) GetTicker24Hour(ctx context.Context) {
	body := http.GetTicker24Hour(th.Symbol)
	fmt.Println(string(body))
	err := json.Unmarshal(body, th)
	if err != nil {
		logger.Log.Errorf(ctx, "unmarshal fail when GetTicker24Hour, err: %v", err)
	}
}
