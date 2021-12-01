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
        logger.Log.Errorf(context.Background(), "ping binanceAPIV3 err: %v", err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Log.Errorf(context.Background(), "read ping binanceAPIV3 body err: %v", err)
    }

    return string(body)
}