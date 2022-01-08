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
    ID        int64    `gorm:"primarykey"`
    UserName  string   `gorm:"column:user_name"`
    UserEmail string   `gorm:"column:user_email"`
    Asset     *float64 `gorm:"column:asset"`
    Profit    *float64 `gorm:"column:profit"`
    State     int      `gorm:"column:state"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
    ChangeMap map[string]interface{}
}

func (User) TableName() string {
    return "quant_user"
}
