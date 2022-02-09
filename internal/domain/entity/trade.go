package entity

import (
    "time"

    "github.com/shopspring/decimal"
    "gorm.io/gorm"
)

/**
 * @author Rancho
 * @date 2022/1/8
 */

type Type int8

const (
    TradeTypeBuy Type = iota + 1
    TradeTypeSell
)

const (
    TradeSideBuy  = "BUY"
    TradeSideSell = "SELL"
)

type Trade struct {
    Id         int64                  `json:"id" structs:",omitempty,underline" gorm:"primarykey" `
    UserEmail  string                 `json:"user_email" structs:",omitempty,underline"`
    Symbol     string                 `json:"symbol" structs:",omitempty,underline"`
    OrderId    string                 `json:"order_id" structs:",omitempty,underline"`
    Type       Type                   `json:"type" structs:",omitempty,underline"`
    Price      decimal.Decimal        `json:"price" structs:",omitempty,underline" gorm:"type:numeric(32,6)"`
    Quantity   decimal.Decimal        `json:"quantity" structs:",omitempty,underline" gorm:"type:numeric(32,6)"`
    IsSimulate bool                   `json:"is_simulate" structs:",omitempty,underline"`
    CreatedAt  time.Time              `json:"created_at" structs:",omitempty,underline"`
    UpdatedAt  time.Time              `json:"updated_at" structs:",omitempty,underline"`
    DeletedAt  gorm.DeletedAt         `json:"deleted_at" structs:",omitempty,underline" gorm:"index"`
    ChangeMap  map[string]interface{} `json:"-" structs:"-" gorm:"-"`
}

func (Trade) TableName() string {
    return "quant_trade"
}
