package mysql

import (
    "context"
    "errors"

    "gorm.io/gorm"

    "quants/internal/domain.model/user"
)

/**
 * @author Rancho
 * @date 2021/12/5
 */

type IUser interface {
    GetUsers(context.Context) ([]*user.User, error)
    SaveUser(context.Context, *user.User) (*user.User, error)
}

type UserRepo struct {
    db *gorm.DB
}

// UserRepo implements the repository.IUser interface
var _ IUser = &UserRepo{}

func NewUserRepository(db *gorm.DB) *UserRepo {
    return &UserRepo{db: db}
}

func (u *UserRepo) GetUsers(ctx context.Context) ([]*user.User, error) {
    var users []*user.User
    err := u.db.Find(&users).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return users, nil
}

func (u *UserRepo) SaveUser(ctx context.Context, user *user.User) (*user.User, error) {
    err := u.db.Create(user).Error
    if err != nil {
        // if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
        // }
        return nil, err
    }

    return user, nil
}
