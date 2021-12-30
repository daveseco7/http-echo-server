package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	configFileEnvKey = "config.yaml"
)

func GetConfig() (*Config, error) {
	return loadConfig(configFileEnvKey)
}

func loadConfig(path string) (*Config, error) {
	configFile, ok := os.LookupEnv(configFileEnvKey)
	if !ok {
		configFile = path
	}

	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
