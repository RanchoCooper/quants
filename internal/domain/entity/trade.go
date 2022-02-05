package entity

import (
    "time"

    "gorm.io/gorm"
)

/**
 * @author Rancho
 * @date 2022/1/8
 */

type Type int8

const (
    TypeInit Type = iota
    TypeBuy
    TypeSell
)

type Trade struct {
    Id         int64                  `json:"id" structs:",omitempty,underline" gorm:"primarykey" `
    UserEmail  string                 `json:"user_email" structs:",omitempty,underline"`
    Symbol     string                 `json:"symbol" structs:",omitempty,underline"`
    OrderId    string                 `json:"order_id" structs:",omitempty,underline"`
    Type       Type                   `json:"type" structs:",omitempty,underline"`
    Price      float64                `json:"price" structs:",omitempty,underline"`
    Quantity   float64                `json:"quantity" structs:",omitempty,underline"`
    IsSimulate bool                   `json:"is_simulate" structs:",omitempty,underline"`
    CreatedAt  time.Time              `json:"created_at" structs:",omitempty,underline"`
    UpdatedAt  time.Time              `json:"updated_at" structs:",omitempty,underline"`
    DeletedAt  gorm.DeletedAt         `json:"deleted_at" structs:",omitempty,underline" gorm:"index"`
    ChangeMap  map[string]interface{} `json:"-" structs:"-" gorm:"-"`
}

func (Trade) TableName() string {
    return "quant_trade"
}
