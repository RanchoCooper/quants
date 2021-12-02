package http

import (
    "context"
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "strconv"
    "time"

    "cash-cow-quantification/config"
    "cash-cow-quantification/util/logger"

    "github.com/spf13/cast"
)

/**
 * @author Rancho
 * @date 2021/12/1
 */

func Ping() string {
    resp, err := http.Get(fmt.Sprintf("%s/ping", BinanceAPIV3))
    if err != nil {
        logger.Log.Errorf(context.Background(), "BinanceAPIV3 GET /ping err: %v", err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Log.Errorf(context.Background(), "read BinanceAPIV3 body err: %v", err)
    }

    return string(body)
}

func GetTickerPrice(symbol string) []byte {
    client := &http.Client{}
    req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/ticker/price", BinanceAPIV3), nil)

    if symbol != "" {
        query := req.URL.Query()
        query.Add("symbol", symbol)
        req.URL.RawQuery = query.Encode()
    }

    resp, err := client.Do(req)
    if err != nil {
        logger.Log.Errorf(context.Background(), "BinanceAPIV3 GET /ticker/price err: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Log.Errorf(context.Background(), "read BinanceAPIV3 body err: %v", err)
    }

    return body
}

func GetTicker24Hour(symbol string) []byte {
    client := &http.Client{}
    req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/ticker/24hr", BinanceAPIV3), nil)

    if symbol != "" {
        query := req.URL.Query()
        query.Add("symbol", symbol)
        req.URL.RawQuery = query.Encode()
    }

    resp, err := client.Do(req)
    if err != nil {
        logger.Log.Errorf(context.Background(), "BinanceAPIV3 GET /ticker/24hr err: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Log.Errorf(context.Background(), "read BinanceAPIV3 body err: %v", err)
    }

    return body
}

func GetTickerKLine(symbol string, interval string, startTime, endTime int64) []byte {
    client := &http.Client{}
    req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/klines", BinanceAPIV3), nil)

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
        logger.Log.Errorf(context.Background(), "BinanceAPIV3 GET /klines err: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Log.Errorf(context.Background(), "read BinanceAPIV3 body err: %v", err)
    }

    return body
}

func signature(params *url.Values) *url.Values {
    params.Add("timestamp", cast.ToString(time.Now().Unix()))
    params.Add("recvWindow", cast.ToString(5000))

    h := hmac.New(sha256.New, []byte(config.Config.Binance.Secret))
    h.Write([]byte(params.Encode()))
    sha := hex.EncodeToString(h.Sum(nil))
    params.Add("signature", sha)

    return params
}

func TradeLimit(symbol, side string, quantity, price *float64) []byte {
    client := &http.Client{}
    req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/order", BinanceAPIV3), nil)

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
    query = *signature(&query)
    req.URL.RawQuery = query.Encode()

    resp, err := client.Do(req)
    if err != nil {
        logger.Log.Errorf(context.Background(), "BinanceAPIV3 POST /order err: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Log.Errorf(context.Background(), "read BinanceAPIV3 body err: %v", err)
    }

    return body
}