package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"ztalk/app/repository/mysql"
	"ztalk/app/repository/redis"
	"ztalk/app/utils"
	"ztalk/logger"
	"ztalk/router"
	"ztalk/settings"
)

// web 应用脚手架

func syncLogger(l *zap.Logger) {
	err := l.Sync()
	if err != nil {
		fmt.Printf("sync logger failed, message:%v\n", err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: config.yaml")
		return
	}

	// 加载配置文件
	if err := settings.Init(os.Args[1]); err != nil {
		fmt.Printf("init settings failed, message:%v\n", err)
		return
	}

	cfg := settings.Conf
	
	// 初始化日志系统
	if err := logger.Init(cfg.LogConfig, cfg.Mode); err != nil {
		fmt.Printf("init logger failed, message:%v\n", err)
		return
	}

	defer syncLogger(zap.L())
	
	zap.L().Debug("logger init success")

	// 初始化 MySQL 连接
	if err := mysql.Init(cfg.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, message:%v\n", err)
		return
	}

	zap.L().Debug("mysql init success")

	defer mysql.Close()

	// 初始化 Redis 连接
	if err := redis.Init(cfg.RedisConfig); err != nil {
		fmt.Printf("init redis failed, message:%v\n", err)
		return
	}

	zap.L().Debug("redis init success")

	defer redis.Close()

	// 初始化 Snowflake 包
	if err := utils.InitSnowflake(cfg.StartTime, cfg.MachineID); err != nil {
		fmt.Printf("init snowflake failed, message:%v\n", err)
		return
	}

	// 初始化 Validator 包
	if err := utils.InitValidator("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}

	// 注册并启动路由器
	mainRouter := router.Setup(cfg.Mode)

	if err := mainRouter.Run(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		fmt.Printf("Run http server faild, message:%v\n", err)
		return
	}

	// 服务器成功关闭，打印退出日志（失效）
	// zap.L().Info("Server exiting")
}
