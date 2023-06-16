package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	RateLimitingTimeoutSeconds      uint   `json:"rateLimitingTimeoutSeconds"`
	AdminRateLimitingTimeoutSeconds uint   `json:"adminRateLimitingTimeoutSeconds"`
	DSN                             string `json:"DSN"`
	FirstRun                        bool   `json:"firstRun"`
	TokenAndKeyLength               uint   `json:"tokenAndKeyLength"`
	UrlPrefix                       string `json:"urlPrefix"`
	Debug                           bool   `json:"debug"`
	Email                           struct {
		From string `json:"from"`
		Auth struct {
			Identity   string `json:"identity"`
			Username   string `json:"username"`
			Password   string `json:"password"`
			Host       string `json:"host"`
			Port       int    `json:"port"`
			Encryption string `json:"encryption"`
		} `json:"auth"`
	} `json:"email"`
	InternalAddress string `json:"internalAddress"`
	Domain          string `json:"domain"`
}

var config Config

func ParseConfig() {
	var data Config
	content, err := os.ReadFile("config.json")
	if err != nil {
		panic("config: failed to open config file")
	}

	err = json.Unmarshal(content, &data)
	if err != nil {
		panic("config: failed to read json")
	}
	if data.TokenAndKeyLength%2 != 0 || data.TokenAndKeyLength == 0 {
		panic("config: tokenAndKeyLength should be divisible by 2 and should not be 0")
	}

	config = data
}

func GetConfig() *Config {
	return &config
}
