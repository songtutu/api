/*
 * @Author: song.tutu
 * @Date: 2020-12-26 09:57:38
 * @Last Modified by:   song.tutu
 * @Last Modified time: 2020-12-26 09:57:38
 */
package routes

import (
	v1 "ginapi/api/v1"
	_ "ginapi/docs"
	"ginapi/utils"
	"ginapi/utils/httpCors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "ginapi/docs"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	r.Use(httpCors.Cors())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rv1 := r.Group("api/v1")
	{
		rv1.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "hello world",
			})
		})
		//User模块的路由接口
		rv1.POST("user/add", v1.AddUser)
		rv1.GET("user/get", v1.GetUsers)
		rv1.POST("user/eidt", v1.EditUser)
		rv1.POST("user/del", v1.DeleteUser)
		//Login模块的路由接口

	}

	r.Run(utils.HttpPort)
}
