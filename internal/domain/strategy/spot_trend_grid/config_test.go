package spot_trend_grid

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"
)

/**
 * @author Rancho
 * @date 2022/2/7
 */

func TestConfig_ReadFromFile(t *testing.T) {
    c := &Config{}
    err := c.ReadFromFile()
    fmt.Println(c)
    assert.NoError(t, err)
    assert.NotEmpty(t, c)
}

func TestConfig_GetBuyPrice(t *testing.T) {
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
