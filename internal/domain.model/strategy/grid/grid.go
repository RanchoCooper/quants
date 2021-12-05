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

type exchangeType int8

const (
    Buy  exchangeType = 0
    Sell exchangeType = 1
)

type Grid struct {
    ConfigJSONFile string `json:"-"`
    RunBet         struct {
        NextBuyPrice  float64 `json:"next_buy_price"`  // 下次开仓价
        GridSellPrice float64 `json:"grid_sell_price"` // 当前止盈价
        Step          int     `json:"step"`            // 当前仓位
    } `json:"run_bet"`
    Config struct {
        ProfitRatio      int       `json:"profit_ratio"`       // 止盈比率
        DoubleThrowRatio int       `json:"double_throw_ratio"` // 补仓比率
        Cointype         string    `json:"cointype"`           // 交易币种
        Quantity         []float64 `json:"quantity"`           // 交易数量
    } `json:"config"`
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

func (b *Grid) GetQuantity(action exchangeType) float64 {
    var curStep int
    if action == Buy {
        curStep = b.RunBet.Step
    }
    if action == Sell {
        curStep = b.RunBet.Step - 1
    }
    quantity := 0.0
    if curStep < len(b.Config.Quantity) {
        if curStep == 0 {
            quantity = b.Config.Quantity[0]
        } else {
            quantity = b.Config.Quantity[curStep]
        }
    } else {
        // 当前仓位大于设置的仓位，取最后一位
        quantity = b.Config.Quantity[len(b.Config.Quantity)-1]
    }

    return quantity
}

func (b *Grid) ModifyPrice(ctx context.Context, dealPrice float64, step int) bool {
    b.RunBet.NextBuyPrice = dealPrice * cast.ToFloat64(1-b.Config.DoubleThrowRatio/100)
    b.RunBet.GridSellPrice = dealPrice * cast.ToFloat64(1+b.Config.ProfitRatio/100)
    b.RunBet.Step = step

    logger.Log.Infof(ctx, "修改后的补仓价格为%f。修改后的网格价格为%f", b.RunBet.NextBuyPrice, b.RunBet.GridSellPrice)
    return b.WriteToJSON(ctx)
}

func (b *Grid) GetBuyPrice() float64 {
    return b.RunBet.NextBuyPrice
}

func (b *Grid) GetSellPrice() float64 {
    return b.RunBet.GridSellPrice
}

func (b *Grid) GetCoinType() string {
    return b.Config.Cointype
}

func (b *Grid) GetStep() int {
    return b.RunBet.Step
}
