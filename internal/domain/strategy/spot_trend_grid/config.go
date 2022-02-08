package spot_trend_grid

import (
    "context"
    "encoding/json"
    "io/ioutil"
    "os"

    "github.com/oleiade/reflections"

    "quants/internal/adapter/dependency/http"
    "quants/util/logger"
)

/**
 * @author Rancho
 * @date 2022/2/7
 */

const ConfigFileName = "data.json"

type Config struct {
    CoinList []string   `json:"coinList"`
    ETHUSDT  CoinConfig `json:"ETHUSDT"`
    BTCUSDT  CoinConfig `json:"BTCUSDT"`
    BNBUSDT  CoinConfig `json:"BNBUSDT"`
}

type CoinConfig struct {
    RunBet struct {
        NextBuyPrice  float64   `json:"next_buy_price"`
        GridSellPrice float64   `json:"grid_sell_price"`
        Step          int       `json:"step"`
        RecordedPrice []float64 `json:"recorded_price"`
    } `json:"runBet"`
    Config struct {
        Cointype         string    `json:"cointype"`
        ProfitRatio      float64   `json:"profit_ratio"`
        DoubleThrowRatio float64   `json:"double_throw_ratio"`
        Quantity         []float64 `json:"quantity"`
    } `json:"config"`
}

func (c *Config) ReadFromFile() error {
    jsonFile, err := os.Open(ConfigFileName)
    if err != nil {
        logger.Log.Errorf(context.Background(), "open config file fail, err: %v", err)
        return err
    }
    defer jsonFile.Close()
    data, _ := ioutil.ReadAll(jsonFile)

    err = json.Unmarshal(data, c)
    if err != nil {
        logger.Log.Errorf(context.Background(), "unmarshal config file fail, err: %v", err)
        return err
    }

    return nil
}

func (c *Config) GetCoinList() []string {
    c.ReadFromFile()
    return c.CoinList
}

func (c *Config) GetBuyPrice(symbol string) float64 {
    c.ReadFromFile()
    value, err := reflections.GetField(c, symbol)
    if err != nil {
        return 0.0
    }
    return value.(CoinConfig).RunBet.NextBuyPrice
}

func (c *Config) GetSellPrice(symbol string) float64 {
    c.ReadFromFile()
    value, err := reflections.GetField(c, symbol)
    if err != nil {
        return 0.0
    }
    return value.(CoinConfig).RunBet.GridSellPrice
}

func (c *Config) GetCoinType(symbol string) string {
    c.ReadFromFile()
    value, err := reflections.GetField(c, symbol)
    if err != nil {
        return ""
    }
    return value.(CoinConfig).Config.Cointype
}

func (c *Config) GetRecordPrice(symbol string) float64 {
    c.ReadFromFile()
    step := c.GetStep(symbol) - 1
    value, err := reflections.GetField(c, symbol)
    if err != nil {
        return 0.0
    }
    return value.(CoinConfig).RunBet.RecordedPrice[step]
}

// GetQuantity true 为买入，false为卖出
func (c *Config) GetQuantity(symbol string, exchangeMethod bool) float64 {
    c.ReadFromFile()
    step := c.GetStep(symbol)
    if !exchangeMethod {
        step -= 1
    }
    quantity := 0.0
    var quantities []float64

    value, err := reflections.GetField(c, symbol)
    if err != nil {
        return 0.0
    }
    quantities = value.(CoinConfig).Config.Quantity

    if step < len(quantities) {
        if step == 0 {
            quantity = quantities[0]
        } else {
            quantity = quantities[step]
        }
    } else {
        quantity = quantities[len(quantities)-1]
    }

    return quantity
}

func (c *Config) GetStep(symbol string) int {
    c.ReadFromFile()
    value, err := reflections.GetField(c, symbol)
    if err != nil {
        return 0
    }
    return value.(CoinConfig).RunBet.Step
}

// GetProfitRatio 补仓比率
func (c *Config) GetProfitRatio(symbol string) float64 {
    c.ReadFromFile()
    value, err := reflections.GetField(c, symbol)
    if err != nil {
        return 0.0
    }
    return value.(CoinConfig).Config.ProfitRatio
}

// GetDoubleThrowRatio 止盈比率
func (c *Config) GetDoubleThrowRatio(symbol string) float64 {
    c.ReadFromFile()
    value, err := reflections.GetField(c, symbol)
    if err != nil {
        return 0.0
    }
    return value.(CoinConfig).Config.DoubleThrowRatio
}

func (c *Config) GetAtr(symbol string) float64 {
    interval := "4h"
    limit := 20
    klines := http.BinanceClinet.GetTickerKLine(context.Background(), symbol, interval, limit, 0, 0)
    percentTotal := 0.0
    for _, kline := range *klines {
        percentTotal += (kline.High - kline.Low) / kline.Close
    }

    return percentTotal / float64(limit) * 100
}

// SetRatio 修改补仓止盈比率
func (c *Config) SetRatio(symbol string) {
    c.ReadFromFile()
    atr := c.GetAtr(symbol)
    switch symbol {
    case "ETHUSDT":
        c.ETHUSDT.Config.ProfitRatio = atr
        c.ETHUSDT.Config.DoubleThrowRatio = atr
    case "BTCUSDT":
        c.BTCUSDT.Config.ProfitRatio = atr
        c.BTCUSDT.Config.DoubleThrowRatio = atr
    case "BNBUSDT":
        c.BNBUSDT.Config.ProfitRatio = atr
        c.BNBUSDT.Config.DoubleThrowRatio = atr
    default:
    }

    c.ModifyJSONData()
}

