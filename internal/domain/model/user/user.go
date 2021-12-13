package model

import (
    "context"
    "time"

    "gorm.io/gorm"
)

/**
 * @author Rancho
 * @date 2021/12/5
 */

type IUserRepo interface {
    GetUser(context.Context, *User) (*User, error)
    GetUsers(context.Context) ([]*User, error)
    SaveUser(context.Context, *User) error
}

type User struct {
    ID        int     `gorm:"primarykey"`
    UserName  string  `gorm:"column:user_name"`
    UserEmail string  `gorm:"column:user_email"`
    Asset     float64 `gorm:"column:asset"`
    Profit    float64 `gorm:"column:profit"`
    State     int     `gorm:"column:state"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string {
    return "quant_user"
}
