package http

import (
    "github.com/gin-gonic/gin"

    "quants/config"
)

func NewServerRoute() *gin.Engine {
    if config.Config.App.Debug {
        gin.SetMode(gin.DebugMode)
    } else {
        gin.SetMode(gin.ReleaseMode)
    }

    router := gin.Default()
    example := router.Group("/example")
    {
        example.POST("", CreateExample)
        example.DELETE("", DeleteExample)
        example.PUT("", UpdateExample)
        example.GET("", GetExample)
    }

    return router
}
