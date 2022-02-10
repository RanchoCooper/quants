package spot_trend_grid

import (
    "context"

    "quants/internal/adapter/dependency/http"
    "quants/internal/domain/entity"
    "quants/internal/domain/repo"
)

/**
 * @author Rancho
 * @date 2022/2/10
 */

type Trader struct {
    Conf *Config
}

var _ repo.ITrader = &Trader{}

func NewTrader() *Trader {
    return &Trader{
        Conf: &Config{},
    }
}

func (t *Trader) Backtest() bool {
    return false
}

func (t *Trader) Buy(ctx context.Context, symbol string, price, quantity float64) bool {
    result := http.BinanceClinet.TradeLimit(ctx, symbol, entity.TradeSideBuy, &quantity, &price)
    if result.OrderId != 0 {
        // 下单成功
        return true
    } else {
        // 下单失败
        return false
    }
}

func (t *Trader) Sell(ctx context.Context, symbol string, price, quantity float64) bool {
    result := http.BinanceClinet.TradeLimit(ctx, symbol, entity.TradeSideSell, &quantity, &price)
    if result.OrderId != 0 {
        // 下单成功
        return true
    } else {
        // 下单失败
        return false
    }

}
