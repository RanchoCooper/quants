package binance

import (
    "context"
    "encoding/json"

    "quants/internal/port.adapter/dependency/http"
    "quants/util/logger"
)

/**
 * @author Rancho
 * @date 2021/12/3
 */

type Trader struct {
}

type TradeInfoVO struct {
    Symbol   string
    Side     string
    Quantity float64
    Price    float64
}

type TradeResp struct {
    Symbol              string `json:"symbol"`
    OrderId             int    `json:"orderId"`
    OrderListId         int    `json:"orderListId"`
    ClientOrderId       string `json:"clientOrderId"`
    TransactTime        int64  `json:"transactTime"`
    Price               string `json:"price"`
    OrigQty             string `json:"origQty"`
    ExecutedQty         string `json:"executedQty"`
    CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
    Status              string `json:"status"`
    TimeInForce         string `json:"timeInForce"`
    Type                string `json:"type"`
    Side                string `json:"side"`
    Fills               []struct {
        Price           string `json:"price"`
        Qty             string `json:"qty"`
        Commission      string `json:"commission"`
        CommissionAsset string `json:"commissionAsset"`
    } `json:"fills"`
    Msg string `json:"msg"`
}

func (tr *TradeResp) String() string {
    b, _ := json.MarshalIndent(tr, "", "    ")
    return string(b)
}

func (u *Trader) Trade(ctx context.Context, vo *TradeInfoVO) (*TradeResp, bool) {
    if vo.Symbol == "" {
        logger.Log.Errorf(ctx, "before trading you should supplement symbol")
        return nil, false
    }
    if vo.Side == "" {
        logger.Log.Errorf(ctx, "before trading you should supplement side(BUY or SELL)")
        return nil, false
    }
    if &vo.Quantity == nil {
        logger.Log.Errorf(ctx, "before trading you should supplement quantity")
        return nil, false
    }
    if &vo.Price == nil {
        logger.Log.Errorf(ctx, "before trading you should supplement price")
        return nil, false
    }

    body := http.BinanceClient.TradeLimit(vo.Symbol, vo.Side, &vo.Quantity, &vo.Price)
    resp := &TradeResp{}
    err := json.Unmarshal(body, resp)
    if err != nil {
        logger.Log.Errorf(ctx, "unmarshal fail when Trade, err: %v", err)
        return nil, false
    }

    return resp, true
}
