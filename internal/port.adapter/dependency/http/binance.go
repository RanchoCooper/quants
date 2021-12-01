package http

import (
    "context"
    "fmt"
    "io/ioutil"
    "net/http"

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