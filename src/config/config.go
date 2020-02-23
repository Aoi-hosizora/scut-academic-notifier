package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ServerConfig struct {
	PollingDuration int `yaml:"polling-duration"` // second
	SendMaxCount    int `yaml:"send-max-count"`
	SendRange       int `yaml:"send-range"` // day
}

type WeChatConfig struct {
	Sckey string `yaml:"sckey"`
}

type Config struct {
	ServerConfig *ServerConfig `yaml:"server"`
	WeChatConfig *WeChatConfig `yaml:"wechat"`
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
