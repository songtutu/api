package v1

import (
	"ginapi/model"
	"ginapi/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

var code int

//查询用户是否存在
func UserExist(c *gin.Context) {

}

//添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	if data.Username == "" {
		code = errmsg.ERROR_PARAM_NULL_CODE
	} else {
		code = model.CheckUser(data.Username)
		if code == errmsg.SUCCSE_CODE {
			model.CreatUser(&data)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个用户
func GetUser(c *gin.Context) {

}

//查询用户列表
func GetUsers(c *gin.Context) {

}

//编辑用户
func EditUser(c *gin.Context) {

}

//删除用户
func DeleteUser(c *gin.Context) {

}
