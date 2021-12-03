package dingding

import (
	"context"

    "quants/internal/port.adapter/dependency/http"
)

/**
 * @author Rancho
 * @date 2021/12/3
 */

type DingDing struct {
	IsAtAll bool
	Message string
}

func (d *DingDing) SendDingMessage(ctx context.Context) bool {
	return http.DingDingClient.SendDingDingMessage(d.Message, d.IsAtAll)
}
