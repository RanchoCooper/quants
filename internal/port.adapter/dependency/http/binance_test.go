package http

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

/**
 * @author Rancho
 * @date 2021/12/1
 */

func TestPing(t *testing.T) {
    assert.Equal(t, "{}", Ping())
}

func TestGetTickerPrice(t *testing.T) {
    result := GetTickerPrice("ETHBTC")
    assert.NotEmpty(t, result)
}