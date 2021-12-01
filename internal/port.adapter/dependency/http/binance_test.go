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
