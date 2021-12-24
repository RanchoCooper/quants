package bet

import (
    "quants/internal/domain.model/service/strategy/grid"
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
