package application

import (
    "context"

    "quants/internal/domain.model/binance"
    "quants/internal/domain.model/strategy/grid"
    "quants/internal/domain.model/strategy/grid/bet"
    "quants/internal/domain.model/trade"
    "quants/internal/port.adapter/repository"
    "quants/util"
    "quants/util/logger"
)

/**
 * @author Rancho
 * @date 2021/12/6
 */

func SimulateGridBetRun(ctx context.Context) {
    gridBet := bet.NewGridBet()
    gridBet.Grid.LoadFromJSON(ctx)
    tp := &binance.TickerPrice{
        Symbol: gridBet.Grid.GetCoinType(),
    }

    for {
        curMarketPrice := tp.GetTickerPrice(ctx).Price      // 当前交易市价
        gridBuyPrice := gridBet.Grid.GetBuyPrice()          // 当前网格买入价格
        gridSellPrice := gridBet.Grid.GetSellPrice()        // 当前网格卖出价格
        buyQuantity := gridBet.Grid.GetQuantity(grid.Buy)   // 买单量
        sellQuantity := gridBet.Grid.GetQuantity(grid.Sell) // 卖单量
        step := gridBet.Grid.GetStep()                      // 当前步数

        // 满足买入价
        if gridBuyPrice >= curMarketPrice {
            // 买入
            t := &trade.Trade{
                UserEmail:  "rancho@simulate.com",
                Symbol:     gridBet.Grid.GetCoinType(),
                OrderId:    "simulate-" + util.RandString(10, false),
                Type:       trade.TypeBuy,
                Price:      gridBuyPrice,
                Quantity:   buyQuantity,
                IsSimulate: true,
            }
            t, err := repository.MySQL.Trade.InsertTrade(ctx, t)
            if err != nil {
                logger.Log.Errorf(ctx, "simulateGridBetRun buy fail when InsertTrade, err: %s", err.Error())
                return
            }
        }

        // 满足卖出价
        if gridSellPrice < curMarketPrice {
            // 防止踏空，跟随价格上涨
            if step == 0 {
                gridBet.Grid.ModifyPrice(ctx, gridSellPrice, step)
            } else {
                // 卖出
                t := &trade.Trade{
                    UserEmail: "rancho@simulate.com",
                    Symbol:    gridBet.Grid.GetCoinType(),
                    OrderId:   "simulate-" + util.RandString(10, false),
                    Type:      trade.TypeSell,
                    Price:     gridSellPrice,
                    Quantity:  sellQuantity,
                }
                t, err := repository.MySQL.Trade.InsertTrade(ctx, t)
                if err != nil {
                    logger.Log.Errorf(ctx, "simulateGridBetRun sell fail when InsertTrade, err: %s", err.Error())
                    return
                }
            }
        }

        logger.Log.Infof(ctx, "当前市价: %f。未能满足交易，继续运行", curMarketPrice)
    }

}
