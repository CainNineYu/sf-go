package config

import (
	"sf-go/internal/dao/db"
)

type ApiSrvCfg struct {
	Server         ApiServiceCfg     `yaml:"server"`
	ReaderDatabase db.DatabaseConfig `yaml:"reader_database"`
	WriterDatabase db.DatabaseConfig `yaml:"writer_database"`
}

type ApiServiceCfg struct {
	Port      string `yaml:"port"`
	JwtSecret string `yaml:"jwt_secret"`
}
