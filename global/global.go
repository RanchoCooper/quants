package global

import (
    "errors"
)

/**
 * @author Rancho
 * @date 2021/12/13
 */

const (
    SimulateUserName     = "rancho-simulate"
    SimulateUserEmail    = "rancho@simulate.com"
    SimulateInitialAsset = 50000
)

var (
    ErrNotFound = errors.New("record not found")
)
