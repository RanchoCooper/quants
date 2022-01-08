package service

import (
    "context"
    "testing"

    "quants/internal/port.adapter/repository"
)

/**
 * @author Rancho
 * @date 2021/12/25
 */

var ctx = context.Background()

func TestMain(m *testing.M) {
    repository.Init(
        repository.WithMySQL(),
        repository.WithRedis(),
    )
    m.Run()
}