func (c *Config) ModifyJSONData() {
    file, _ := json.MarshalIndent(c, "", "   ")
    _ = ioutil.WriteFile(ConfigFileName, file, 0644)
}

// SetRecordPrice 记录交易价格
func (c *Config) SetRecordPrice(symbol string, price float64) {
    c.ReadFromFile()
    switch symbol {
    case "ETHUSDT":
        c.ETHUSDT.RunBet.RecordedPrice = append(c.ETHUSDT.RunBet.RecordedPrice, price)
    case "BTCUSDT":
        c.BTCUSDT.RunBet.RecordedPrice = append(c.BTCUSDT.RunBet.RecordedPrice, price)
    case "BNBUSDT":
        c.BNBUSDT.RunBet.RecordedPrice = append(c.BNBUSDT.RunBet.RecordedPrice, price)
    default:
    }
}

func (c *Config) RemoveRecordPrice(symbol string) {
    c.ReadFromFile()
    var recordedPrice []float64
    switch symbol {
    case "ETHUSDT":
        recordedPrice = c.ETHUSDT.RunBet.RecordedPrice
    case "BTCUSDT":
        recordedPrice = c.BTCUSDT.RunBet.RecordedPrice
    case "BNBUSDT":
        recordedPrice = c.BNBUSDT.RunBet.RecordedPrice
    default:
        recordedPrice = []float64{0.0}
    }
    size := len(recordedPrice)
    recordedPrice = recordedPrice[:size-1]
    c.ModifyJSONData()
}

func (c *Config) ModifyPrice(symbol string, dealPrice float64, step int, marketPrice float64) {
    c.ReadFromFile()

    switch symbol {
    case "ETHUSDT":
        c.ETHUSDT.RunBet.NextBuyPrice = dealPrice * (1 - c.ETHUSDT.Config.DoubleThrowRatio/100)
        c.ETHUSDT.RunBet.GridSellPrice = dealPrice * (1 + c.ETHUSDT.Config.ProfitRatio/100)
        // 如果修改后的价格满足立刻卖出的条件，则再次更改
        if c.ETHUSDT.RunBet.NextBuyPrice > marketPrice {
            c.ETHUSDT.RunBet.NextBuyPrice = marketPrice * (1 - c.ETHUSDT.Config.DoubleThrowRatio/100)
        } else if c.ETHUSDT.RunBet.GridSellPrice < marketPrice {
            c.ETHUSDT.RunBet.GridSellPrice = marketPrice * (1 + c.ETHUSDT.Config.ProfitRatio/100)
        }
        c.ETHUSDT.RunBet.Step = step
        logger.Log.Infof(context.Background(), "修改后的补仓价格为: %f，修改后的网格价格为: %f", c.ETHUSDT.RunBet.NextBuyPrice, c.ETHUSDT.RunBet.GridSellPrice)
    case "BTCUSDT":
        c.BTCUSDT.RunBet.NextBuyPrice = dealPrice * (1 - c.BTCUSDT.Config.DoubleThrowRatio/100)
        c.BTCUSDT.RunBet.GridSellPrice = dealPrice * (1 + c.BTCUSDT.Config.ProfitRatio/100)
        // 如果修改后的价格满足立刻卖出的条件，则再次更改
        if c.BTCUSDT.RunBet.NextBuyPrice > marketPrice {
            c.BTCUSDT.RunBet.NextBuyPrice = marketPrice * (1 - c.BTCUSDT.Config.DoubleThrowRatio/100)
        } else if c.BTCUSDT.RunBet.GridSellPrice < marketPrice {
            c.BTCUSDT.RunBet.GridSellPrice = marketPrice * (1 + c.BTCUSDT.Config.ProfitRatio/100)
        }
        c.BTCUSDT.RunBet.Step = step
        logger.Log.Infof(context.Background(), "修改后的补仓价格为: %f，修改后的网格价格为: %f", c.BTCUSDT.RunBet.NextBuyPrice, c.BTCUSDT.RunBet.GridSellPrice)
    case "BNBUSDT":
        c.BNBUSDT.RunBet.NextBuyPrice = dealPrice * (1 - c.BNBUSDT.Config.DoubleThrowRatio/100)
        c.BNBUSDT.RunBet.GridSellPrice = dealPrice * (1 + c.BNBUSDT.Config.ProfitRatio/100)
        // 如果修改后的价格满足立刻卖出的条件，则再次更改
        if c.BNBUSDT.RunBet.NextBuyPrice > marketPrice {
            c.BNBUSDT.RunBet.NextBuyPrice = marketPrice * (1 - c.BNBUSDT.Config.DoubleThrowRatio/100)
        } else if c.BNBUSDT.RunBet.GridSellPrice < marketPrice {
            c.BNBUSDT.RunBet.GridSellPrice = marketPrice * (1 + c.BNBUSDT.Config.ProfitRatio/100)
        }
        c.BNBUSDT.RunBet.Step = step
        logger.Log.Infof(context.Background(), "修改后的补仓价格为: %f，修改后的网格价格为: %f", c.BNBUSDT.RunBet.NextBuyPrice, c.BNBUSDT.RunBet.GridSellPrice)
    default:
    }

    c.ModifyJSONData()
}
