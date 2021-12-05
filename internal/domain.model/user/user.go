package user

import (
    "gorm.io/gorm"
)

/**
 * @author Rancho
 * @date 2021/12/5
 */

type User struct {
    gorm.Model
    UserName  string `gorm:"column:user_name"`
    UserEmail string `gorm:"column:user_email"`
    State     int    `gorm:"column:state"`
}

func (User) TableName() string {
    return "quant_user"
}
