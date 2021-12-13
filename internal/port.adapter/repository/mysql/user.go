package mysql

import (
    "context"
    "errors"

    "github.com/fatih/structs"
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

func (u *UserRepo) CreateUser(ctx context.Context, user *model.User) error {
    u.db = u.db.Session(&gorm.Session{
        NewDB:   true,
        Context: ctx,
    })
    err := u.db.Create(user).Error
    if err != nil {
        // if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
        // }
        return err
    }

    return nil
}

func (u *UserRepo) DeleteUser(ctx context.Context, user *model.User) error {
    u.db = u.db.Session(&gorm.Session{
        NewDB:   true,
        Context: ctx,
    })
    if user.ID != 0 {
        return u.db.Delete(user).Error
    }
    if user.UserName != "" {
        return u.db.Where("user_name = ?", user.UserName).Delete(&model.User{}).Error
    }
    if user.UserEmail != "" {
        return u.db.Where("user_email = ?", user.UserEmail).Delete(&model.User{}).Error
    }
    return nil
}

func (u *UserRepo) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
    u.db = u.db.Session(&gorm.Session{
        NewDB:   true,
        Context: ctx,
    })
    if structs.IsZero(user) {
        return nil, errors.New("can not delete with empty model")
    }
    if user.ID == 0 {
        return nil, errors.New("update fail, need ID")
    }
    m := make(map[string]interface{})
    if user.UserEmail != "" {
        m["user_email"] = user.UserEmail
    }
    if user.UserName != "" {
        m["user_name"] = user.UserName
    }
    if user.Asset != nil {
        m["assert"] = user.Asset
    }
    if user.Profit != nil {
        m["profit"] = user.Profit
    }
    if user.State != 0 {
        m["state"] = user.State
    }
    err := u.db.Model(user).Updates(m).Error
    return user, err
}

func (u *UserRepo) GetUser(ctx context.Context, user *model.User) (*model.User, error) {
    u.db = u.db.Session(&gorm.Session{
        NewDB:   true,
        Context: ctx,
    })
    var result *model.User
    if user.UserName != "" {
        u.db = u.db.Where(user, "user_name")
    }
    if user.UserEmail != "" {
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
    u.db = u.db.Session(&gorm.Session{
        NewDB:   true,
        Context: ctx,
    })
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

// UserRepo implements the IUserRepo interface
var _ model.IUserRepo = &UserRepo{}
