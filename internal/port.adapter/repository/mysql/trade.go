package mysql

import (
    "context"

    "github.com/pkg/errors"

    "quants/internal/domain.model/entity"
    "quants/internal/domain.model/repo"
)

/**
 * @author Rancho
 * @date 2021/12/25
 */

func GetTradeInstance(mysql IMySQL) *Trade {
    return &Trade{
        IMySQL: mysql,
    }
}

type Trade struct {
    IMySQL
}

var _ repo.ITradeRepo = &Trade{}

func (u Trade) Create(ctx context.Context, trade *entity.Trade) (*entity.Trade, error) {
    err := u.GetDB(ctx).Create(trade).Error
    if err != nil {
        return nil, err
    }

    return trade, nil
}

func (u Trade) Delete(ctx context.Context, ID int64) error {
    if ID == 0 {
        return errors.New("delete fail. need ID")
    }
    err := u.GetDB(ctx).Delete(&entity.Trade{}, ID).Error
    return err
}

func (u Trade) Update(ctx context.Context, trade *entity.Trade) error {
    return u.GetDB(ctx).Table(trade.TableName()).Updates(trade.GetChangeMap()).Error
}

func (u Trade) Get(ctx context.Context, ID int64) (*entity.Trade, error) {
    var record *entity.Trade
    if ID == 0 {
        return nil, errors.New("get fail. need ID")
    }
    err := u.GetDB(ctx).Find(record, ID).Error
    return record, err
}

func (u Trade) FindByOrderID(ctx context.Context, orderID string) (*entity.Trade, error) {
    var record *entity.Trade
    if orderID == "" {
        return nil, errors.New("FindByOrderID fail. need orderID")
    }
    err := u.GetDB(ctx).Where("order_id = ?", orderID).Last(record).Error
    return record, err
}
