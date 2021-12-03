package file

import (
    "context"
    "io/ioutil"
    "os"
    "path/filepath"

    "quants/util/logger"
)

/**
 * @author Rancho
 * @date 2021/12/3
 */

func ReadFile(filePath string) (content []byte) {
    filePath, err := filepath.Abs(filePath)
    if err != nil {
        return
    }
    file, err := os.Open(filePath)
    if err != nil {
        logger.Log.Errorf(context.Background(), "os.Open file fail. filePath: %s, err: %v", filePath, err)
        return nil
    }
    defer file.Close()

    content, _ = ioutil.ReadAll(file)
    return
}