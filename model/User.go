/*
 * @Author: song.tutu
 * @Date: 2020-12-26 09:56:59
 * @Last Modified by:   song.tutu
 * @Last Modified time: 2020-12-26 09:56:59
 */
package model

import (
	"ginapi/utils/errmsg"

	"github.com/jinzhu/gorm"

	"ginapi/utils/mymd5"
)

type User struct {
	gorm.Model
	Username string `gorm:"type: varchar(20);not null" json:"username"`
	Password string `gorm:"type: varchar(20);not null" json:"password"`
	//角色
	Role int `gorm:"type: int" json:"role"`
}

//检查用户是否存在
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED_CODE
	}
	return errmsg.SUCCSE_CODE
}

//添加用户
func CreatUser(data *User) int {
	data.Password = mymd5.Encryption(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR_CODE
	}
	return errmsg.SUCCSE_CODE
}
