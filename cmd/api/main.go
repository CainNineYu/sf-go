package main

import (
	"sf-go/internal/api"
	"sf-go/internal/config"
	dbpkg "sf-go/internal/dao/db"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"os"
	"sf-go/logs"
)

func main() {
	// 设置日志
	logs.Setlogs(zap.DebugLevel)

	// 加载配置文件，从路径
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		logs.Logger.Fatal("配置文件路径未设置")
	}
	open, err := os.Open(configPath)
	if err != nil {
		logs.Logger.Fatal("打开配置文件失败", zap.Error(err))
	}

	decoder := yaml.NewDecoder(open)
	cfg := &config.ApiSrvCfg{}
	if err = decoder.Decode(cfg); err != nil {
		logs.Logger.Fatal("解析配置文件失败", zap.Error(err))
	}
	cloErr := open.Close()
	if cloErr != nil {
		logs.Logger.Fatal("关闭文件失败", zap.Error(cloErr))
	}
	// 初始化数据库
	dbCfg, err := dbpkg.NewDB(cfg.WriterDatabase, cfg.ReaderDatabase)
	if err != nil {
		logs.Logger.Fatal("初始化数据库失败", zap.Error(err))
	}

	// 4. 启动 gin，端口号用 cfg.GinPort
	r := api.Router(dbCfg, cfg)

	// 注册路由...
	err = r.Run(cfg.Server.Port)
	if err != nil {
		logs.Logger.Fatal("Init gin error", zap.Error(err))
	}
}
