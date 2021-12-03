package application

import (
    "context"
    "fmt"
    "time"

    "quants/internal/domain.model/binance"
    "quants/internal/domain.model/dingding"
    "quants/internal/domain.model/strategy/bet"
    "quants/util/logger"
)

/**
 * @author Rancho
 * @date 2021/12/4
 */

func EthRun(ctx context.Context) {
    b := &bet.BetData{}
    b.LoadFromJSON(context.Background())
    tp := &binance.TickerPrice{
        Symbol: b.GetCoinType(),
    }
    trader := &binance.Trader{}
    ding := &dingding.DingDing{}

    for {
        curMarketPrice := tp.GetTickerPrice(ctx).Price // 当前交易市价
        gridBuyPrice := b.GetBuyPrice()                // 当前网格买入价格
        gridSellPrice := b.GetSellPrice()              // 当前网格卖出价格
        buyQuantity := b.GetQuantity(ctx, bet.Buy)     // 买单量
        sellQuantity := b.GetQuantity(ctx, bet.Sell)   // 卖单量
        step := b.GetStep()                            // 当前步数

        // 满足买入价
        if gridBuyPrice >= curMarketPrice {
            tradeResp, ok := trader.Trade(ctx, &binance.TradeInfoVO{
                Symbol:   b.GetCoinType(),
                Side:     "BUY",
                Quantity: buyQuantity,
                Price:    gridBuyPrice,
            })
            if ok && tradeResp.OrderId != 0 {
                // 挂单成功
                ding.Message = fmt.Sprintf("：币种%s买单价为：%f。买单量为：%f", b.GetCoinType(), gridBuyPrice, buyQuantity)
                ding.SendDingMessage(ctx)
                b.ModifyPrice(ctx, gridBuyPrice, step+1)
                // 挂单后停止运行1分钟
                time.Sleep(1 * time.Minute)
            } else {
                ding.Message = fmt.Sprintf("：币种%s买单失败。api返回内容为:%s", b.GetCoinType(), tradeResp.Msg)
                ding.SendDingMessage(ctx)
                break
            }
        }

        // 满足卖出价
        if gridSellPrice < curMarketPrice {
            // 防止踏空，跟随价格上涨
            if step == 0 {
                b.ModifyPrice(ctx, gridSellPrice, step)
            } else {
                tradeResp, ok := trader.Trade(ctx, &binance.TradeInfoVO{
                    Symbol:   b.GetCoinType(),
                    Side:     "SELL",
                    Quantity: sellQuantity,
                    Price:    gridSellPrice,
                })
                if ok && tradeResp.OrderId != 0 {
                    // 挂单成功
                    ding.Message = fmt.Sprintf("：币种%s卖单价为：%f。卖买单量为：%f", b.GetCoinType(), gridBuyPrice, buyQuantity)
                    ding.SendDingMessage(ctx)
                    b.ModifyPrice(ctx, gridSellPrice, step-1)
                    // 挂单后停止运行1分钟
                    time.Sleep(1 * time.Minute)
                } else {
                    ding.Message = fmt.Sprintf("：币种%s卖买单失败。api返回内容为:%s", b.GetCoinType(), tradeResp.Msg)
                    ding.SendDingMessage(ctx)
                    break
                }
            }
        }

        logger.Log.Infof(ctx, "当前市价: %f。未能满足交易，继续运行", curMarketPrice)
    }
}
