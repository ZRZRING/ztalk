package mysql

import (
	"fmt"
	"ztalk/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(config *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	zap.L().Debug("mysql info", zap.Any("db", db))
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	return
}

func Close() {
	_ = db.Close()
}