/*
 * @Author: song.tutu
 * @Date: 2020-12-26 09:57:25
 * @Last Modified by:   song.tutu
 * @Last Modified time: 2020-12-26 09:57:25
 */
package v1

import (
	"ginapi/model"
	"ginapi/utils"
	"ginapi/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var code int

//查询用户是否存在
func UserExist(c *gin.Context) {

}

//添加用户
func AddUser(c *gin.Context) {
	var data model.User
	if err := c.ShouldBind(&data); err != nil {
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User information is not complete " + err.Error()})
			return
		}
	}
	if data.Username == "" || data.Password == "" {
		code = errmsg.ERROR_PARAM_NULL_CODE
	} else {
		code = model.CheckUser(data.Username)
		if code == errmsg.SUCCSE_CODE {
			model.CreatUser(&data)
		}
	}
	utils.ReturnJson(c, http.StatusOK, code, data)
}

//查询单个用户
func GetUser(c *gin.Context) {

}

//查询用户列表
func GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	rows, _ := strconv.Atoi(c.Query("rows"))
	//-1是不加限制
	if page == 0 {
		page = -1
	}
	if rows == 0 {
		rows = -1
	}
	var count = 0
	userList, count := model.GetUserList(page, rows)
	utils.ReturnPageJson(c, http.StatusOK, errmsg.SUCCSE_CODE, userList, count)
}

//编辑用户
func EditUser(c *gin.Context) {

}

//删除用户
func DeleteUser(c *gin.Context) {

}
