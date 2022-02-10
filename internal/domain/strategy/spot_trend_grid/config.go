package spot_trend_grid

import (
    "context"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
    "time"

    "github.com/oleiade/reflections"
    "github.com/spf13/cast"

    "quants/internal/adapter/dependency/http"
    "quants/util"
    "quants/util/logger"
)

/**
 * @author Rancho
 * @date 2022/2/7
 */

const ConfigFileName = "/data.json"

type Config struct {
    StartTime string     `json:"start_time"`
    EndTime   string     `json:"end_time"`
    Interval  string     `json:"interval"`
    CoinList  []string   `json:"coinList"`
    ETHUSDT   CoinConfig `json:"ETHUSDT"`
    BTCUSDT   CoinConfig `json:"BTCUSDT"`
    BNBUSDT   CoinConfig `json:"BNBUSDT"`
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
    jsonFile, err := os.Open(util.GetCurrentPath() + ConfigFileName)
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

func (c *Config) GetStartTime() int64 {
    c.ReadFromFile()
    layout := "2006-01-02 15:04"
    t, err := time.Parse(layout, c.StartTime)
    if err != nil {
        logger.Log.Errorf(context.Background(), "GetStartTime fail when startTime=%s, err: %v", c.StartTime, err)
        return 0
    }
    return t.Unix()
}

func (c *Config) GetEndTime() int64 {
    c.ReadFromFile()
    layout := "2006-01-02 15:04"
    t, _ := time.Parse(layout, c.StartTime)
    t2 := t.Add(cast.ToDuration(strings.ToLower(c.Interval)))
    return t2.Unix()
}

func (c *Config) GetInterval() string {
    return c.Interval
}

func (c *Config) UpdateStartTime() {
    c.ReadFromFile()
    layout := "2006-01-02 15:04"
    t, _ := time.Parse(layout, c.StartTime)
    t2 := t.Add(cast.ToDuration(strings.ToLower(c.Interval)))
    // t2 := t.Add(cast.ToDuration(strings.ToLower(c.Interval)) * 10)
    c.StartTime = t2.Format(layout)

    c.ModifyJSONData()
}

func (c *Config) GetCoinList() []string {
    c.ReadFromFile()
    return c.CoinList
}

func (c *Config) GetBuyPrice(symbol string) float64 {
    c.ReadFromFile()
    item, err := reflections.GetField(c, symbol)
    if err != nil {
        logger.Log.Errorf(context.Background(), "GetBuyPrice fail when symbol=%s, err: %v", symbol, err)
        return 0.0
    }
    return item.(CoinConfig).RunBet.NextBuyPrice
}

func (c *Config) GetSellPrice(symbol string) float64 {
    c.ReadFromFile()
    item, err := reflections.GetField(c, symbol)
    if err != nil {
        logger.Log.Errorf(context.Background(), "GetSellPrice fail when symbol=%s, err: %v", symbol, err)
        return 0.0
    }
    return item.(CoinConfig).RunBet.GridSellPrice
}

func (c *Config) GetCoinType(symbol string) string {
    c.ReadFromFile()
    item, err := reflections.GetField(c, symbol)
    if err != nil {
        logger.Log.Errorf(context.Background(), "GetCoinType fail when symbol=%s, err: %v", symbol, err)
        return ""
    }
    return item.(CoinConfig).Config.Cointype
}

func (c *Config) GetRecordPrice(symbol string) float64 {
    c.ReadFromFile()
    step := c.GetStep(symbol) - 1
    item, err := reflections.GetField(c, symbol)
    if err != nil {
        logger.Log.Errorf(context.Background(), "GetRecordPrice fail when symbol=%s, err: %v", symbol, err)
        return 0.0
    }
    return item.(CoinConfig).RunBet.RecordedPrice[step]
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

    item, err := reflections.GetField(c, symbol)
    if err != nil {
        logger.Log.Errorf(context.Background(), "GetQuantity fail when symbol=%s, err: %v", symbol, err)
        return 0.0
    }
    quantities = item.(CoinConfig).Config.Quantity

    if step < len(quantities) {
        if step < 0 {
            step = 0
        } else if step == 0 {
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
    item, err := reflections.GetField(c, symbol)
    if err != nil {
        logger.Log.Errorf(context.Background(), "GetStep fail when symbol=%s, err: %v", symbol, err)
        return 0
    }
    return item.(CoinConfig).RunBet.Step
}

// GetProfitRatio 补仓比率
func (c *Config) GetProfitRatio(symbol string) float64 {
    c.ReadFromFile()
    item, err := reflections.GetField(c, symbol)
    if err != nil {
        logger.Log.Errorf(context.Background(), "GetProfitRatio fail when symbol=%s, err: %v", symbol, err)
        return 0.0
    }
    return item.(CoinConfig).Config.ProfitRatio
}

// GetDoubleThrowRatio 止盈比率
func (c *Config) GetDoubleThrowRatio(symbol string) float64 {
    c.ReadFromFile()
    item, err := reflections.GetField(c, symbol)
    if err != nil {
        logger.Log.Errorf(context.Background(), "GetDoubleThrowRatio fail when symbol=%s, err: %v", symbol, err)
        return 0.0
    }
    return item.(CoinConfig).Config.DoubleThrowRatio
}

func (c *Config) GetAtr(symbol string) float64 {
    interval := "4h"
    limit := 20
    klines := http.BinanceClinet.GetTickerKLine(context.Background(), symbol, interval, limit, 0, 0)
    percentTotal := 0.0
    for _, kline := range *klines {
        percentTotal += (kline.High - kline.Low) / kline.Close
    }

    return cast.ToFloat64(fmt.Sprintf("%.2f", percentTotal/float64(limit)*100))
}

func (c *Config) ModifyJSONData() {
    file, _ := json.MarshalIndent(c, "", "   ")
    err := ioutil.WriteFile(util.GetCurrentPath()+ConfigFileName, file, 0644)
    if err != nil {
        logger.Log.Errorf(context.Background(), "ModifyJSONData fail, err: %v", err)
    }
}

// SetRatio 修改补仓止盈比率
func (c *Config) SetRatio(symbol string) {
    c.ReadFromFile()
    atr := c.GetAtr(symbol)
    item, err := reflections.GetField(c, symbol)
    if err != nil {
        logger.Log.Errorf(context.Background(), "SetRatio fail when symbol=%s, err: %v", symbol, err)
        return
    }
    newCoinConfig := CoinConfig{}
    newCoinConfig = item.(CoinConfig)
    newCoinConfig.Config.ProfitRatio = atr
    newCoinConfig.Config.DoubleThrowRatio = atr
    err = reflections.SetField(c, symbol, newCoinConfig)
    if err != nil {
        logger.Log.Errorf(context.Background(), "SetRatio fail, err: %v", err)
    }

    c.ModifyJSONData()
}

// SetRecordPrice 记录交易价格
func (c *Config) SetRecordPrice(symbol string, price float64) {
    c.ReadFromFile()
    item, err := reflections.GetField(c, symbol)
    if err != nil {
        logger.Log.Errorf(context.Background(), "SetRecordPrice fail when symbol=%s, err: %v", symbol, err)
        return
    }
    newCoinConfig := CoinConfig{}
    newCoinConfig = item.(CoinConfig)
    newCoinConfig.RunBet.RecordedPrice = append(newCoinConfig.RunBet.RecordedPrice, price)
    err = reflections.SetField(c, symbol, newCoinConfig)
    if err != nil {
        logger.Log.Errorf(context.Background(), "SetRecordPrice fail, err: %v", err)
    }
    c.ModifyJSONData()
}

func (c *Config) RemoveRecordPrice(symbol string) {
    c.ReadFromFile()
    item, err := reflections.GetField(c, symbol)
    if err != nil {
        logger.Log.Errorf(context.Background(), "RemoveRecordPrice fail when symbol=%s, err: %v", symbol, err)
        return
    }
    newCoinConfig := CoinConfig{}
    newCoinConfig = item.(CoinConfig)
    size := len(newCoinConfig.RunBet.RecordedPrice)
    if size == 0 {
        logger.Log.Errorf(context.Background(), "RemoveRecordPrice fail, current %s RecordedPrice size is 0", symbol)
        return
    }
    newCoinConfig.RunBet.RecordedPrice = newCoinConfig.RunBet.RecordedPrice[:size-1]
    err = reflections.SetField(c, symbol, newCoinConfig)
    if err != nil {
        logger.Log.Errorf(context.Background(), "RemoveRecordPrice fail when SetField, err: %v", err)
        return
    }
    c.ModifyJSONData()
}

func (c *Config) ModifyPrice(symbol string, dealPrice float64, step int, marketPrice float64) {
    c.ReadFromFile()
    item, err := reflections.GetField(c, symbol)
    if err != nil {
        logger.Log.Errorf(context.Background(), "ModifyPrice fail when symbol=%s, err: %v", symbol, err)
        return
    }
    newCoinConfig := CoinConfig{}
    newCoinConfig = item.(CoinConfig)
    newCoinConfig.RunBet.NextBuyPrice = dealPrice * (1 - c.ETHUSDT.Config.DoubleThrowRatio/100)
    newCoinConfig.RunBet.GridSellPrice = dealPrice * (1 + c.ETHUSDT.Config.ProfitRatio/100)
    // 如果修改后的价格满足立刻卖出的条件，则再次更改
    if newCoinConfig.RunBet.NextBuyPrice > marketPrice {
        newCoinConfig.RunBet.NextBuyPrice = marketPrice * (1 - c.ETHUSDT.Config.DoubleThrowRatio/100)
    } else if newCoinConfig.RunBet.GridSellPrice < marketPrice {
        newCoinConfig.RunBet.GridSellPrice = marketPrice * (1 + c.ETHUSDT.Config.ProfitRatio/100)
    }
    newCoinConfig.RunBet.Step = step
    err = reflections.SetField(c, symbol, newCoinConfig)
    if err != nil {
        logger.Log.Errorf(context.Background(), "ModifyPrice fail when SetField, err: %v", err)
        return
    }

    c.ModifyJSONData()
    logger.Log.Infof(context.Background(), "币种: %s, 修改后的买入价为: %f，修改后的卖出价为: %f", symbol, newCoinConfig.RunBet.NextBuyPrice, newCoinConfig.RunBet.GridSellPrice)
}
