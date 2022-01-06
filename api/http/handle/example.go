package handle

import (
    "github.com/gin-gonic/gin"

    "quants/internal/domain.model/service"
    "quants/util/logger"

    "quants/api/http/dto"
    "quants/api/http/errcode"
    "quants/api/http/validator"
)

/**
 * @author Rancho
 * @date 2022/1/5
 */

func CreateExample(ctx *gin.Context) {
    response := NewResponse(ctx)
    body := dto.CreateExampleReq{}
    valid, errs := validator.BindAndValid(ctx, &body, ctx.ShouldBindJSON)
    if !valid {
        logger.Log.Errorf(ctx, "createExample.BindAndList errs: %v", errs)
        errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
        response.ToErrorResponse(errResp)
        return
    }
    example, err := service.Service.ExampleService.Create(ctx, body)
    if err != nil {
        logger.Log.Errorf(ctx, "create example failed.%v", err.Error())
        ctx.Abort()
        return
    }
    response.ToResponse(example)
}
