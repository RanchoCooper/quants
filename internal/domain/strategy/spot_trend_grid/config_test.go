package spot_trend_grid

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"

    "quants/config"
)

/**
 * @author Rancho
 * @date 2022/2/7
 */

func TestConfig_ReadFromFile(t *testing.T) {
    if config.Config.Env == "local" {
        c := &Config{}
        err := c.ReadFromFile()
        fmt.Println(c)
        assert.NoError(t, err)
        assert.NotEmpty(t, c)
    }
}

func TestConfig_GetBuyPrice(t *testing.T) {
    if config.Config.Env == "local" {
        t.Run("basic", func(t *testing.T) {
            c := &Config{}
            price := c.GetBuyPrice("BNBUSDT")
            assert.NotEmpty(t, price)
        })

        t.Run("not exists", func(t *testing.T) {
            c := &Config{}
            price := c.GetBuyPrice("not exists")
            assert.Empty(t, price)
        })
    }

}

func TestConfig_SetRatio(t *testing.T) {
    t.SkipNow()
    c := &Config{}
    c.SetRatio("BNBUSDT")
}

func TestConfig_SetRecordPrice(t *testing.T) {
    t.SkipNow()
    c := &Config{}
    c.SetRecordPrice("BNBUSDT", 430.2)
}

func TestConfig_RemoveRecordPrice(t *testing.T) {
    t.SkipNow()
    c := &Config{}
    c.RemoveRecordPrice("BNBUSDT")
}

func TestConfig_GetStartTime(t *testing.T) {
    t.SkipNow()
    c := &Config{}
    fmt.Println(c.GetStartTime())
}

func TestConfig_GetEndTime(t *testing.T) {
    t.SkipNow()
    c := &Config{}
    fmt.Println(c.GetEndTime())
}

func TestConfig_UpdateStartTime(t *testing.T) {
    c := &Config{}
    c.UpdateStartTime()
}
