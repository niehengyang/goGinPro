package Mysql

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "goGinPro/Services"
)

// 3. 初始化全局变量 / 也可以在函数内定义
var (
    DB *gorm.DB
)

func init() {
    var err error

    configMap := Services.InitConfig("Config/app.conf")

    dbhost := configMap["dbhost"]
    dbport := configMap["dbport"]
    dbuser := configMap["dbuser"]
    dbpassword := configMap["dbpassword"]
    db := configMap["db"]

    connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbport, db)

    DB, err = gorm.Open("mysql", connArgs)
    if err != nil {
        fmt.Printf("mysql connect error %v", err)
        return
    }

    if DB.Error != nil {
        fmt.Printf("database error %v", DB.Error)
        return
    }

    DB.SingularTable(true)

    fmt.Printf("数据库连接成功！\n")
}
