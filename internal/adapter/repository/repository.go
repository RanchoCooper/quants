package repository

import (
    "context"

    "quants/config"
    "quants/util/logger"

    "quants/internal/adapter/repository/mysql"
    "quants/internal/adapter/repository/redis"
)

var (
    Clients     = &client{}
    HealthCheck *redis.HealthCheck
    Example     *mysql.Example
    User        *mysql.User
    Trade       *mysql.Trade
)

type client struct {
    MySQL mysql.IMySQL
    Redis redis.IRedis
}

func (c *client) close(ctx context.Context) {
    if c.MySQL != nil {
        c.MySQL.Close(ctx)
    }
    if c.Redis != nil {
        c.Redis.Close(ctx)
    }
}

type Option func(*client)

func WithMySQL() Option {
    return func(c *client) {
        if c.MySQL == nil {
            if config.Config.MySQL != nil {
                c.MySQL = mysql.NewMySQLClient()
            } else {
                panic("init repository with empty MySQL config")
            }
        }
        // inject repository
        if Example == nil {
            Example = mysql.NewExample(Clients.MySQL)
        }
        if User == nil {
            User = mysql.NewUser(Clients.MySQL)
        }
        if Trade == nil {
            Trade = mysql.NewTrade(Clients.MySQL)
        }
    }
}

func WithRedis() Option {
    return func(c *client) {
        if c.Redis == nil {
            if config.Config.Redis != nil {
                c.Redis = redis.NewRedisClient()
            } else {
                panic("init repository with empty Redis config")
            }
        }
        if HealthCheck == nil {
            HealthCheck = redis.NewHealthCheck(Clients.Redis)
        }
    }
}

func Init(opts ...Option) {
    for _, opt := range opts {
        opt(Clients)
    }
    logger.Log.Info(context.Background(), "repository init successfully")
}

func Close(ctx context.Context) {
    Clients.close(ctx)
    logger.Log.Info(ctx, "repository is closed.")
}
