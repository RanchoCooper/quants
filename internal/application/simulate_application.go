package application

import (
    "context"
    "time"

    "quants/internal/domain.model/binance"
    "quants/internal/domain.model/strategy/grid"
    "quants/internal/domain.model/strategy/grid/bet"
    "quants/internal/domain.model/trade"
    "quants/util"
    "quants/util/logger"
)

/**
 * @author Rancho
 * @date 2021/12/6
 */

func SimulateGridBetRun(ctx context.Context) {
    gb := bet.NewGridBet()
    gb.Grid.LoadFromJSON(ctx)
    tp := &binance.TickerPrice{
        Symbol: gb.Grid.GetCoinType(),
    }

    for {
        curMarketPrice := tp.GetTickerPrice(ctx).Price // 当前交易市价
        gridBuyPrice := gb.Grid.GetBuyPrice()          // 当前网格买入价格
        gridSellPrice := gb.Grid.GetSellPrice()        // 当前网格卖出价格
        buyQuantity := gb.Grid.GetQuantity(grid.Buy)   // 买单量
        sellQuantity := gb.Grid.GetQuantity(grid.Sell) // 卖单量
        step := gb.Grid.GetStep()                      // 当前步数

        // 满足买入价
        if gb.Grid.ShouldBuy(curMarketPrice) {
            // 买入
            logger.Log.Infof(ctx, "simulate meet buy price, price: %f", gridBuyPrice)
            t := &trade.Trade{
                UserEmail:  SimulateUserEmail,
                Symbol:     gb.Grid.GetCoinType(),
                OrderId:    "simulate-" + util.RandString(10, false),
                Type:       trade.TypeBuy,
                Price:      gridBuyPrice,
                Quantity:   buyQuantity,
                IsSimulate: true,
            }
            _, err := t.InsertTrade(ctx)
            if err != nil {
                logger.Log.Errorf(ctx, "simulateGridBetRun buy fail when InsertTrade, err: %s", err.Error())
                return
            }
        } else if gb.Grid.ShouldSell(curMarketPrice) {
            // 满足卖出价
            // 防止踏空，跟随价格上涨
            if step == 0 {
                gb.Grid.AdjustPrice(ctx, gridSellPrice, step)
            } else {
                // 卖出
                logger.Log.Infof(ctx, "simulate meet sell price, price: %f", gridSellPrice)
                t := &trade.Trade{
                    UserEmail:  SimulateUserEmail,
                    Symbol:     gb.Grid.GetCoinType(),
                    OrderId:    "simulate-" + util.RandString(10, false),
                    Type:       trade.TypeSell,
                    Price:      gridSellPrice,
                    Quantity:   sellQuantity,
                    IsSimulate: true,
                }
                _, err := t.InsertTrade(ctx)
                if err != nil {
                    logger.Log.Errorf(ctx, "simulateGridBetRun sell fail when InsertTrade, err: %s", err.Error())
                    return
                }
            }
        } else {
            logger.Log.Infof(ctx, "当前市价: %f。未能满足交易，继续运行", curMarketPrice)
        }

        time.Sleep(2 * time.Second)
    }
}
