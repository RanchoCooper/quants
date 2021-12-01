package http

import (
    "context"
    "fmt"
    "io/ioutil"
    "net/http"
    "strconv"

    "go-hexagonal/util/logger"
)

/**
 * @author Rancho
 * @date 2021/12/1
 */

func Ping() string {
    resp, err := http.Get(fmt.Sprintf("%s/ping", BinanceAPIV3))
    if err != nil {
        logger.Log.Errorf(context.Background(), "binanceAPIV3 GET /ping err: %v", err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Log.Errorf(context.Background(), "read binanceAPIV3 body err: %v", err)
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
        logger.Log.Errorf(context.Background(), "binanceAPIV3 GET /ticker/price err: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Log.Errorf(context.Background(), "read binanceAPIV3 body err: %v", err)
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
        logger.Log.Errorf(context.Background(), "binanceAPIV3 GET /ticker/24hr err: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Log.Errorf(context.Background(), "read binanceAPIV3 body err: %v", err)
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
        logger.Log.Errorf(context.Background(), "binanceAPIV3 GET /klines err: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Log.Errorf(context.Background(), "read binanceAPIV3 body err: %v", err)
    }

    return body
}