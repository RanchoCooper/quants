package application

import (
    "context"
    "testing"

    "quants/config"
)

/**
 * @author Rancho
 * @date 2022/2/8
 */

func TestSpotTrendGridLoop(t *testing.T) {
    if config.Config.Env == "local" {
        t.Run("simulate", func(t *testing.T) {
            SpotTrendGridLoop(context.Background(), true)
        })

        t.Run("real", func(t *testing.T) {
            SpotTrendGridLoop(context.Background(), false)
        })
    }
}
