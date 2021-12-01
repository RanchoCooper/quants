package binance

import (
    "context"
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"
)

/**
 * @author Rancho
 * @date 2021/12/1
 */

func TestTickerPrice_GetTickerPrice(t *testing.T) {
    tk := &TickerPrice{
        Symbol: "ETHBTC",
    }
    tk.GetTickerPrice(context.Background())
    assert.NotEmpty(t, tk.Price)
}

func TestTickerPrice_GetTicker24Hour(t *testing.T) {
    tk := &Ticker24Hour{
        Symbol: "ETHBTC",
    }
    tk.GetTicker24Hour(context.Background())
    fmt.Printf("%v", tk)
    assert.NotEmpty(t, tk.PriceChange)
}