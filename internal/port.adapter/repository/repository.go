package repository

import (
    "context"
    "fmt"

    "quants/config"
    "quants/util/logger"

    "quants/internal/port.adapter/repository/mysql"
    "quants/internal/port.adapter/repository/redis"
)

var (
    Clients = &client{}
    Example *mysql.Example
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

func WithMySQL(ctx context.Context) Option {
    return func(c *client) {
        if c.MySQL == nil {
            if config.Config.MySQL != nil {
                c.MySQL = mysql.NewMySQLClient()
            }
        }
        // inject repository
        if Example == nil {
            Example = mysql.GetExampleInstance(Clients.MySQL)
        }
    }
}

func WithRedis(ctx context.Context) Option {
    return func(c *client) {
        if c.Redis == nil {
            if config.Config.Redis != nil {
                c.Redis = redis.NewRedisClient()
            }
        }
    }
}

func Init(opts ...Option) {
    for _, opt := range opts {
        opt(Clients)
    }
    if Clients.MySQL == nil {
        fmt.Println("opps!!!")
    }
    logger.Log.Info(context.Background(), "repository init successfully")
}

func Close(ctx context.Context) {
    Clients.close(ctx)
    logger.Log.Info(ctx, "repository is closed.")
}
