package binance

import (
    "quants/internal/port.adapter/dependency/http"
)

/**
 * @author Rancho
 * @date 2021/12/1
 */

type Binance struct {

}

func (b Binance) Ping() string {
    return http.Ping()
}