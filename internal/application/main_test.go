package application

import (
    "context"
    "testing"

    "quants/internal/adapter/repository"
    "quants/internal/domain/service"
)

/**
 * @author Rancho
 * @date 2022/2/9
 */

var ctx = context.Background()

func TestMain(m *testing.M) {
    repository.Init(
        repository.WithMySQL(),
        repository.WithRedis(),
    )
    service.Init(ctx)
    m.Run()
}
