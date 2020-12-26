package main

import (
	"ginapi/model"
	"ginapi/routes"
)

func main() {
	//引用数据库
	db := model.InitDb()
	routes.InitRouter()
	db.Close()
}
