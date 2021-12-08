package bet

import (
    "quants/internal/domain.model/strategy/grid"
    "quants/util"
)

/**
 * @author Rancho
 * @date 2021/12/3
 */

const gridBetJSONFile = "bet.json"

type GridBet struct {
    Grid grid.Grid
}

func NewGridBet() *GridBet {
    return &GridBet{
        Grid: grid.Grid{
            ConfigJSONFile: util.GetCurrentPath() + "/" + gridBetJSONFile,
        },
    }
}

func (gb GridBet) ShouldBuy(curMarketPrice float64) bool {
    return gb.Grid.GetBuyPrice() > curMarketPrice
}

func (gb GridBet) ShouldSell(curMarketPrice float64) bool {
    return gb.Grid.GetSellPrice() < curMarketPrice
}
