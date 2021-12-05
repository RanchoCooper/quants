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
    mySQLRepository, err := repository.NewMySQLRepository()
    if err != nil {
        logger.Log.Errorf(ctx, "init MySQLRepository fail, err: %v", err)
        return false
    }

    _, err = mySQLRepository.User.SaveUser(ctx, &user.User{
        UserName:  userName,
        UserEmail: userEmail,
    })
    if err != nil {
        logger.Log.Errorf(ctx, "AddUser fail when SaveUser, err: %v", err)
        return false
    }

    return true
}
