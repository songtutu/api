package utils

import (
	"ginapi/utils/errmsg"

	"github.com/gin-gonic/gin"
)

type Swagger struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ReturnJson(c *gin.Context, status int, code int, data interface{}) {
	var swagger Swagger
	swagger.Code = code
	swagger.Data = data
	swagger.Message = errmsg.GetErrMsg(code)
	c.JSON(status, Swagger{
		Code:    code,
		Message: errmsg.GetErrMsg(code),
		Data:    data,
	})
}

func ReturnPageJson(c *gin.Context, status int, code int, data interface{}, total int) {
	c.JSON(status, gin.H{
		"code":    code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"total":   total,
	})
}
