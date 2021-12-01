package binance

import (
    "context"
    "encoding/json"
    "strconv"

    "go-hexagonal/internal/port.adapter/dependency/http"
    "go-hexagonal/util/logger"
)

/**
 * @author Rancho
 * @date 2021/12/1
 */

type TickerPrice struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}

type TickerPriceVO struct {
    Symbol string  `json:"symbol"`
    Price  string `json:"price"`
}

func (tp *TickerPrice) GetTickerPrice(ctx context.Context) {
    body := http.GetTickerPrice(tp.Symbol)
    vo := &TickerPriceVO{}
    err := json.Unmarshal(body, vo)
    if err != nil {
        logger.Log.Errorf(ctx, "unmarshal fail when GetTickerPrice, err: %v", err)
    }
    price, err := strconv.ParseFloat(vo.Price, 64)
    if err != nil {
        logger.Log.Errorf(ctx, "convert string to float64 fail when GetTickerPrice, price: %s, err: %v", vo.Price, err)
    }
    tp.Price = price
}
