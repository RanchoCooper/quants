package binance

import (
    "context"
    "testing"

    "github.com/stretchr/testify/assert"
)

/**
 * @author Rancho
 * @date 2021/12/1
 */

func TestTickerPrice_GetTickerPrice(t *testing.T) {
    tp := &TickerPrice{
        Symbol: "ETHBTC",
    }
    tp.GetTickerPrice(context.Background())
    assert.NotEmpty(t, tp.Price)
}
