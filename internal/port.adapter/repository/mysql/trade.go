package mysql

import (
    "context"
    "errors"

    "gorm.io/gorm"

    "quants/internal/domain/model/trade"
)

/**
 * @author Rancho
 * @date 2021/12/6
 */

// @autowire(set=repository,trade.ITradeRepo)
type TradeRepo struct {
    db *gorm.DB
}

func NewTradeRepo(db *gorm.DB) *TradeRepo {
    return &TradeRepo{db: db}
}

func (t TradeRepo) GetTrades(ctx context.Context) ([]*model.Trade, error) {
    var trades []*model.Trade
    err := t.db.Find(&trades).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return trades, nil
}

func (t TradeRepo) GetTradesByUser(ctx context.Context, userEmail string) ([]*model.Trade, error) {
    var trades []*model.Trade
    err := t.db.Where("user_email = ?", userEmail).Find(&trades).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return trades, nil
}

func (t TradeRepo) GetTradesByOrderId(ctx context.Context, orderId string) (*model.Trade, error) {
    var trade *model.Trade
    err := t.db.Where("user_email = ?", orderId).Take(&trade).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return trade, nil
}

func (t TradeRepo) InsertTrade(ctx context.Context, trade *model.Trade) (*model.Trade, error) {
    err := t.db.Create(trade).Error
    if err != nil {
        return nil, err
    }

    return trade, nil
}

// TradeRepo implements the ITradeRepo interface
var _ model.ITradeRepo = &TradeRepo{}
