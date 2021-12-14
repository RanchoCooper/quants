package mysql

import (
    "regexp"
    "testing"

    "github.com/DATA-DOG/go-sqlmock"
    "github.com/stretchr/testify/require"
    "github.com/stretchr/testify/suite"
    "gorm.io/gorm"

    "quants/global"
    "quants/internal/domain/model/user"
    "quants/util"
)

/**
 * @author Rancho
 * @date 2021/12/13
 */

type Suite struct {
    suite.Suite
    DB   *gorm.DB
    mock sqlmock.Sqlmock
}

func (s *Suite) SetupSuite() {
    s.DB, s.mock = mockMySQL()
}

func (s *Suite) AfterTest(_, _ string) {
    require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestUserRepo_GetUsers(t *testing.T) {
    suite.Run(t, new(Suite))
}

func (s *Suite) TestUserRepo_GetUsers() {
    s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `quant_user` WHERE `quant_user`.`deleted_at` IS NULL")).WillReturnRows(sqlmock.NewRows([]string{"id", "user_name"}).AddRow(1, "test1"))
    us, err := MySQL.User.GetUsers(ctx)
    s.NoError(err)
    s.NotEmpty(us)
}

func (s *Suite) TestUserRepo_SaveUser() {
    s.mock.ExpectBegin()
    s.mock.ExpectExec("INSERT INTO `quant_user`").WillReturnResult(sqlmock.NewResult(1, 1))
    s.mock.ExpectCommit()
    asset := 0.0
    profit := 0.0
    user := &model.User{
        UserName:  util.RandString(10, false),
        UserEmail: util.RandString(10, false),
        Asset:     &asset,
        Profit:    &profit,
    }
    err := MySQL.User.CreateUser(ctx, user)
    s.NoError(err)
    s.NotEmpty(user.ID)
}

func (s *Suite) TestUserRepo_UpdateUser() {
    s.Run("update empty object", func() {
        user := &model.User{}
        user, err := MySQL.User.UpdateUser(ctx, user)
        s.Error(err)
    })

    s.Run("normal update", func() {
        s.mock.ExpectBegin()
        s.mock.ExpectExec("UPDATE `quant_user` SET").WillReturnResult(sqlmock.NewResult(1, 1))
        s.mock.ExpectCommit()
        user := &model.User{
            ID:    2,
            State: 2,
        }
        user, err := MySQL.User.UpdateUser(ctx, user)
        s.NoError(err)
        s.Equal(2, user.State)
    })

    s.Run("zero value update", func() {
        s.mock.ExpectBegin()
        s.mock.ExpectExec("UPDATE `quant_user` SET").WillReturnResult(sqlmock.NewResult(1, 1))
        s.mock.ExpectCommit()
        user := &model.User{
            ID:    2,
            State: 0,
        }
        user, err := MySQL.User.UpdateUser(ctx, user)
        s.NoError(err)
        s.Equal(0, user.State)
    })
}

func (s *Suite) TestUserRepo_GetUser() {
    s.Run("normal GetUser", func() {
        s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `quant_user` WHERE `quant_user`.`user_name` = ? AND `quant_user`.`user_email` = ? AND `quant_user`.`deleted_at` IS NULL ORDER BY `quant_user`.`id` LIMIT 1")).WillReturnRows(sqlmock.NewRows([]string{"id", "user_name", "user_email"}).AddRow(2, "rancho-simulate", "rancho@simulate.com"))
        user := &model.User{
            UserName:  global.SimulateUserName,
            UserEmail: global.SimulateUserEmail,
        }
        u, err := MySQL.User.GetUser(ctx, user)
        s.NoError(err)
        s.Equal(2, u.ID)
        s.Equal(user.UserName, u.UserName)
        s.Equal(user.UserEmail, u.UserEmail)
    })

    s.Run("get first record", func() {
        s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `quant_user` WHERE `quant_user`.`deleted_at` IS NULL ORDER BY `quant_user`.`id` LIMIT 1")).WillReturnRows(sqlmock.NewRows([]string{"id", "user_name", "user_email"}).AddRow(1, "rancho-simulate", "rancho@simulate.com"))
        user := &model.User{}
        u, err := MySQL.User.GetUser(ctx, user) // return first user record
        s.NoError(err)
        s.Equal(1, u.ID)
    })

    s.Run("not exists, get first user record", func() {
        s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `quant_user` WHERE `quant_user`.`user_name` = ? AND `quant_user`.`deleted_at` IS NULL ORDER BY `quant_user`.`id` LIMIT 1")).WillReturnRows(sqlmock.NewRows([]string{"id", "user_name", "user_email"}).AddRow(1, "rancho-simulate", "rancho@simulate.com"))
        user := &model.User{
            UserName: "not exists",
        }
        u, err := MySQL.User.GetUser(ctx, user) // return first user record
        s.NoError(err)
        s.Equal(1, u.ID)
    })
}
