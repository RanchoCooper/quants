package mysql

import (
    "context"
    "testing"
)

/**
 * @author Rancho
 * @date 2021/12/13
 */

var ctx = context.Background()

func TestMain(m *testing.M) {
    _, err := NewMySQLRepository()
    if err != nil {
        panic("init testing MySQL fail, err: " + err.Error())
    }
}
