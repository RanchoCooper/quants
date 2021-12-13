package grid

import (
    "context"
    "fmt"
    "time"

    "quants/internal/domain/binance"
    "quants/internal/domain/dingding"
    "quants/internal/domain/strategy/grid"
    "quants/internal/domain/strategy/grid/bet"
    "quants/util/logger"
)

/**
 * @author Rancho
 * @date 2021/12/4
 */

func GridBetRun(ctx context.Context) {
    gb := bet.NewGridBet()
    gb.Grid.LoadFromJSON(context.Background())
    b := &binance.Binance{
        TickerPrice: binance.TickerPrice{
            Symbol: gb.Grid.GetCoinType(),
        },
    }
    ding := &dingding.DingDing{}

    for {
        curMarketPrice := b.GetTickerPrice(ctx).Price  // 当前交易市价
        gridBuyPrice := gb.Grid.GetBuyPrice()          // 当前网格买入价格
        gridSellPrice := gb.Grid.GetSellPrice()        // 当前网格卖出价格
        buyQuantity := gb.Grid.GetQuantity(grid.Buy)   // 买单量
        sellQuantity := gb.Grid.GetQuantity(grid.Sell) // 卖单量
        step := gb.Grid.GetStep()                      // 当前步数

        // 满足买入价
        if gb.Grid.ShouldBuy(curMarketPrice) {
            tradeResp, ok := b.Trade(ctx, &binance.TradeInfoVO{
                Symbol:   gb.Grid.GetCoinType(),
                Side:     "BUY",
                Quantity: buyQuantity,
                Price:    gridBuyPrice,
            })
            if ok && tradeResp.OrderId != 0 {
                // 挂单成功
                ding.Message = fmt.Sprintf("：币种%s买单价为：%f。买单量为：%f", gb.Grid.GetCoinType(), gridBuyPrice, buyQuantity)
                ding.SendDingMessage(ctx)
                gb.Grid.AdjustPrice(ctx, gridBuyPrice, step+1)
                // 挂单后停止运行1分钟
                time.Sleep(1 * time.Minute)
            } else {
                ding.Message = fmt.Sprintf("：币种%s买单失败。api返回内容为:%s", gb.Grid.GetCoinType(), tradeResp.Msg)
                ding.SendDingMessage(ctx)
                break
            }
        }

        // 满足卖出价
        if gb.Grid.ShouldSell(curMarketPrice) {
            // 防止踏空，跟随价格上涨
            if step == 0 {
                gb.Grid.AdjustPrice(ctx, gridSellPrice, step)
            } else {
                tradeResp, ok := b.Trade(ctx, &binance.TradeInfoVO{
                    Symbol:   gb.Grid.GetCoinType(),
                    Side:     "SELL",
                    Quantity: sellQuantity,
                    Price:    gridSellPrice,
                })
                if ok && tradeResp.OrderId != 0 {
                    // 挂单成功
                    ding.Message = fmt.Sprintf("：币种%s卖单价为：%f。卖买单量为：%f", gb.Grid.GetCoinType(), gridBuyPrice, buyQuantity)
                    ding.SendDingMessage(ctx)
                    gb.Grid.AdjustPrice(ctx, gridSellPrice, step-1)
                    // 挂单后停止运行1分钟
                    time.Sleep(1 * time.Minute)
                } else {
                    ding.Message = fmt.Sprintf("：币种%s卖买单失败。api返回内容为:%s", gb.Grid.GetCoinType(), tradeResp.Msg)
                    ding.SendDingMessage(ctx)
                    break
                }
            }
        }
        logger.Log.Infof(ctx, "当前市价: %f。未能满足交易，一分钟后继续运行", curMarketPrice)

        time.Sleep(1 * time.Minute)
    }
}
