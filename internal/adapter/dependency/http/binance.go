package http

import (
    "context"
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "strconv"
    "sync"
    "time"

    "github.com/spf13/cast"

    "quants/config"
    "quants/internal/domain/vo"
    "quants/util/logger"
)

/**
 * @author Rancho
 * @date 2021/12/1
 */

var (
    once          sync.Once
    BinanceClinet *client
)

func init() {
    once.Do(func() {
        BinanceClinet = NewBinanceAPI()
    })
}

type client struct {
}

func NewBinanceAPI() *client {
    return new(client)
}

func (b *client) signature(params *url.Values) *url.Values {
    params.Add("timestamp", cast.ToString(time.Now().Unix()))
    params.Add("recvWindow", cast.ToString(5000))

    h := hmac.New(sha256.New, []byte(config.Config.Binance.Secret))
    h.Write([]byte(params.Encode()))
    sha := hex.EncodeToString(h.Sum(nil))
    params.Add("signature", sha)

    return params
}

func (b *client) Ping(ctx context.Context) *vo.PingResp {
    resp, err := http.Get(fmt.Sprintf("%s/ping", BinanceAPIV3Url))
    if err != nil {
        logger.Log.Errorf(ctx, "Ping err: %v", err)
        return nil
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Log.Errorf(ctx, "Ping error when read response body err: %v", err)
        return nil
    }

    if resp.StatusCode != http.StatusOK {
        logger.Log.Errorf(ctx, "Ping error with status code: %d, errMsg: %s", resp.StatusCode, string(body))
        return nil
    }

    result := &vo.PingResp{}
    err = json.Unmarshal(body, result)
    if err != nil {
        logger.Log.Errorf(ctx, "Ping error when Unmarshal response body err: %v", err)
        return nil
    }

    return result
}

func (b *client) GetTickerPrice(ctx context.Context, symbol string) *vo.TickerPrice {
    client := &http.Client{}
    req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/ticker/price", BinanceAPIV3Url), nil)

    if symbol != "" {
        query := req.URL.Query()
        query.Add("symbol", symbol)
        req.URL.RawQuery = query.Encode()
    }

    resp, err := client.Do(req)
    if err != nil {
        logger.Log.Errorf(ctx, "GetTickerPrice err: %v", err)
        return nil
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Log.Errorf(ctx, "GetTickerPrice error when read body err: %v", err)
        return nil
    }

    if resp.StatusCode != http.StatusOK {
        logger.Log.Errorf(ctx, "Ping error with status code: %d, errMsg: %s", resp.StatusCode, string(body))
        return nil
    }

    result := &vo.TickerPrice{}
    err = json.Unmarshal(body, result)
    if err != nil {
        logger.Log.Errorf(ctx, "GetTickerPrice error when Unmarshal response body err: %v", err)
        return nil
    }

    return result
}

func (b *client) GetTicker24Hour(ctx context.Context, symbol string) *vo.Ticker24Hour {
    client := &http.Client{}
    req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/ticker/24hr", BinanceAPIV3Url), nil)

    if symbol != "" {
        query := req.URL.Query()
        query.Add("symbol", symbol)
        req.URL.RawQuery = query.Encode()
    }

    resp, err := client.Do(req)
    if err != nil {
        logger.Log.Errorf(ctx, "GetTicker24Hour err: %v", err)
        return nil
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Log.Errorf(ctx, "GetTicker24Hour error when read body err: %v", err)
        return nil
    }

    if resp.StatusCode != http.StatusOK {
        logger.Log.Errorf(ctx, "Ping error with status code: %d, errMsg: %s", resp.StatusCode, string(body))
        return nil
    }

    result := &vo.Ticker24Hour{}
    err = json.Unmarshal(body, result)
    if err != nil {
        logger.Log.Errorf(ctx, "GetTicker24Hour error when Unmarshal response body err: %v", err)
        return nil
    }

    return result
}

func (b *client) GetTickerKLine(ctx context.Context, symbol string, interval string, startTime, endTime int64) *[]vo.KLine {
    // TODO
    client := &http.Client{}
    req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/klines", BinanceAPIV3Url), nil)

    query := req.URL.Query()
    query.Add("symbol", symbol)
    query.Add("interval", interval)

    if startTime != 0 && endTime != 0 {
        query.Add("startTime", strconv.FormatInt(startTime, 10))
        query.Add("startTime", strconv.FormatInt(endTime, 10))
    }
    req.URL.RawQuery = query.Encode()

    resp, err := client.Do(req)
    if err != nil {
        logger.Log.Errorf(ctx, "GetTickerKLine err: %v", err)
        return nil
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Log.Errorf(ctx, "GetTickerKLine error when read body err: %v", err)
        return nil
    }

    if resp.StatusCode != http.StatusOK {
        logger.Log.Errorf(ctx, "Ping error with status code: %d, errMsg: %s", resp.StatusCode, string(body))
        return nil
    }

    result := &[]vo.KLine{}
    err = json.Unmarshal(body, result)
    if err != nil {
        logger.Log.Errorf(ctx, "GetTickerKLine error when Unmarshal response body err: %v", err)
        return nil
    }

    return result
}

func (b *client) TradeLimit(ctx context.Context, symbol, side string, quantity, price *float64) *vo.TradeResult {
    // TODO
    client := &http.Client{}
    req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/order", BinanceAPIV3Url), nil)

    req.Header.Set("X-MBX-APIKEY", config.Config.Binance.Key)
    query := req.URL.Query()
    query.Add("symbol", symbol)
    query.Add("side", side)
    query.Add("quantity", fmt.Sprintf("%.8f", *quantity))
    if price == nil {
        query.Add("type", "MARKET")
    } else {
        query.Add("type", "LIMIT")
    }
    query = *b.signature(&query)
    req.URL.RawQuery = query.Encode()

    resp, err := client.Do(req)
    if err != nil {
        logger.Log.Errorf(ctx, "TradeLimit err: %v", err)
        return nil
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Log.Errorf(ctx, "TradeLimit error when read body err: %v", err)
        return nil
    }

    if resp.StatusCode != http.StatusOK {
        logger.Log.Errorf(ctx, "Ping error with status code: %d, errMsg: %s", resp.StatusCode, string(body))
        return nil
    }

    if resp.StatusCode != http.StatusOK {
        logger.Log.Errorf(ctx, "TradeLimit error with status code: %d, errMsg: %s", resp.StatusCode, string(body))
        return nil
    }

    result := &vo.TradeResult{}
    err = json.Unmarshal(body, result)
    if err != nil {
        logger.Log.Errorf(ctx, "TradeLimit error when Unmarshal response body err: %v", err)
        return nil
    }

    return result
}
