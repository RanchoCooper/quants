package binance

import (
    "context"
    "testing"

    "github.com/stretchr/testify/assert"
)

/**
 * @author Rancho
 * @date 2022/1/8
 */

func TestTickerPrice_GetTickerPrice(t *testing.T) {
    tk := &TickerPrice{
        Symbol: "ETHBTC",
    }
    tk.GetTickerPrice(context.Background())
    assert.NotEmpty(t, tk.Price)
}

func TestGetAllTickerPrice(t *testing.T) {
    tp := &TickerPrice{}
    tp.GetAllTickerPrice(context.Background())
    assert.NotEmpty(t, tp.TickerPrices)
}

func TestTickerPrice_GetTicker24Hour(t *testing.T) {
    tk := &Ticker24Hour{
        Symbol: "ETHBTC",
    }
    tk.GetTicker24Hour(context.Background())
    assert.NotEmpty(t, tk.PriceChange)
}

func TestTickerKLine_GetTickerKLine(t *testing.T) {
    vo := &TickerKLineQueryVO{
        Symbol:   "ETHBTC",
        Interval: "1M",
    }
    tkl := TickerKLine{}
    r := tkl.GetTickerKLine(context.Background(), vo)
    assert.NotEmpty(t, len(r))
    assert.NotEmpty(t, r[0])
}