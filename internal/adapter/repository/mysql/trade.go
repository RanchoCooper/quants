package mysql

import (
    "context"

    "github.com/RanchoCooper/structs"
    "github.com/pkg/errors"
    "gorm.io/gorm"

    "quants/internal/domain/entity"
    "quants/internal/domain/repo"
)

/**
 * @author Rancho
 * @date 2022/2/6
 */

func NewTrade(mysql IMySQL) *Trade {
    return &Trade{IMySQL: mysql}
}

type Trade struct {
    IMySQL
}

var _ repo.ITradeRepo = &Trade{}

func (e *Trade) Create(ctx context.Context, tx *gorm.DB, trade *entity.Trade) (result *entity.Trade, err error) {
    if tx == nil {
        tx = e.GetDB(ctx).Begin()
        defer func() {
            if r := recover(); r != nil {
                tx.Rollback()
                return
            }
            if err != nil {
                tx.Rollback()
                return
            }
            err = errors.WithStack(tx.Commit().Error)
        }()
    }
    err = tx.Table(trade.TableName()).Create(trade).Error
    if err != nil {
        return nil, err
    }

    return trade, nil
}

func (e *Trade) Delete(ctx context.Context, tx *gorm.DB, id int) (err error) {
    if tx == nil {
        tx = e.GetDB(ctx).Begin()
        defer func() {
            if r := recover(); r != nil {
                tx.Rollback()
                return
            }
            if err != nil {
                tx.Rollback()
                return
            }
            err = errors.WithStack(tx.Commit().Error)
        }()
    }
    if id == 0 {
        return errors.New("delete fail. need Id")
    }
    trade := &entity.Trade{}
    err = tx.Table(trade.TableName()).Delete(trade, id).Error
    // hard delete with .Unscoped()
    // err := e.GetDB(ctx).Table(trade.TableName()).Unscoped().Delete(trade, Id).Error
    return err
}

func (e *Trade) Update(ctx context.Context, tx *gorm.DB, trade *entity.Trade) (err error) {
    if tx == nil {
        tx = e.GetDB(ctx).Begin()
        defer func() {
            if r := recover(); r != nil {
                tx.Rollback()
                return
            }
            if err != nil {
                tx.Rollback()
                return
            }
            err = errors.WithStack(tx.Commit().Error)
        }()
    }
    trade.ChangeMap = structs.Map(trade)
    return tx.Table(trade.TableName()).Where("id = ? AND deleted_at IS NULL", trade.Id).Updates(trade.ChangeMap).Error
}

func (e *Trade) Get(ctx context.Context, id int) (*entity.Trade, error) {
    record := &entity.Trade{}
    if id == 0 {
        return nil, errors.New("get fail. need Id")
    }
    err := e.GetDB(ctx).Table(record.TableName()).Find(record, id).Error
    return record, err
}

func (e *Trade) FindByOrderID(ctx context.Context, email string) (*entity.Trade, error) {
    record := &entity.Trade{}
    if email == "" {
        return nil, errors.New("FindByEmail fail. need name")
    }
    err := e.GetDB(ctx).Table(record.TableName()).Where("trade_email = ?", email).Last(record).Error
    return record, err
}
