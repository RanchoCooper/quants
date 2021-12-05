package application

import (
    "context"
    "time"

    "quants/internal/domain.model/binance"
    "quants/internal/domain.model/dingding"
    "quants/internal/domain.model/strategy/grid/bothway"
)

/**
 * @author Rancho
 * @date 2021/12/5
 */

func GridBothwayRun(ctx context.Context) {
    gridBothway := bothway.NewGridBothway()
    gridBothway.Grid.LoadFromJSON(context.Background())
    b := &binance.Binance{TickerPrice: binance.TickerPrice{
        Symbol: gridBothway.Grid.GetCoinType(),
    }}
    ding := &dingding.DingDing{}

    for {
        curMarketPrice := b.GetTickerPrice(ctx).Price // 当前交易市价

        // 超出范围不允许网格策略
        if curMarketPrice < gridBothway.FloorPrice || curMarketPrice > gridBothway.CeilPrice {
            time.Sleep(10 * time.Minute)
        }

        // 开多
        if gridBothway.SpotBuyPrice < curMarketPrice { // FIXME index.calcAngle(self.coinType,"5m",False,right_size)
            // 做多买入，趋势拉升，不买入
        }

        // 开空

        // 平多

        // 平空

        time.Sleep(1 * time.Minute)
    }
}
