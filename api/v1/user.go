/*
 * @Author: song.tutu
 * @Date: 2020-12-26 09:57:25
 * @Last Modified by: song.tutu
 * @Last Modified time: 2020-12-30 10:21:19
 */
package v1

import (
	"fmt"
	"ginapi/model"
	"ginapi/utils"
	"ginapi/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var code int

func UserExist(c *gin.Context) {

}

// @Summary Add user
// @Produce  json
// @Param user body model.User true "User"
// @Success 200 {ojbect} utils.Swagger
// @Router /api/v1/user/add [post]
func AddUser(c *gin.Context) {
	fmt.Println(c.Param("id"))
	var data model.User
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User information is not complete " + err.Error()})
		return
	}
	if data.Username == "" || data.Password == "" {
		code = errmsg.ERROR_PARAM_NULL_CODE
	} else {
		code = model.CheckUser(data.Username)
		if code == errmsg.SUCCSE_CODE {
			code = model.CreatUser(&data)
		}
	}
	utils.ReturnJson(c, http.StatusOK, code, data)
}

func GetUser(c *gin.Context) {

}

func GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	rows, _ := strconv.Atoi(c.Query("rows"))
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

func EditUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
