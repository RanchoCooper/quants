package application

import (
    "testing"

    "quants/config"
    "quants/internal/domain/service"
    "quants/internal/domain/strategy/spot_trend_grid"
)

/**
 * @author Rancho
 * @date 2022/2/8
 */

func TestSpotTrendGridLoop(t *testing.T) {
    service.Init(ctx)
    realTrader := spot_trend_grid.NewTrader()
    simulateTrader := service.NewSimulatorService(ctx)
    if config.Config.Env == "local" {
        t.Run("simulate", func(t *testing.T) {
            SpotTrendGridLoop(ctx, simulateTrader)
        })

        t.Run("real", func(t *testing.T) {
            SpotTrendGridLoop(ctx, realTrader)
        })
    }
}
