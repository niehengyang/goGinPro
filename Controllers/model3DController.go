package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goGinPro/Models"
	"goGinPro/Services"
)

func CreateModel(c *gin.Context) {
	var data Models.Models
	responseServer := Services.Gin{Ctx: c}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		fmt.Print("数据获取失败！")
	}

	//创建
	modelId, result := Models.CreateModel(data)

	if result {
		responseServer.Response(0, "创建成功", modelId)
	}
}

func EditModel(c *gin.Context) {
	var data Models.Models
	responseServer := Services.Gin{Ctx: c}

	modelId := c.Param("id")
	err := c.ShouldBindJSON(&data)
	if err != nil {
		fmt.Print("数据获取失败！")
	}

	//编辑
	result := Models.EditModel(Services.StringToInt(modelId), data)

	if result {
		responseServer.Response(0, "创建成功", nil)
	}
}

func GetModelItem(c *gin.Context) {

	responseServer := Services.Gin{Ctx: c}
	id := c.Param("id")

	//查找
	modelItem, status := Models.GetModelItem(Services.StringToInt(id))

	if status {
		responseServer.Response(0, "查询成功", modelItem)
	} else {
		responseServer.Response(1, "查询失败", modelItem)
	}
}

func GetAllModel(c *gin.Context) {

	groupId := c.DefaultQuery("groupId", "0")

	mdoels := Models.GetAllModel(Services.StringToInt(groupId))

	responseServer := Services.Gin{Ctx: c}
	responseServer.Response(0, "查询成功", mdoels)
}

func SaveModels(c *gin.Context) {
	var data []Models.SceneToModel
	responseServer := Services.Gin{Ctx: c}

	sceneId := c.Param("sceneId")
	err := c.ShouldBindJSON(&data)
	if err != nil {
		fmt.Print("数据获取失败！")
	}

	//save
	status := Models.SaveModels(Services.StringToInt(sceneId), data)

	if status {
		responseServer.Response(0, "保存成功", nil)
	}
}

func GetSceneModels(c *gin.Context) {
	mdoels := Models.GetSceneModels()

	responseServer := Services.Gin{Ctx: c}
	responseServer.Response(0, "查询成功", mdoels)
}
