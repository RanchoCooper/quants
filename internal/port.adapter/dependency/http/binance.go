package http

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

/**
 * @author Rancho
 * @date 2021/12/1
 */

func Ping() string {
    resp, err := http.Get(fmt.Sprintf("%s/ping", BinanceAPIV3))
    if err != nil {
        log.Fatalln(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    return string(body)
}