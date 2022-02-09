package service

import (
    "context"
    "sync"
)

/**
 * @author Rancho
 * @date 2021/12/10
 */

var (
    once         sync.Once
    ExampleSvc   *ExampleService
    UserSvc      *UserService
    SimulatorSvc *SimulatorService
)

func Init(ctx context.Context) {
    once.Do(func() {
        ExampleSvc = NewExampleService(ctx)
        UserSvc = NewUserService(ctx)
        SimulatorSvc = NewSimulatorService(ctx)
    })
}
