package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResData struct {
	Code ResCode     `json:"response"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Error(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
