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
    assert.NoError(t, err)
    assert.NotEmpty(t, c)
    fmt.Println(c)
}
