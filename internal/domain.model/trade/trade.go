package trade

import (
    "context"

    "gorm.io/gorm"

    "quants/internal/port.adapter/repository"
)

/**
 * @author Rancho
 * @date 2021/12/6
 */

type Type int8

const (
    TypeInit Type = iota
    TypeBuy
    TypeSell
)

type Trade struct {
    gorm.Model
    UserEmail  string  `gorm:"column:user_email"`
    Symbol     string  `gorm:"column:symbol"`
    OrderId    string  `gorm:"column:order_id"`
    Type       Type    `gorm:"column:type"`
    Price      float64 `gorm:"column:price"`
    Quantity   float64 `gorm:"column:quantity"`
    IsSimulate bool    `gorm:"column:is_simulate"`
}

func (Trade) TableName() string {
    return "quant_trade"
}

func (t *Trade) InsertTrade(ctx context.Context) (*Trade, error) {
    return repository.MySQL.Trade.InsertTrade(ctx, t)
}
