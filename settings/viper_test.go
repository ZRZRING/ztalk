package settings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySQLConfig(t *testing.T) {
	// 初始化配置
	err := Init("../dev.yaml")
	assert.NoError(t, err, "Failed to load config file")

	// 获取 MySQL 配置
	mysqlConfig := Conf.MySQLConfig

	// 断言 MySQL 配置中的各个字段的值是否正确
	assert.Equal(t, "127.0.0.1", mysqlConfig.Host, "MySQL host should be '127.0.0.1'")
	assert.Equal(t, 3306, mysqlConfig.Port, "MySQL port should be 3306")
	assert.Equal(t, "root", mysqlConfig.Username, "MySQL username should be 'root'")
	assert.Equal(t, "123456", mysqlConfig.Password, "MySQL password should be '123456'")
	assert.Equal(t, "ztalk", mysqlConfig.Database, "MySQL database should be 'ztalk'")
	assert.Equal(t, 200, mysqlConfig.MaxOpenConns, "MySQL max_open_conns should be 200")
	assert.Equal(t, 50, mysqlConfig.MaxIdleConns, "MySQL max_idle_conns should be 50")
}
