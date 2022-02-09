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
 * @date 2022/2/9
 */

const (
    SimulatorUserName  = "simulator"
    SimulatorUserEmail = "rancho@simulator.com"
)

type SimulatorService struct {
    UserRepository  repo.IUserRepo
    TradeRepository repo.ITradeRepo
}

func NewSimulatorService(ctx context.Context) *SimulatorService {
    srv := &SimulatorService{UserRepository: repository.User, TradeRepository: repository.Trade}
    logger.Log.Info(ctx, "simulator service init successfully")
    return srv
}

func (ss *SimulatorService) FindOrCreateSimulateUser(ctx context.Context) *entity.User {
    user, err := ss.UserRepository.FindByEmail(ctx, SimulatorUserEmail)
    if err != nil {
        logger.Log.Errorf(ctx, "find simulate user fail when simulate, err: %v", err)
        return nil
    }
    if user == nil {
        user, err = ss.UserRepository.Create(ctx, nil, &entity.User{
            UserName:  SimulatorUserName,
            UserEmail: SimulatorUserEmail,
            Asset:     10000,
            Profit:    0,
            State:     entity.UserStateEnable,
        })
        if err != nil {
            logger.Log.Errorf(ctx, "create simulate user fail when simulate, err: %v", err)
            return nil
        }
    }
    return user
}

func (ss *SimulatorService) Buy(ctx context.Context, symbol string, price, quantity float64) {
    user := ss.FindOrCreateSimulateUser(ctx)
    if user == nil {
        return
    }

    trade := &entity.Trade{
        UserEmail:  user.UserEmail,
        Symbol:     symbol,
        Type:       entity.TradeTypeBuy,
        Price:      price,
        Quantity:   quantity,
        IsSimulate: true,
    }
    trade, err := ss.TradeRepository.Create(ctx, nil, trade)
    if err != nil {
        logger.Log.Errorf(ctx, "create user trade fail when simulate, trade: %v, err: %v", trade, err)
        return
    }

    // update user asset
    user.Asset -= price * quantity
    err = ss.UserRepository.Update(ctx, nil, user)
    if err != nil {
        logger.Log.Errorf(ctx, "update user fail when simulate, err: %v", err)
        return
    }
}

func (ss *SimulatorService) Sell(ctx context.Context, symbol string, price, quantity float64) {
    user := ss.FindOrCreateSimulateUser(ctx)
    if user == nil {
        return
    }

    trade := &entity.Trade{
        UserEmail:  user.UserEmail,
        Symbol:     symbol,
        Type:       entity.TradeTypeSell,
        Price:      price,
        Quantity:   quantity,
        IsSimulate: true,
    }
    trade, err := ss.TradeRepository.Create(ctx, nil, trade)
    if err != nil {
        logger.Log.Errorf(ctx, "create user trade fail when simulate, trade: %v, err: %v", trade, err)
        return
    }

    // update user asset
    user.Asset += price * quantity
    err = ss.UserRepository.Update(ctx, nil, user)
    if err != nil {
        logger.Log.Errorf(ctx, "update user fail when simulate, err: %v", err)
        return
    }
}
