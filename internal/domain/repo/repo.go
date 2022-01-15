package repo

import (
    "context"

    "gorm.io/gorm"

    "quants/internal/domain/entity"
)

/**
 * @author Rancho
 * @date 2021/12/24
 */

type IExampleRepo interface {
    Create(ctx context.Context, tx *gorm.DB, entity *entity.Example) (*entity.Example, error)
    Delete(ctx context.Context, tx *gorm.DB, Id int) error
    Save(ctx context.Context, tx *gorm.DB, entity *entity.Example) error
    Get(ctx context.Context, Id int) (entity *entity.Example, e error)
    FindByName(ctx context.Context, name string) (*entity.Example, error)
}

type IHealthCheckRepository interface {
    HealthCheck(ctx context.Context) error
}

type ITradeRepo interface {
    Create(context.Context, *entity.Trade) (*entity.Trade, error)
    Delete(context.Context, int64) error
    Update(ctx context.Context, user *entity.Trade) error
    Get(context.Context, int64) (*entity.Trade, error)
    FindByOrderID(context.Context, string) (*entity.Trade, error)
}

type IUserRepo interface {
    Create(context.Context, *entity.User) (*entity.User, error)
    Delete(context.Context, int64) error
    Update(ctx context.Context, user *entity.User) error
    Get(context.Context, int64) (*entity.User, error)
    FindByEmail(context.Context, string) (*entity.User, error)
}
