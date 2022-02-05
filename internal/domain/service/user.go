package service

import (
    "context"

    "quants/internal/adapter/repository"
    "quants/internal/domain/entity"
    "quants/internal/domain/repo"
    "quants/util/logger"
)

/**
 * @author Rancho
 * @date 2022/2/6
 */

type UserService struct {
    Repository repo.IUserRepo
}

func NewUserService(ctx context.Context) *UserService {
    srv := &UserService{Repository: repository.User}
    logger.Log.Info(ctx, "user service init successfully")
    return srv
}

func (e *UserService) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
    user, err := e.Repository.Create(ctx, nil, user)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (e *UserService) Delete(ctx context.Context, id int) error {
    err := e.Repository.Delete(ctx, nil, id)
    if err != nil {
        return err
    }
    return nil
}

func (e *UserService) Update(ctx context.Context, user *entity.User) error {
    err := e.Repository.Update(ctx, nil, user)
    if err != nil {
        return err
    }
    return nil
}

func (e *UserService) Get(ctx context.Context, id int) (*entity.User, error) {
    user, err := e.Repository.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return user, nil
}
