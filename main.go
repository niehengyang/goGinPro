package main

import (
	"goGinPro/Databases"
	"goGinPro/Models"
	"goGinPro/Router"
)

func main() {

	//当整个程序完成之后关闭数据库连接
	defer Mysql.DB.Close()
	Models.StartMigrate(Mysql.DB)

	//路由
	Router.InitRouter()
}
