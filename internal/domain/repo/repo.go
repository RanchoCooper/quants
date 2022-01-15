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
    Create(ctx context.Context, tx *gorm.DB, entity *entity.Trade) (*entity.Trade, error)
    Delete(ctx context.Context, tx *gorm.DB, Id int) error
    Update(ctx context.Context, tx *gorm.DB, entity *entity.Trade) error
    Get(ctx context.Context, Id int) (entity *entity.Trade, e error)
    FindByOrderID(context.Context, string) (*entity.Trade, error)
}

type IUserRepo interface {
    Create(ctx context.Context, tx *gorm.DB, entity *entity.User) (*entity.User, error)
    Delete(ctx context.Context, tx *gorm.DB, Id int) error
    Update(ctx context.Context, tx *gorm.DB, entity *entity.User) error
    Get(ctx context.Context, Id int) (entity *entity.User, e error)
    FindByEmail(context.Context, string) (*entity.User, error)
}
