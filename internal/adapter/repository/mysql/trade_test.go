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

func TestTrade_Create(t *testing.T) {
    tradeRepo := NewTrade(NewMySQLClient())
    DB, mock := tradeRepo.MockClient()
    tradeRepo.SetDB(DB)
    mock.ExpectBegin()
    mock.ExpectExec("INSERT INTO `quant_trade`").WillReturnResult(sqlmock.NewResult(1, 1))
    mock.ExpectCommit()
    e := &entity.Trade{
        UserEmail:  "random",
        Symbol:     "BTC",
        OrderId:    "order-id",
        Type:       1,
        Price:      1.0,
        Quantity:   10,
        IsSimulate: true,
    }
    trade, err := tradeRepo.Create(ctx, nil, e)
    assert.NoError(t, err)
    assert.NotEmpty(t, trade.Id)
    assert.Equal(t, int64(1), trade.Id)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}

func TestTrade_Delete(t *testing.T) {
    tradeRepo := NewTrade(NewMySQLClient())
    DB, mock := tradeRepo.MockClient()
    tradeRepo.SetDB(DB)
    mock.ExpectBegin()
    mock.ExpectExec("UPDATE `quant_trade`").WillReturnResult(sqlmock.NewResult(1, 1))
    mock.ExpectCommit()
    err := tradeRepo.Delete(ctx, nil, 1)
    assert.NoError(t, err)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}

func TestTrade_Update(t *testing.T) {
    tradeRepo := NewTrade(NewMySQLClient())
    DB, mock := tradeRepo.MockClient()
    tradeRepo.SetDB(DB)
    mock.ExpectBegin()
    mock.ExpectExec("UPDATE `quant_trade`").WillReturnResult(sqlmock.NewResult(1, 1))
    mock.ExpectCommit()
    d := &entity.Trade{
        Id:        1,
        UserEmail: "random",
        Symbol:    "BTC",
        OrderId:   "order-id",
    }
    d.ChangeMap = structs.Map(d)
    err := tradeRepo.Update(ctx, nil, d)
    assert.NoError(t, err)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}

func TestTrade_Get(t *testing.T) {
    tradeRepo := NewTrade(NewMySQLClient())
    DB, mock := tradeRepo.MockClient()
    tradeRepo.SetDB(DB)
    mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `quant_trade` WHERE `quant_trade`.`id` = ? AND `quant_trade`.`deleted_at` IS NULL")).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test1"))
    trade, err := tradeRepo.Get(ctx, 1)
    assert.NoError(t, err)
    assert.Equal(t, int64(1), trade.Id)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}
