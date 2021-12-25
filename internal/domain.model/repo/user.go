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
    Create(context.Context, *entity.User) (*entity.User, error)
    Delete(context.Context, int64) error
    Update(ctx context.Context, user *entity.User) error
    Get(context.Context, int64) (*entity.User, error)
    FindByEmail(context.Context, string) (*entity.User, error)
}
