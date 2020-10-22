package main

import (
	"goGinPro/Databases"
	"goGinPro/Models"
	"goGinPro/Router"

	_ "goGinPro/docs"
)

// @title Three.js 项目API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://www.topgoer.com

// @contact.name www.topgoer.com
// @contact.url https://www.topgoer.com
// @contact.email 790227542@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8088
// @BasePath v1

func main() {

	//当整个程序完成之后关闭数据库连接
	defer Mysql.DB.Close()
	Models.StartMigrate(Mysql.DB)

	//路由
	Router.InitRouter()
}
