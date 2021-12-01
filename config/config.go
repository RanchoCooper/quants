package config

import (
	"io/ioutil"
	"log"
	"path/filepath"
    "runtime"

    "gopkg.in/yaml.v3"
)

const (
	configFilePath        = "/config.yaml"
	privateConfigFilePath = "/config.private.yaml"
)

type binanceConfig struct {
	Key    string `yaml:"key"`
	Secret string `yaml:"secret"`
}

type logConfig struct {
	LogSavePath string `yaml:"log_save_path"`
	LogFileName string `yaml:"log_file_name"`
	LogFileExt  string `yaml:"log_file_ext"`
}

type config struct {
	Binance *binanceConfig `yaml:"binance"`
	Log     *logConfig     `yaml:"log"`
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
    _, file, _, _ := runtime.Caller(0)
    rootPath := filepath.Dir(file)

	readYamlConfig(rootPath + configFilePath)

    // read private sensitive configs
	if Config.Binance.Key == "" || Config.Binance.Secret == "" {
		// read private config
		readYamlConfig(rootPath + privateConfigFilePath)
	}
}
