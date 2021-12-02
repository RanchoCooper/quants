package dingding

import (
    "context"
    "testing"
)

/**
 * @author Rancho
 * @date 2021/12/3
 */

func TestDingDing_SendDingMessage(t *testing.T) {
    Dinger.Message = "【构建系统】"
    Dinger.IsAtAll = false
    Dinger.SendDingMessage(context.Background())
}
