package http

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"

    "quants/internal/domain/vo"
)

/**
 * @author Rancho
 * @date 2021/12/1
 */

func TestPing(t *testing.T) {
    binance := NewBinanceAPI()
    assert.Equal(t, &vo.PingResp{}, binance.Ping())
}

func TestGetTickerPrice(t *testing.T) {
    binance := NewBinanceAPI()
    body := binance.GetTickerPrice("ETHBTC")
    fmt.Println(body)
    assert.NotEmpty(t, body)
}

func TestGetTicker24Hour(t *testing.T) {
    binance := NewBinanceAPI()
    body := binance.GetTicker24Hour("ETHBTC")
    fmt.Println(body)
    assert.NotEmpty(t, body)
}

func TestGetTickerKLine(t *testing.T) {
    binance := NewBinanceAPI()
    body := binance.GetTickerKLine("ETHBTC", "1M", 0, 0)
    fmt.Println(body)
    assert.NotEmpty(t, body)
}

func TestTradeLimit(t *testing.T) {
    binance := NewBinanceAPI()
    quantity := 1.0
    price := 0.00000001
    body := binance.TradeLimit("ETHBTC", "BUY", &quantity, &price)
    // FIXME
    // {"code":-1021,"msg":"Timestamp for this request is outside of the recvWindow."}
    fmt.Println(string(body))
    assert.NotEmpty(t, body)
}
