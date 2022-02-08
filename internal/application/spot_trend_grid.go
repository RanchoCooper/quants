package application

import (
    "context"
    "fmt"
    "time"

    "quants/internal/adapter/dependency/http"
    "quants/internal/domain/strategy/spot_trend_grid"
)

/**
 * @author Rancho
 * @date 2022/2/8
 */

func SpotTrendGridLoop(ctx context.Context, isSimulate bool) {
    c := &spot_trend_grid.Config{}
    for {
        for _, coinType := range c.GetCoinList() {
            // 当前网格买入价格
            gridBuyPrice := c.GetBuyPrice(coinType)
            // 当前网格卖出价格
            gridSellPrice := c.GetSellPrice(coinType)
            // 买入量
            quantity := c.GetQuantity(coinType, true)
            // 当前步数
            step := c.GetStep(coinType)
            // 当前交易市价
            marketPrice := http.BinanceClinet.GetTickerPrice(ctx, coinType).Price

            if gridBuyPrice >= marketPrice {
                // 满足买入价

                if !isSimulate {
                    result := http.BinanceClinet.TradeLimit(ctx, coinType, "BUY", &quantity, &gridBuyPrice)
                    if result.OrderId != 0 {
                        // 下单成功
                        http.DingDingClient.SendDingDingMessage(fmt.Sprintf("买入成功。币种: %s, 数量: %f, 价格: %f", coinType, quantity, gridBuyPrice), false)
                        c.SetRatio(coinType)
                        c.SetRecordPrice(coinType, result.Price)
                        c.ModifyPrice(coinType, marketPrice, step+1, marketPrice)
                        // 停止运行1min
                        time.Sleep(time.Minute)
                    } else {
                        http.DingDingClient.SendDingDingMessage(fmt.Sprintf("买入失败。币种: %s, 数量: %f, 价格: %f", coinType, quantity, gridBuyPrice), false)
                        break
                    }
                } else {
                    // handle simulate
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

                    if !isSimulate {
                        result := http.BinanceClinet.TradeLimit(ctx, coinType, "SELL", &quantity, &gridBuyPrice)
                        if result.OrderId != 0 {
                            // 下单成功
                            http.DingDingClient.SendDingDingMessage(fmt.Sprintf("卖出成功。币种: %s, 数量: %f, 价格: %f, 预计盈利: %f", coinType, quantity, gridBuyPrice, profitUSDT), false)
                            c.SetRatio(coinType)
                            c.ModifyPrice(coinType, marketPrice, step-1, marketPrice)
                            c.RemoveRecordPrice(coinType)
                            // 停止运行1min
                            time.Sleep(time.Minute)
                        } else {
                            http.DingDingClient.SendDingDingMessage(fmt.Sprintf("卖出失败。币种: %s, 数量: %f, 价格: %f", coinType, quantity, gridBuyPrice), false)
                            break
                        }
                    } else {
                        // handle simulate
                    }
                }
            }

        }
    }
}
