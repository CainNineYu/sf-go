package config

import (
	"sf-go/internal/dao/db"
)

type ConsoleSrvCfg struct {
	Server         ConsoleServiceCfg `yaml:"server"`
	ReaderDatabase db.DatabaseConfig `yaml:"reader_database"`
	WriterDatabase db.DatabaseConfig `yaml:"writer_database"`
	RedisCfg       db.RedisConfig    `yaml:"redis"`
}

type ConsoleServiceCfg struct {
	Port      string `yaml:"port"`
	JwtSecret string `yaml:"jwt_secret"`
}
