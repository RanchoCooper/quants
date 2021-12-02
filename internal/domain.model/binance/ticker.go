package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

    "quants/internal/port.adapter/dependency/http"
    "quants/util/logger"

    "github.com/spf13/cast"
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

type TickerKLineQueryVO struct {
	Symbol    string
	Interval  string // e.g. 1m 3m 5m 15m 30m 1h 2h 4h 6h 8h 12h 1d 3d 1w 1M
	StartTime int64
	EndTime   int64
}

type TickerKLine struct {
	OpenTime                 int64
	Open                     float64
	High                     float64
	Low                      float64
	Close                    float64
	Volume                   float64
	CloseTime                int64
	QuoteAssetVolume         float64
	NumberOfTrades           int
	TakerBuyBaseAssetVolume  float64
	TakerBuyQuoteAssetVolume float64
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

func (tkl *TickerKLine) UnmarshalJSON(b []byte) error {
	// We're deserializing into a struct, but in JSON it's a mixed-type array.
	var arr []interface{}
	err := json.Unmarshal(b, &arr)
	if err != nil {
		return fmt.Errorf("unmarshal Item underlying array: %w", err)
	}
	// JSON numbers will become float64 when loaded into interface{} but we want int
	tkl.OpenTime = cast.ToInt64(arr[0])
	tkl.Open = cast.ToFloat64(arr[1])
	tkl.High = cast.ToFloat64(arr[2])
	tkl.Low = cast.ToFloat64(arr[3])
	tkl.Close = cast.ToFloat64(arr[4])
	tkl.Volume = cast.ToFloat64(arr[5])
	tkl.CloseTime = cast.ToInt64(arr[6])
	tkl.QuoteAssetVolume = cast.ToFloat64(arr[7])
	tkl.NumberOfTrades = cast.ToInt(arr[8])
	tkl.TakerBuyBaseAssetVolume = cast.ToFloat64(arr[9])
	tkl.TakerBuyQuoteAssetVolume = cast.ToFloat64(arr[10])
	return nil
}

func (tkl *TickerKLineQueryVO) GetTickerKLine(ctx context.Context) []*TickerKLine {
	result := make([]*TickerKLine, 0)
	body := http.GetTickerKLine(tkl.Symbol, tkl.Interval, tkl.StartTime, tkl.EndTime)
	err := json.Unmarshal(body, &result)
	if err != nil {
		logger.Log.Errorf(ctx, "unmarshal fail when GetTickerKLine, err: %v", err)
	}

	return result
}
