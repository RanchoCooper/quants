package mysql

import (
    "context"
    "errors"

    "gorm.io/gorm"

    "quants/global"
    "quants/internal/domain/model/user"
)

/**
 * @author Rancho
 * @date 2021/12/5
 */

type UserRepo struct {
    db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
    return &UserRepo{db: db}
}

func (u *UserRepo) GetUser(ctx context.Context, user *model.User) (*model.User, error) {
    u.db = u.db.Session(&gorm.Session{NewDB: true})
    var result *model.User
    if len(user.UserName) != 0 {
        u.db = u.db.Where(user, "user_name")
    }
    if len(user.UserEmail) != 0 {
        u.db = u.db.Where(user, "user_email")
    }
    err := u.db.First(&result).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, global.ErrNotFound
    }
    if err != nil {
        return nil, err
    }
    return result, nil
}

func (u *UserRepo) GetUsers(ctx context.Context) ([]*model.User, error) {
    var users []*model.User
    err := u.db.Find(&users).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    if err != nil {
        return nil, global.ErrNotFound
    }
    return users, nil
}

func (u *UserRepo) SaveUser(ctx context.Context, user *model.User) error {
    err := u.db.Create(user).Error
    if err != nil {
        // if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
        // }
        return err
    }

    return nil
}

// UserRepo implements the IUserRepo interface
var _ model.IUserRepo = &UserRepo{}
