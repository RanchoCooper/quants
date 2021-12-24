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
    GetTrades(context.Context) ([]*entity.Trade, error)
    GetTradesByUser(context.Context, string) ([]*entity.Trade, error)
    GetTradesByOrderId(context.Context, string) (*entity.Trade, error)
    InsertTrade(context.Context, *entity.Trade) (*entity.Trade, error)
}
