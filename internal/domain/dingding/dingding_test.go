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
    d := DingDing{
        Message : "【构建系统】",
        IsAtAll : false,
    }
    d.SendDingMessage(context.Background())
}
