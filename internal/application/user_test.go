package application

import (
    "context"
    "testing"

    "github.com/stretchr/testify/assert"

    "quants/config"
    "quants/util"
)

/**
 * @author Rancho
 * @date 2021/12/5
 */

func TestAddUser(t *testing.T) {
    if config.Config.Env == string(config.EnvTesting) {
        return
    }
    assert.True(t, AddUser(context.Background(),
        "rancho-test-"+util.RandString(10, false),
        util.RandString(10, false)+"@gmail.com",
    ))
}

func TestGetUsers(t *testing.T) {
    if config.Config.Env == string(config.EnvTesting) {
        assert.Empty(t, GetUsers(context.Background()))
    } else {
        assert.NotEmpty(t, GetUsers(context.Background()))
    }
}
