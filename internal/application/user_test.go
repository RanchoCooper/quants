package application

import (
    "context"
    "testing"

    "github.com/stretchr/testify/assert"

    "quants/util"
)

/**
 * @author Rancho
 * @date 2021/12/5
 */

func TestAddUser(t *testing.T) {
    assert.True(t, AddUser(context.Background(),
        "rancho-test-"+util.RandString(10, false),
        util.RandString(10, false)+"@gmail.com",
    ))
}

func TestGetUsers(t *testing.T) {
    assert.NotEmpty(t, GetUsers(context.Background()))
}
