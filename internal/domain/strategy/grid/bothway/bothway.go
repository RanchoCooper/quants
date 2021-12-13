package bothway

import (
    "quants/internal/domain/strategy/grid"
    "quants/util"
)

/**
 * @author Rancho
 * @date 2021/12/5
 */

const gridBothwayJSONFile = "bothway.json"

type GridBothway struct {
    Grid grid.Grid

    SpotBuyPrice      float64       `json:"spot_buy_price"`  // 现货买入价格
    SpotSellPrice     float64       `json:"spot_sell_price"` // 现货卖出价格
    SpotStep          int           `json:"spot_step"`       // 当前现货步数
    RecordSpotPrice   []interface{} `json:"record_spot_price"`
    FutureBuyPrice    float64       `json:"future_buy_price"`
    FutureSellPrice   float64       `json:"future_sell_price"`
    FutureStep        int           `json:"future_step"`
    RecordFuturePrice []interface{} `json:"record_future_price"`
    ProfitRatio       float64       `json:"profit_ratio"`
    DoubleThrowRatio  float64       `json:"double_throw_ratio"`
    Cointype          string        `json:"cointype"`
    FloorPrice        float64       `json:"floor_price"`
    CeilPrice         float64       `json:"ceil_price"`
    SpotQuantity      []float64     `json:"spot_quantity"`
    FutureQuantity    []float64     `json:"future_quantity"`
}

func NewGridBothway() *GridBothway {
    return &GridBothway{
        Grid: grid.Grid{
            ConfigJSONFile: util.GetCurrentPath() + "/" + gridBothwayJSONFile,
        },
    }
}
func (b *GridBothway) GetQuantity(action grid.ExchangeType) float64 {
    var curStep int
    if action == grid.Buy {
        curStep = b.SpotStep
    }
    if action == grid.Sell {
        curStep = b.SpotStep - 1
    }
    quantity := 0.0
    if curStep < len(b.SpotQuantity) {
        if curStep == 0 {
            quantity = b.SpotQuantity[0]
        } else {
            quantity = b.SpotQuantity[curStep]
        }
    } else {
        // 当前仓位大于设置的仓位，取最后一位
        quantity = b.SpotQuantity[len(b.SpotQuantity)-1]
    }

    return quantity
}
