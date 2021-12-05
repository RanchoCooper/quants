package bet

import (
    "quants/internal/domain.model/strategy/grid"
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
            ConfigJSONFile: gridBetJSONFile,
        },
    }
}
