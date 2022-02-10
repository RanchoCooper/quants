package repo

import (
    "context"
)

/**
 * @author Rancho
 * @date 2022/2/10
 */

type ITrader interface {
    Backtest() bool
    Buy(ctx context.Context, symbol string, buyPrice, quantity float64) bool
    Sell(ctx context.Context, symbol string, sellPrice, quantity float64) bool
}
