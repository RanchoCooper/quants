package application

import (
    "context"

    "quants/global"
    model "quants/internal/domain.model/entity/user"
    "quants/internal/port.adapter/repository"
    "quants/util/logger"
)

/**
 * @author Rancho
 * @date 2021/12/5
 */

func InitEmulateUser(ctx context.Context) {
    u := &model.User{UserEmail: global.SimulateUserEmail}
    u, err := repository.MySQL.User.GetUser(ctx, u)
    if u == nil {
        u = &model.User{
            UserName:  global.SimulateUserName,
            UserEmail: global.SimulateUserEmail,
        }
    }
    if err != nil {
        logger.Log.Errorf(ctx, "InitEmulateUser when GetUser, err: %s", err.Error())
        return
    }
    asset := global.SimulateInitialAsset
    u.Asset = &asset
    err = repository.MySQL.User.CreateUser(ctx, u)
    if err != nil {
        logger.Log.Errorf(ctx, "InitEmulateUser when CreateUser, err: %s", err.Error())
        return
    }
}

func AddUser(ctx context.Context, userName, userEmail string) bool {
    err := repository.MySQL.User.CreateUser(ctx, &model.User{
        UserName:  userName,
        UserEmail: userEmail,
    })
    if err != nil {
        logger.Log.Errorf(ctx, "AddUser fail when CreateUser, err: %v", err)
        return false
    }

    return true
}

func GetUsers(ctx context.Context) []*model.User {
    users, err := repository.MySQL.User.GetUsers(ctx)
    if err != nil {
        logger.Log.Errorf(ctx, "GetUsers fail, err: %v", err)
        return nil
    }

    return users
}
