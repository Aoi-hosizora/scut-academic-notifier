package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	ServerConfig struct {
		PollingDuration int32 `json:"polling-duration"`
	} `json:"server"`
	WxConfig struct {
		Sckey string `json:"sckey"`
	} `json:"wx"`
}

func LoadConfig(path string) (*Config, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = yaml.Unmarshal(f, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
