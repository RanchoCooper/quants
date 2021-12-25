package repo

import (
    "context"

    "quants/internal/domain.model/entity"
)

/**
 * @author Rancho
 * @date 2021/12/23
 */

type ITradeRepo interface {
    Create(context.Context, *entity.Trade) (*entity.Trade, error)
    Delete(context.Context, int64) error
    Update(ctx context.Context, user *entity.Trade) error
    Get(context.Context, int64) (*entity.Trade, error)
    FindByOrderID(context.Context, string) (*entity.Trade, error)
}
