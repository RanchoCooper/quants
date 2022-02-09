package application

import (
    "testing"

    "quants/config"
    "quants/internal/domain/service"
)

/**
 * @author Rancho
 * @date 2022/2/8
 */

func TestSpotTrendGridLoop(t *testing.T) {
    service.Init(ctx)
    if config.Config.Env == "local" {
        t.Run("simulate", func(t *testing.T) {
            SpotTrendGridLoop(ctx, true)
        })

        t.Run("real", func(t *testing.T) {
            SpotTrendGridLoop(ctx, false)
        })
    }
}
