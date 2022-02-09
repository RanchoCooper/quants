package entity

import (
    "time"

    "gorm.io/gorm"
)

/**
 * @author Rancho
 * @date 2022/1/8
 */

const (
    UserStateEnable = 1
    UserStateUnable = 2
)

type User struct {
    Id        int64                  `json:"id" structs:",omitempty,underline" gorm:"primarykey" `
    UserName  string                 `json:"user_name" structs:",omitempty,underline"`
    UserEmail string                 `json:"user_email" structs:",omitempty,underline"`
    Asset     float64                `json:"asset" structs:",omitempty,underline"`
    Profit    float64                `json:"profit" structs:",omitempty,underline"`
    State     int                    `json:"state" structs:",omitempty,underline"`
    CreatedAt time.Time              `json:"created_at" structs:",omitempty,underline"`
    UpdatedAt time.Time              `json:"updated_at" structs:",omitempty,underline"`
    DeletedAt gorm.DeletedAt         `json:"deleted_at" structs:",omitempty,underline" gorm:"index"`
    ChangeMap map[string]interface{} `json:"-" structs:"-" gorm:"-"`
}

func (User) TableName() string {
    return "quant_user"
}
