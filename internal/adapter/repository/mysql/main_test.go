package mysql

import (
    "context"
    "testing"

    "quants/internal/domain/entity"
)

/**
 * @author Rancho
 * @date 2022/1/15
 */

var ctx = context.Background()

func TestMain(m *testing.M) {
    db := NewExample(NewMySQLClient()).GetDB(ctx)
    _ = db.AutoMigrate(&entity.Example{})
    m.Run()
}
