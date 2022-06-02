package config

import "time"

func NewConfig() Config {
	return Config{
		ID: time.Now().Unix(),
	}
}
