package repository

import (
    "fmt"

    "github.com/DATA-DOG/go-sqlmock"
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
    User  *mysql.UserRepo
    Trade *mysql.TradeRepo
    db    *gorm.DB
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

func mockMySQL() *gorm.DB {
    sqlDB, mock, err := sqlmock.New()
    if err != nil {
        panic("mock MySQL fail, err: " + err.Error())
    }
    dialector := driver.New(driver.Config{
        Conn:       sqlDB,
        DriverName: "mysql",
    })
    // a SELECT VERSION() query will be run when gorm opens the database, so we need to expect that here
    columns := []string{"version"}
    mock.ExpectQuery("SELECT VERSION()").WithArgs().WillReturnRows(
        mock.NewRows(columns).FromCSVString("1"),
    )
    mock.ExpectExec("INSERT INTO `quant_user`").WillReturnResult(sqlmock.NewResult(1, 1))
    mock.ExpectQuery("SELECT (.+) FROM `quant_user`").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test1"))
    db, err := gorm.Open(dialector, &gorm.Config{})

    return db
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

    if config.Config.Env == string(config.EnvTesting) {
        db := mockMySQL()
        MySQL = &MySQLRepository{
            User:  mysql.NewUserRepo(db),
            Trade: mysql.NewTradeRepo(db),
            db:    db,
        }

        return MySQL, nil
    }

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
        User: mysql.NewUserRepo(db),
        db:   db,
    }

    return MySQL, nil
}
