package request

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetPageInfo 获取分页参数
func GetPageInfo(c *gin.Context) (page int64, size int64) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	var err error
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
