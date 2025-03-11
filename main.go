package main

import (
	"fmt"
	"os"
	"ztalk/internal/mysql"
	"ztalk/internal/redis"
	"ztalk/internal/router"
	"ztalk/pkg/logger"
	"ztalk/pkg/settings"
	"ztalk/pkg/translate"
	"ztalk/pkg/utils"

	"go.uber.org/zap"
)

func syncLogger(l *zap.Logger) {
	err := l.Sync()
	if err != nil {
		fmt.Printf("sync logger failed, message:%v\n", err)
	}
}

// @title ztalk
// @version 1.0
// @description 在线论坛平台
// @termsOfService http://swagger.io/terms/
// @contact.name zrzring
// @contact.url http://zrzring.com/support
// @contact.email zrzring@126.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host http://127.0.0.1
// @BasePath /api/v1
func main() {
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: dev.yaml")
		return
	}

	// 加载配置文件
	if err := settings.Init(os.Args[1]); err != nil {
		fmt.Printf("配置文件加载失败, message:%v\n", err)
		return
	}
	fmt.Println("配置文件加载成功")

	cfg := settings.Conf

	// 初始化日志系统
	if err := logger.Init(cfg.LogConfig, cfg.Mode); err != nil {
		fmt.Printf("logger\t初始化失败, message:%v\n", err)
		return
	}
	// defer syncLogger(zap.L())
	zap.L().Info("logger\t初始化完成")

	// 初始化 MySQL 连接
	if err := mysql.Init(cfg.MySQLConfig); err != nil {
		zap.L().Error("mysql\t初始化失败", zap.Error(err))
		return
	}
	zap.L().Info("mysql\t初始化完成")
	defer mysql.Close()

	// 初始化 Redis 连接
	if err := redis.Init(cfg.RedisConfig); err != nil {
		zap.L().Error("redis\t初始化失败", zap.Error(err))
		return
	}
	zap.L().Info("redis\t初始化完成")
	defer redis.Close()

	// 初始化 Snowflake 包
	if err := utils.InitSnowflake(cfg.StartTime, cfg.MachineID); err != nil {
		zap.L().Error("snowflake\t初始化失败", zap.Error(err))
		return
	}
	zap.L().Info("snowflake\t初始化完成")

	// 初始化 Validator 包
	if err := translate.Init("zh"); err != nil {
		zap.L().Error("translate\t初始化失败", zap.Error(err))
		return
	}
	zap.L().Info("translate\t初始化完成")

	// 注册并启动路由器
	mainRouter := router.Setup(cfg.Mode)
	if err := mainRouter.Run(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		zap.L().Error("路由启动失败", zap.Error(err))
		return
	}
	zap.L().Info("路由启动完成")

	// 服务器成功关闭，打印退出日志（失效）
	// zap.L().Info("Server exiting")
}
