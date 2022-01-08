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
    ID         int64                  `gorm:"primarykey" structs:",omitempty,underline"`
    UserEmail  string                 `json:"user_email" structs:",omitempty,underline"`
    Symbol     string                 `json:"symbol" structs:",omitempty,underline"`
    OrderId    string                 `json:"order_id" structs:",omitempty,underline"`
    Type       Type                   `json:"type" structs:",omitempty,underline"`
    Price      float64                `json:"price" structs:",omitempty,underline"`
    Quantity   float64                `json:"quantity" structs:",omitempty,underline"`
    IsSimulate bool                   `json:"is_simulate" structs:",omitempty,underline"`
    CreatedAt  time.Time              `json:"created_at" structs:",omitempty,underline"`
    UpdatedAt  time.Time              `json:"updated_at" structs:",omitempty,underline"`
    DeletedAt  gorm.DeletedAt         `gorm:"index" structs:",omitempty,underline"`
    ChangeMap  map[string]interface{} `json:"-" structs:"-"`
}

func (Trade) TableName() string {
    return "quant_trade"
}
