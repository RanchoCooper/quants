package util

import (
    "math/rand"
    "path/filepath"
    "runtime"
    "time"
)

/**
 * @author Rancho
 * @date 2021/12/5
 */

const (
    AsciiLowercase = "abcdefghijklmnopqrstuvwxyz"
    AsciiUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    AsciiLetters   = AsciiLowercase + AsciiUppercase
    Digits         = "0123456789"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func GetCurrentPath() string {
    _, file, _, _ := runtime.Caller(1)
    return filepath.Dir(file)
}

// RandString returns random string.
func RandString(length int, onlyDigital bool) string {
    var char string
    if onlyDigital {
        char = Digits
    } else {
        char = AsciiLetters + Digits
    }
    bytes := make([]byte, length)
    for i := 0; i < length; i++ {
        bytes[i] = char[rand.Intn(len(char)-1)]
    }
    return string(bytes)
}
