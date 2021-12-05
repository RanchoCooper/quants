package application

import (
    "context"

    "quants/internal/domain.model/user"
    "quants/internal/port.adapter/repository"
    "quants/util/logger"
)

/**
 * @author Rancho
 * @date 2021/12/5
 */

func AddUser(ctx context.Context, userName, userEmail string) bool {
    _, err := repository.MySQL.User.SaveUser(ctx, &user.User{
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