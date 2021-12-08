package bothway

import (
    "quants/internal/domain.model/strategy/grid"
    "quants/util"
)

/**
 * @author Rancho
 * @date 2021/12/5
 */

const gridBothwayJSONFile = "bothway.json"

type GridBothway struct {
    Grid grid.Grid
}

func NewGridBothway() *GridBothway {
    return &GridBothway{
        Grid: grid.Grid{
            ConfigJSONFile: util.GetCurrentPath() + "/" + gridBothwayJSONFile,
        },
    }
}
