package mysql

import (
    "regexp"
    "testing"

    "github.com/DATA-DOG/go-sqlmock"
    "github.com/RanchoCooper/structs"
    "github.com/stretchr/testify/assert"

    "quants/internal/domain/entity"
)

/**
 * @author Rancho
 * @date 2022/2/6
 */

func TestUser_Create(t *testing.T) {
    userRepo := NewUser(NewMySQLClient())
    DB, mock := userRepo.MockClient()
    userRepo.SetDB(DB)
    mock.ExpectBegin()
    mock.ExpectExec("INSERT INTO `quant_user`").WillReturnResult(sqlmock.NewResult(1, 1))
    mock.ExpectCommit()
    e := &entity.User{
        UserName:  "random",
        UserEmail: "random",
    }
    user, err := userRepo.Create(ctx, nil, e)
    assert.NoError(t, err)
    assert.NotEmpty(t, user.Id)
    assert.Equal(t, int64(1), user.Id)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}

func TestUser_Delete(t *testing.T) {
    userRepo := NewUser(NewMySQLClient())
    DB, mock := userRepo.MockClient()
    userRepo.SetDB(DB)
    mock.ExpectBegin()
    mock.ExpectExec("UPDATE `quant_user`").WillReturnResult(sqlmock.NewResult(1, 1))
    mock.ExpectCommit()
    err := userRepo.Delete(ctx, nil, 1)
    assert.NoError(t, err)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}

func TestUser_Update(t *testing.T) {
    userRepo := NewUser(NewMySQLClient())
    DB, mock := userRepo.MockClient()
    userRepo.SetDB(DB)
    mock.ExpectBegin()
    mock.ExpectExec("UPDATE `quant_user`").WillReturnResult(sqlmock.NewResult(1, 1))
    mock.ExpectCommit()
    d := &entity.User{
        Id:        1,
        UserName:  "random",
        UserEmail: "random",
    }
    d.ChangeMap = structs.Map(d)
    err := userRepo.Update(ctx, nil, d)
    assert.NoError(t, err)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}

func TestUser_Get(t *testing.T) {
    userRepo := NewUser(NewMySQLClient())
    DB, mock := userRepo.MockClient()
    userRepo.SetDB(DB)
    mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `quant_user` WHERE `quant_user`.`id` = ? AND `quant_user`.`deleted_at` IS NULL")).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test1"))
    user, err := userRepo.Get(ctx, 1)
    assert.NoError(t, err)
    assert.Equal(t, int64(1), user.Id)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}
