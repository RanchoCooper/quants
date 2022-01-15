package http

import (
    "context"
    "flag"
    "testing"

    "quants/config"
    "quants/internal/adapter/repository"
    "quants/internal/domain/service"
)

/**
 * @author Rancho
 * @date 2022/1/7
 */

var ctx = context.Background()

func TestMain(m *testing.M) {
    if err := flag.Set("cf", "../../config/config.yaml"); err != nil {
        panic(err)
    }
    config.Init()
    repository.Init(
        repository.WithMySQL(),
        repository.WithRedis(),
    )
    service.Init(ctx)

    m.Run()
}
