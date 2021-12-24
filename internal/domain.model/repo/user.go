package repo

import (
    "context"

    "quants/internal/domain.model/entity"
)

/**
 * @author Rancho
 * @date 2021/12/23
 */

type IUserRepo interface {
    CreateUser(context.Context, *entity.User) error
    DeleteUser(context.Context, *entity.User) error
    GetUser(context.Context, *entity.User) (*entity.User, error)
    GetUsers(context.Context) ([]*entity.User, error)
}
