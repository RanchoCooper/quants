package mysql

import (
    "context"
    "time"

    "github.com/RanchoCooper/structs"
    "github.com/pkg/errors"
    "gorm.io/gorm"

    "quants/internal/domain/entity"
    "quants/internal/domain/repo"
)

/**
 * @author Rancho
 * @date 2022/1/15
 */

func NewUser(mysql IMySQL) *User {
    return &User{IMySQL: mysql}
}

type User struct {
    IMySQL
}

var _ repo.IUserRepo = &User{}

func (e *User) Create(ctx context.Context, tx *gorm.DB, user *entity.User) (result *entity.User, err error) {
    if tx == nil {
        tx = e.GetDB(ctx).Begin()
        defer func() {
            if r := recover(); r != nil {
                tx.Rollback()
                return
            }
            if err != nil {
                tx.Rollback()
                return
            }
            err = errors.WithStack(tx.Commit().Error)
        }()
    }
    err = tx.Model(user).Create(user).Error
    if err != nil {
        return nil, err
    }

    return user, nil
}

func (e *User) Delete(ctx context.Context, tx *gorm.DB, id int) (err error) {
    if tx == nil {
        tx = e.GetDB(ctx).Begin()
        defer func() {
            if r := recover(); r != nil {
                tx.Rollback()
                return
            }
            if err != nil {
                tx.Rollback()
                return
            }
            err = errors.WithStack(tx.Commit().Error)
        }()
    }
    if id == 0 {
        return errors.New("delete fail. need Id")
    }
    user := &entity.User{}
    err = tx.Model(user).Delete(user, id).Error
    // hard delete with .Unscoped()
    // err := e.GetDB(ctx).Table(user.TableName()).Unscoped().Delete(user, Id).Error
    return err
}

func (e *User) Update(ctx context.Context, tx *gorm.DB, user *entity.User) (err error) {
    if tx == nil {
        tx = e.GetDB(ctx).Begin()
        defer func() {
            if r := recover(); r != nil {
                tx.Rollback()
                return
            }
            if err != nil {
                tx.Rollback()
                return
            }
            err = errors.WithStack(tx.Commit().Error)
        }()
    }
    user.ChangeMap = structs.Map(user)
    user.ChangeMap["updated_at"] = time.Now()
    return tx.Model(user).Where("id = ? AND deleted_at IS NULL", user.Id).Updates(user.ChangeMap).Error
}

func (e *User) Get(ctx context.Context, id int) (*entity.User, error) {
    record := &entity.User{}
    if id == 0 {
        return nil, errors.New("get fail. need Id")
    }
    err := e.GetDB(ctx).Model(record).Find(record, id).Error
    return record, err
}

func (e *User) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
    record := &entity.User{}
    if email == "" {
        return nil, errors.New("FindByEmail fail. need name")
    }
    err := e.GetDB(ctx).Model(record).Where("user_email = ?", email).Last(record).Error
    return record, err
}
