package vo

/**
 * @author Rancho
 * @date 2022/1/16
 */

type PingResp struct {
}

type TickerPrice struct {
    Symbol string  `json:"symbol"`
    Price  float64 `json:"price,string"`
}

type Ticker24Hour struct {
    Symbol             string  `json:"symbol"`
    PriceChange        float64 `json:"priceChange,string"`
    PriceChangePercent float64 `json:"priceChangePercent,string"`
    WeightedAgvPrice   float64 `json:"weightedAgvPrice,string"`
    HighPrice          float64 `json:"highPrice,string"`
    LowPrice           float64 `json:"lowPrice,string"`
    Count              int64   `json:"count"`
}

type KLine struct {
    OpenTime       int64
    CloseTime      int64
    Open           float64
    Close          float64
    High           float64
    Low            float64
    Volume         float64
    NumberOfTrades int64
}

type TradeResult struct {
    Symbol              string  `json:"symbol"`
    OrderId             int     `json:"orderId"`
    OrderListId         int     `json:"orderListId"`
    ClientOrderId       string  `json:"clientOrderId"`
    TransactTime        int64   `json:"transactTime"`
    Price               float64 `json:"price,string"`
    OrigQty             float64 `json:"origQty,string"`
    ExecutedQty         float64 `json:"executedQty,string"`
    CummulativeQuoteQty float64 `json:"cummulativeQuoteQty"`
    Status              string  `json:"status"`
    TimeInForce         string  `json:"timeInForce"`
    Type                string  `json:"type"`
    Side                string  `json:"side"`
}
