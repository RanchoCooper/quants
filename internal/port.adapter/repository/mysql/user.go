package mysql

import (
    "context"

    "github.com/pkg/errors"

    "quants/internal/domain.model/entity"
    "quants/internal/domain.model/repo"
)

/**
 * @author Rancho
 * @date 2021/12/25
 */

func GetUserInstance(mysql IMySQL) *User {
    return &User{
        IMySQL: mysql,
    }
}

type User struct {
    IMySQL
}

var _ repo.IUserRepo = &User{}

func (u User) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
    err := u.GetDB(ctx).Create(user).Error
    if err != nil {
        return nil, err
    }

    return user, nil
}

func (u User) Delete(ctx context.Context, ID int64) error {
    if ID == 0 {
        return errors.New("delete fail. need ID")
    }
    err := u.GetDB(ctx).Delete(&entity.User{}, ID).Error
    return err
}

func (u User) Update(ctx context.Context, user *entity.User) error {
    return u.GetDB(ctx).Table(user.TableName()).Updates(user.GetChangeMap()).Error
}

func (u User) Get(ctx context.Context, ID int64) (*entity.User, error) {
    var record *entity.User
    if ID == 0 {
        return nil, errors.New("get fail. need ID")
    }
    err := u.GetDB(ctx).Find(record, ID).Error
    return record, err
}

func (u User) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
    var record *entity.User
    if email == "" {
        return nil, errors.New("FindByEmail fail. need email")
    }
    err := u.GetDB(ctx).Where("user_email = ?", email).Last(record).Error
    return record, err
}
