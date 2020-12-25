package main

import (
	"ginapi/model"
	"ginapi/routes"
)

func main() {
	//引用数据库
	model.InitDb()

	routes.InitRouter()
}
