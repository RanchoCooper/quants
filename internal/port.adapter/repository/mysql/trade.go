package mysql

import (
    "context"
    "errors"

    "gorm.io/gorm"

    "quants/internal/domain.model/trade"
)

/**
 * @author Rancho
 * @date 2021/12/6
 */

type ITrade interface {
    GetTrades(context.Context) ([]*trade.Trade, error)
    GetTradesByUser(context.Context, string) ([]*trade.Trade, error)
    GetTradesByOrderId(context.Context, string) (*trade.Trade, error)
    InsertTrade(context.Context, *trade.Trade) (*trade.Trade, error)
}

type TradeRepo struct {
    db *gorm.DB
}

func NewTradeRepository(db *gorm.DB) *TradeRepo {
    return &TradeRepo{db: db}
}

func (t TradeRepo) GetTrades(ctx context.Context) ([]*trade.Trade, error) {
    var trades []*trade.Trade
    err := t.db.Find(&trades).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return trades, nil
}

func (t TradeRepo) GetTradesByUser(ctx context.Context, userEmail string) ([]*trade.Trade, error) {
    var trades []*trade.Trade
    err := t.db.Where("user_email = ?", userEmail).Find(&trades).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return trades, nil
}

func (t TradeRepo) GetTradesByOrderId(ctx context.Context, orderId string) (*trade.Trade, error) {
    var trade *trade.Trade
    err := t.db.Where("user_email = ?", orderId).Take(&trade).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return trade, nil
}

func (t TradeRepo) InsertTrade(ctx context.Context, trade *trade.Trade) (*trade.Trade, error) {
    err := t.db.Create(trade).Error
    if err != nil {
        return nil, err
    }

    return trade, nil
}

// TradeRepo implements the ITrade interface
var _ ITrade = &TradeRepo{}
