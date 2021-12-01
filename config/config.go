package config

import (
    "io/ioutil"
    "log"
    "path/filepath"

    "gopkg.in/yaml.v3"
)

const (
	configFilePath = "config/config.yaml"
    privateConfigFilePath = "config/config.private.yaml"
)

type binanceConfig struct {
	Key    string `yaml:"key"`
	Secret string `yaml:"secret"`
}

type config struct {
	Binance *binanceConfig `yaml:"binance"`
}

var Config = &config{}

func readYamlConfig(configPath string) {
    yamlFile, err := filepath.Abs(configPath)
    if err != nil {
        log.Fatalf("invalid config file path, err: %v", err)
    }
    content, err := ioutil.ReadFile(yamlFile)
    if err != nil {
        log.Fatalf("read config file fail, err: %v", err)
    }
    err = yaml.Unmarshal(content, Config)
    if err != nil {
        log.Fatalf("config file unmarshal fail, err: %v", err)
    }
}

func init() {
    readYamlConfig(configFilePath)

    if Config.Binance.Key == "" || Config.Binance.Secret == "" {
        // read private config
        readYamlConfig(privateConfigFilePath)
    }
}
