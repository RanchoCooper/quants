package dingding

import (
	"context"

	"cash-cow-quantification/internal/port.adapter/dependency/http"
)

/**
 * @author Rancho
 * @date 2021/12/3
 */

var Dinger DingDing

type DingDing struct {
	IsAtAll bool
	Message string
}

func init() {
	Dinger = DingDing{}
}

func (d *DingDing) SendDingMessage(ctx context.Context) bool {
	http.SendDingDingMessage(d.Message, d.IsAtAll)

	return true
}
