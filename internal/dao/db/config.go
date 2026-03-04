package db

import "fmt"

type DatabaseConfig struct {
	User         string `yaml:"user"`
	Database     string `yaml:"database"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Password     string `yaml:"password"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
}

func (c DatabaseConfig) GetDNS() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		"utf8mb4",
	)
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	IsTls    bool   `yaml:"is_tls"`
}
