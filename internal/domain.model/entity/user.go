package entity

import (
    "time"

    "gorm.io/gorm"
)

/**
 * @author Rancho
 * @date 2022/1/8
 */

type User struct {
    ID        int64                  `gorm:"primarykey" structs:",omitempty,underline"`
    UserName  string                 `json:"user_name" structs:",omitempty,underline"`
    UserEmail string                 `json:"user_email" structs:",omitempty,underline"`
    Asset     *float64               `json:"asset" structs:",omitempty,underline"`
    Profit    *float64               `json:"profit" structs:",omitempty,underline"`
    State     int                    `json:"state" structs:",omitempty,underline"`
    CreatedAt time.Time              `json:"created_at" structs:",omitempty,underline"`
    UpdatedAt time.Time              `json:"updated_at" structs:",omitempty,underline"`
    DeletedAt gorm.DeletedAt         `gorm:"index" structs:",omitempty,underline"`
    ChangeMap map[string]interface{} `json:"-" structs:"-"`
}

func (User) TableName() string {
    return "quant_user"
}
