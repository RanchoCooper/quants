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

func (g Grid) String() string {
    bf, _ := json.MarshalIndent(g, "", "    ")
    return string(bf)
}

func (g *Grid) GetBuyPrice() float64 {
    return g.NextBuyPrice
}

func (g *Grid) GetSellPrice() float64 {
    return g.GridSellPrice
}

func (g *Grid) GetCoinType() string {
    return g.Cointype
}

func (g *Grid) GetStep() int {
    return g.Step
}

func (g Grid) ShouldBuy(curMarketPrice float64) bool {
    return g.GetBuyPrice() > curMarketPrice
}

func (g Grid) ShouldSell(curMarketPrice float64) bool {
    return g.GetSellPrice() < curMarketPrice
}

func (g *Grid) LoadFromJSON(ctx context.Context) {
    content := file.ReadFile(g.ConfigJSONFile)
    if content == nil {
        return
    }
    err := json.Unmarshal(content, g)
    if err != nil {
        logger.Log.Errorf(ctx, "json.Unmarshal fail when LoadFromJSON. err: %v", err)
    }
}

func (g *Grid) WriteToJSON(ctx context.Context) bool {
    filePath := g.ConfigJSONFile
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

    content, err := json.MarshalIndent(g, "", "    ")
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

func (g *Grid) GetQuantity(action ExchangeType) float64 {
    var curStep int
    if action == Buy {
        curStep = g.Step
    }
    if action == Sell {
        curStep = g.Step + len(g.Quantity) - 1
    }
    quantity := 0.0
    if curStep < len(g.Quantity) {
        if curStep == 0 {
            quantity = g.Quantity[0]
        } else {
            quantity = g.Quantity[curStep]
        }
    } else {
        // 当前仓位大于设置的仓位，取最后一位
        quantity = g.Quantity[len(g.Quantity)-1]
    }

    return quantity
}

func (g *Grid) AdjustPrice(ctx context.Context, dealPrice float64, step int) bool {
    g.NextBuyPrice = dealPrice * cast.ToFloat64(1-g.DoubleThrowRatio/100)
    g.GridSellPrice = dealPrice * cast.ToFloat64(1+g.ProfitRatio/100)
    g.Step = step

    logger.Log.Infof(ctx, "修改后的补仓价格为%f。修改后的网格价格为%f", g.NextBuyPrice, g.GridSellPrice)
    return g.WriteToJSON(ctx)
}
