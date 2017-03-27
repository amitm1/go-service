package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Memcached             MemcachedConfig
	Debug                 string
	DefaultCacheExpireSec int
}

type MemcachedConfig struct {
	Host string
	Port int
}

func GetConfig() *Config {
	conf := Config{}

	configFile := "config/config.cfg"
	if _, err := toml.DecodeFile(configFile, &conf); err != nil {
		log.Fatal(err)
	}

	return &conf
}
