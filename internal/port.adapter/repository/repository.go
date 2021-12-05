package repository

import (
    "fmt"

    driver "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/schema"

    "quants/config"
    "quants/internal/port.adapter/repository/mysql"
)

/**
 * @author Rancho
 * @date 2021/12/5
 */

var MySQL *MySQLRepository

type MySQLRepository struct {
    User *mysql.UserRepo
    db   *gorm.DB
}

func init() {
    if MySQL == nil {
        MySQL, err := NewMySQLRepository()
        if err != nil {
            panic("init MySQL fail, err: " + err.Error())
        }
        _ = MySQL
    }
}

func NewMySQLRepository() (*MySQLRepository, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=%s",
        config.Config.MySQL.User,
        config.Config.MySQL.Password,
        config.Config.MySQL.Host,
        config.Config.MySQL.Database,
        config.Config.MySQL.CharSet,
        config.Config.MySQL.ParseTime,
        config.Config.MySQL.TimeZone,
    )
    db, err := gorm.Open(driver.Open(dsn), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true,
        },
    })
    if err != nil {
        return nil, err
    }

    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }
    sqlDB.SetMaxIdleConns(config.Config.MySQL.MaxIdleConns)
    sqlDB.SetMaxOpenConns(config.Config.MySQL.MaxOpenConns)

    MySQL = &MySQLRepository{
        User: mysql.NewUserRepository(db),
        db:   db,
    }

    return MySQL, nil
}
