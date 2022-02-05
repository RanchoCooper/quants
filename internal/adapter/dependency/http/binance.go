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

type binanceAPI struct {
}

func NewBinanceAPI() *binanceAPI {
    return new(binanceAPI)
}

func (b *binanceAPI) Ping() *vo.PingResp {
    resp, err := http.Get(fmt.Sprintf("%s/ping", BinanceAPIV3Url))
    if err != nil {
        logger.Log.Errorf(context.Background(), "BinanceAPIV3 GET /ping err: %v", err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Log.Errorf(context.Background(), "read BinanceAPIV3 response body err: %v", err)
    }

    result := &vo.PingResp{}
    err = json.Unmarshal(body, result)
    if err != nil {
        logger.Log.Errorf(context.Background(), "Unmarshal BinanceAPIV3 response body err: %v", err)
        return nil
    }

    return result
}

func (b *binanceAPI) GetTickerPrice(symbol string) []byte {
    client := &http.Client{}
    req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/ticker/price", BinanceAPIV3Url), nil)

    if symbol != "" {
        query := req.URL.Query()
        query.Add("symbol", symbol)
        req.URL.RawQuery = query.Encode()
    }

    resp, err := client.Do(req)
    if err != nil {
        logger.Log.Errorf(context.Background(), "BinanceAPIV3Url GET /ticker/price err: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Log.Errorf(context.Background(), "read BinanceAPIV3Url body err: %v", err)
    }

    return body
}

func (b *binanceAPI) GetTicker24Hour(symbol string) []byte {
    client := &http.Client{}
    req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/ticker/24hr", BinanceAPIV3Url), nil)

    if symbol != "" {
        query := req.URL.Query()
        query.Add("symbol", symbol)
        req.URL.RawQuery = query.Encode()
    }

    resp, err := client.Do(req)
    if err != nil {
        logger.Log.Errorf(context.Background(), "BinanceAPIV3Url GET /ticker/24hr err: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Log.Errorf(context.Background(), "read BinanceAPIV3Url body err: %v", err)
    }

    return body
}

func (b *binanceAPI) GetTickerKLine(symbol string, interval string, startTime, endTime int64) []byte {
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
        logger.Log.Errorf(context.Background(), "BinanceAPIV3Url GET /klines err: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Log.Errorf(context.Background(), "read BinanceAPIV3Url body err: %v", err)
    }

    return body
}

func (b *binanceAPI) signature(params *url.Values) *url.Values {
    params.Add("timestamp", cast.ToString(time.Now().Unix()))
    params.Add("recvWindow", cast.ToString(5000))

    h := hmac.New(sha256.New, []byte(config.Config.Binance.Secret))
    h.Write([]byte(params.Encode()))
    sha := hex.EncodeToString(h.Sum(nil))
    params.Add("signature", sha)

    return params
}

func (b *binanceAPI) TradeLimit(symbol, side string, quantity, price *float64) []byte {
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
        logger.Log.Errorf(context.Background(), "BinanceAPIV3Url POST /order err: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Log.Errorf(context.Background(), "read BinanceAPIV3Url body err: %v", err)
    }

    return body
}
