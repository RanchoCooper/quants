package binance

import (
    "context"
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"
)

/**
 * @author Rancho
 * @date 2021/12/3
 */

func TestTrader_Trade(t *testing.T) {
    vo := &TradeInfoVO{
        Symbol:   "ETHBTC",
        Side:     "BUY",
        Quantity: 1.0,
        Price:    0.0000001,
    }
    trader := &Trader{}
    resp, _ := trader.Trade(context.Background(), vo)
    // FIXME
    fmt.Printf("%+v", resp)
    assert.NotEmpty(t, resp)
}
