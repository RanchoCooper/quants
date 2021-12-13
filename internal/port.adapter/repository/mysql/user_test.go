package mysql

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "quants/global"
    "quants/internal/domain/model/user"
)

/**
 * @author Rancho
 * @date 2021/12/13
 */

func TestUserRepo_GetUser(t *testing.T) {
    user := &model.User{
        UserName:  global.SimulateUserName,
        UserEmail: global.SimulateUserEmail,
    }
    u, err := MySQL.User.GetUser(ctx, user)
    assert.Nil(t, err)
    assert.Equal(t, 2, u.ID)
    assert.Equal(t, user.UserName, u.UserName)
    assert.Equal(t, user.UserEmail, u.UserEmail)

    user = &model.User{}
    u, err = MySQL.User.GetUser(ctx, user) // return first user record
    assert.Nil(t, err)
    assert.Equal(t, 1, u.ID)

    user = &model.User{
        UserName: "not exists",
    }
    u, err = MySQL.User.GetUser(ctx, user) // return first user record
    assert.NotNil(t, err)
    assert.Nil(t, u)
}
