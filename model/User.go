/*
 * @Author: song.tutu
 * @Date: 2020-12-26 09:56:59
 * @Last Modified by:   song.tutu
 * @Last Modified time: 2020-12-26 09:56:59
 */
package model

import (
	"fmt"
	"ginapi/utils/errmsg"

	"ginapi/utils/mymd5"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type: varchar(20);not null" json:"username" form:"username"`
	Password string `gorm:"type: varchar(50);not null" json:"password" form:"password"`
	//角色
	Role int `gorm:"type: int" json:"role" form:"role"`
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

	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR_CODE
	}
	return errmsg.SUCCSE_CODE
}

//查询用户列表
func GetUserList(page int, rows int) ([]User, int) {
	var userList []User
	var count = 0
	err := db.Table("User").Count(&count).Limit(rows).Offset((page - 1) * rows).Find(&userList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	fmt.Println(userList)
	return userList, count
}

//钩子函数
func (u *User) BeforeSave() {
	u.Password = mymd5.Encryption(u.Password)
}
