package http

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

/**
 * @author Rancho
 * @date 2021/12/1
 */

func TestPing(t *testing.T) {
    assert.Equal(t, "{}", Ping())
}

func TestGetTickerPrice(t *testing.T) {
    body := GetTickerPrice("ETHBTC")
    assert.NotEmpty(t, body)
}

func TestGetTicker24Hour(t *testing.T) {
    body := GetTicker24Hour("ETHBTC")
    assert.NotEmpty(t, body)
}

func TestGetTickerKLine(t *testing.T) {
    body := GetTickerKLine("ETHBTC", "1M", 0, 0)
    assert.NotEmpty(t, body)
}

func TestTradeLimit(t *testing.T) {
    quantity := 1.0
    price:= 0.00000001
    body := TradeLimit("ETHBTC", "BUY", &quantity, &price)
    // FIXME
    // {"code":-1021,"msg":"Timestamp for this request is outside of the recvWindow."}
    assert.NotEmpty(t, body)
}