package mysql

import (
    "context"
    "errors"

    "gorm.io/gorm"

    "quants/internal/domain.model/entity"
)

/**
 * @author Rancho
 * @date 2021/12/6
 */

type TradeRepo struct {
    db *gorm.DB
}

func NewTradeRepo(db *gorm.DB) *TradeRepo {
    return &TradeRepo{db: db}
}

func (t TradeRepo) GetTrades(ctx context.Context) ([]*entity.Trade, error) {
    var trades []*entity.Trade
    err := t.db.Find(&trades).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return trades, nil
}

func (t TradeRepo) GetTradesByUser(ctx context.Context, userEmail string) ([]*entity.Trade, error) {
    var trades []*entity.Trade
    err := t.db.Where("user_email = ?", userEmail).Find(&trades).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return trades, nil
}

func (t TradeRepo) GetTradesByOrderId(ctx context.Context, orderId string) (*entity.Trade, error) {
    var trade *entity.Trade
    err := t.db.Where("user_email = ?", orderId).Take(&trade).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return trade, nil
}

func (t TradeRepo) InsertTrade(ctx context.Context, trade *entity.Trade) (*entity.Trade, error) {
    err := t.db.Create(trade).Error
    if err != nil {
        return nil, err
    }

    return trade, nil
}

// TradeRepo implements the ITradeRepo interface
var _ entity.ITradeRepo = &TradeRepo{}
