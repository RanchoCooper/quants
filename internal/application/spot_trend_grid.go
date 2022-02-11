package application

import (
    "context"
    "fmt"
    "time"

    "quants/internal/adapter/dependency/http"
    "quants/internal/domain/repo"
    "quants/internal/domain/strategy/spot_trend_grid"
    "quants/util/logger"
)

/**
 * @author Rancho
 * @date 2022/2/8
 */

func SpotTrendGridLoop(ctx context.Context, trader repo.ITrader) {
    c := &spot_trend_grid.Config{}
    for {
        coinList := c.GetCoinList()
        for _, coinType := range coinList {
            // 当前网格买入价格
            gridBuyPrice := c.GetBuyPrice(coinType)
            // 当前网格卖出价格
            gridSellPrice := c.GetSellPrice(coinType)
            // 买入量
            gridBuyQuantity := c.GetQuantity(coinType, true)
            // 卖出量
            gridSellQuantity := c.GetQuantity(coinType, false)
            // 当前步数
            step := c.GetStep(coinType)

            // 当前交易市价
            var marketPrice float64
            if c.Backtest {
                if c.GetStartTime() >= c.GetEndTime() {
                    logger.Log.Infof(ctx, "请重新配置回测时间")
                    return
                }
                result := http.BinanceClinet.GetTickerKLine(ctx, coinType, c.GetInterval(), 1, c.GetStartTime(), c.GetEndTime())
                marketPrice = ((*result)[0].High - (*result)[0].Low) / 2
                c.UpdateStartTime()
            } else {
                marketPrice = http.BinanceClinet.GetTickerPrice(ctx, coinType).Price
            }

            if gridBuyPrice >= marketPrice {
                // 满足买入价

                ok := trader.Buy(ctx, coinType, gridBuyPrice, gridBuyQuantity)
                if ok {
                    // 买入成功
                    http.DingDingClient.SendDingDingMessage(fmt.Sprintf("买入成功。币种: %s, 数量: %f, 价格: %f", coinType, gridBuyQuantity, gridBuyPrice), false)
                    c.SetRatio(coinType)
                    c.SetRecordPrice(coinType, gridBuyPrice)
                    c.ModifyPrice(coinType, marketPrice, step+1, marketPrice)
                } else {
                    // 买入失败
                    http.DingDingClient.SendDingDingMessage(fmt.Sprintf("买入失败。币种: %s, 数量: %f, 价格: %f", coinType, gridBuyQuantity, gridBuyPrice), false)
                }

            } else if gridSellPrice < marketPrice {
                // 满足卖出价

                if step == 0 {
                    // 防止踏空，跟随价格上涨
                    c.ModifyPrice(coinType, gridSellPrice, step, marketPrice)
                } else {
                    lastPrice := c.GetRecordPrice(coinType)
                    sellAmount := c.GetQuantity(coinType, false)
                    profitUSDT := (marketPrice - lastPrice) * sellAmount // 预计盈利

                    ok := trader.Sell(ctx, coinType, gridSellPrice, gridSellQuantity)
                    if ok {
                        // 卖出成功
                        http.DingDingClient.SendDingDingMessage(fmt.Sprintf("卖出成功。币种: %s, 数量: %f, 价格: %f, 预计盈利: %f", coinType, gridSellQuantity, gridSellPrice, profitUSDT), false)
                        c.SetRatio(coinType)
                        c.ModifyPrice(coinType, marketPrice, step-1, marketPrice)
                        c.RemoveRecordPrice(coinType)
                    } else {
                        // 卖出失败
                        http.DingDingClient.SendDingDingMessage(fmt.Sprintf("卖出失败。币种: %s, 数量: %f, 价格: %f", coinType, gridSellQuantity, gridSellPrice), false)
                    }
                }
            } else {
                logger.Log.Infof(ctx, "未满足交易。币种: %s, 当前市价: %f, 买入价: %f, 卖出价: %f, 等待下次运行", coinType, marketPrice, gridBuyPrice, gridSellPrice)
            }
            if trader.Backtest() {
                time.Sleep(time.Second * 2)
            } else {
                time.Sleep(time.Minute)
            }
        }
    }
}
