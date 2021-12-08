package grid

import (
    "context"
    "encoding/json"
    "io"
    "os"

    "github.com/spf13/cast"

    "quants/util/file"
    "quants/util/logger"
)

/**
 * @author Rancho
 * @date 2021/12/5
 */

type ExchangeType int8

const (
    Buy  ExchangeType = 0
    Sell ExchangeType = 1
)

type Grid struct {
    ConfigJSONFile   string    `json:"-"`
    NextBuyPrice     float64   `json:"next_buy_price"`     // 下次开仓价
    GridSellPrice    float64   `json:"grid_sell_price"`    // 当前止盈价
    Step             int       `json:"step"`               // 当前仓位
    ProfitRatio      int       `json:"profit_ratio"`       // 止盈比率
    DoubleThrowRatio int       `json:"double_throw_ratio"` // 补仓比率
    Cointype         string    `json:"cointype"`           // 交易币种
    Quantity         []float64 `json:"quantity"`           // 交易数量
}

func (b Grid) String() string {
    bf, _ := json.MarshalIndent(b, "", "    ")
    return string(bf)
}

func (b *Grid) LoadFromJSON(ctx context.Context) {
    content := file.ReadFile(b.ConfigJSONFile)
    if content == nil {
        return
    }
    err := json.Unmarshal(content, b)
    if err != nil {
        logger.Log.Errorf(ctx, "json.Unmarshal fail when LoadFromJSON. err: %v", err)
    }
}

func (b *Grid) WriteToJSON(ctx context.Context) bool {
    filePath := b.ConfigJSONFile
    if _, err := os.Stat(filePath); err != nil {
        if os.IsNotExist(err) {
            _, err2 := os.Create(filePath)
            if err2 != nil {
                logger.Log.Errorf(ctx, "WriteToJSON fail when Create f, err: %v", err)
                return false
            }
            return true
        } else {
            logger.Log.Errorf(ctx, "WriteToJSON fail, err: %v", err)
            return false
        }
    }
    f, err := os.OpenFile(filePath, os.O_WRONLY, 0666)
    if err != nil {
        logger.Log.Errorf(ctx, "WriteToJSON fail when os.OpenFile, err: %v", err)
        return false
    }
    defer f.Close()

    content, err := json.MarshalIndent(b, "", "    ")
    if err != nil {
        logger.Log.Errorf(ctx, "WriteToJSON fail when json.MarshalIndent, err: %v", err)
        return false
    }

    _, err = io.WriteString(f, string(content))
    if err != nil {
        logger.Log.Errorf(ctx, "WriteToJSON fail when io.WriteString, err: %v", err)
        return false
    }

    return true
}

func (b *Grid) GetQuantity(action ExchangeType) float64 {
    var curStep int
    if action == Buy {
        curStep = b.Step
    }
    if action == Sell {
        curStep = b.Step + len(b.Quantity) - 1
    }
    quantity := 0.0
    if curStep < len(b.Quantity) {
        if curStep == 0 {
            quantity = b.Quantity[0]
        } else {
            quantity = b.Quantity[curStep]
        }
    } else {
        // 当前仓位大于设置的仓位，取最后一位
        quantity = b.Quantity[len(b.Quantity)-1]
    }

    return quantity
}

func (b *Grid) AdjustPrice(ctx context.Context, dealPrice float64, step int) bool {
    b.NextBuyPrice = dealPrice * cast.ToFloat64(1-b.DoubleThrowRatio/100)
    b.GridSellPrice = dealPrice * cast.ToFloat64(1+b.ProfitRatio/100)
    b.Step = step

    logger.Log.Infof(ctx, "修改后的补仓价格为%f。修改后的网格价格为%f", b.NextBuyPrice, b.GridSellPrice)
    return b.WriteToJSON(ctx)
}

func (b *Grid) GetBuyPrice() float64 {
    return b.NextBuyPrice
}

func (b *Grid) GetSellPrice() float64 {
    return b.GridSellPrice
}

func (b *Grid) GetCoinType() string {
    return b.Cointype
}

func (b *Grid) GetStep() int {
    return b.Step
}
