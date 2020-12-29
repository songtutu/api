package utils

import (
	"ginapi/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func ReturnJson(c *gin.Context, status int, code int, data interface{}) {
	c.JSON(status, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func ReturnPageJson(c *gin.Context, status int, code int, data interface{}, total int) {
	c.JSON(status, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"total":   total,
	})
}
