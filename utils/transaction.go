/*
 * @Author: song.tutu
 * @Date: 2020-12-30 09:45:31
 * @Last Modified by: song.tutu
 * @Last Modified time: 2020-12-30 11:25:31
 */
package utils

import (
	"github.com/jinzhu/gorm"
)

//开启事务操作数据库

func ExecSqlWithTransaction(db *gorm.DB, handle func(tx *gorm.DB) error) (err error) {
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	if err = handle(tx); err != nil {
		return err
	}
	tx.Commit()
	return nil
}
