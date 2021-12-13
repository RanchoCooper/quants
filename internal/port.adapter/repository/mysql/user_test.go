package mysql

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "quants/global"
    "quants/internal/domain/model/user"
    "quants/util"
)

/**
 * @author Rancho
 * @date 2021/12/13
 */

func TestUserRepo_GetUsers(t *testing.T) {
    us, err := MySQL.User.GetUsers(ctx)
    assert.Nil(t, err)
    assert.NotEmpty(t, us)
}

func TestUserRepo_SaveUser(t *testing.T) {
    user := &model.User{
        UserName:  util.RandString(10, false),
        UserEmail: util.RandString(10, false),
    }
    err := MySQL.User.CreateUser(ctx, user)
    assert.Nil(t, err)
    assert.NotEmpty(t, user.ID)
}

func TestUserRepo_UpdateUser(t *testing.T) {
    t.Run("update empty object", func(t *testing.T) {
        user := &model.User{}
        user, err := MySQL.User.UpdateUser(ctx, user)
        assert.NotNil(t, err)
    })

    t.Run("normal update", func(t *testing.T) {
        user := &model.User{
            ID:    2,
            State: 2,
        }
        user, err := MySQL.User.UpdateUser(ctx, user)
        assert.Nil(t, err)
        assert.Equal(t, 2, user.State)
    })

    t.Run("zero value update", func(t *testing.T) {
        user := &model.User{
            ID:    2,
            State: 0,
        }
        user, err := MySQL.User.UpdateUser(ctx, user)
        assert.Nil(t, err)
        assert.Equal(t, 0, user.State)
    })
}

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
