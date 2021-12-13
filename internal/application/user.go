package application

import (
    "context"

    "quants/global"
    "quants/internal/port.adapter/repository"
    "quants/util/logger"
)

/**
 * @author Rancho
 * @date 2021/12/5
 */

func InitEmulateUser(ctx context.Context) {
    u := &user.User{UserEmail: global.SimulateUserEmail}
    u, err := repository.MySQL.User.GetUser(ctx, u)
    if u == nil {
        u = &user.User{
            UserName:  global.SimulateUserName,
            UserEmail: global.SimulateUserEmail,
        }
    }
    if err != nil {
        logger.Log.Errorf(ctx, "InitEmulateUser when GetUser, err: %s", err.Error())
        return
    }
    u.Asset = global.SimulateInitialAsset
    err = repository.MySQL.User.SaveUser(ctx, u)
    if err != nil {
        logger.Log.Errorf(ctx, "InitEmulateUser when SaveUser, err: %s", err.Error())
        return
    }
}

func AddUser(ctx context.Context, userName, userEmail string) bool {
    err := repository.MySQL.User.SaveUser(ctx, &user.User{
        UserName:  userName,
        UserEmail: userEmail,
    })
    if err != nil {
        logger.Log.Errorf(ctx, "AddUser fail when SaveUser, err: %v", err)
        return false
    }

    return true
}

func GetUsers(ctx context.Context) []*user.User {
    users, err := repository.MySQL.User.GetUsers(ctx)
    if err != nil {
        logger.Log.Errorf(ctx, "GetUsers fail, err: %v", err)
        return nil
    }

    return users
}
