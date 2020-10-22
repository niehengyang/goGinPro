package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goGinPro/Models"
	"goGinPro/Services"
)

// @Summary 新建3D模型接口
// @Description 新建3D模型接口
// @Tags 模型相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param data body Services.Models true "3D模型结构体"
// @Security ApiKeyAuth
// @Success 200 {object} Services.response
// @Router /createmodel [post]
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

// @Summary 编辑3D模型接口
// @Description 编辑3D模型接口
// @Tags 模型相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path int true "3D模型ID"
// @Param data body Services.Models true "3D模型结构体"
// @Security ApiKeyAuth
// @Success 200 {object} Services.response
// @Router /editmodel [put]
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

// @Summary 查询单个模型
// @Description 查询单个模型接口
// @Tags 模型相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id query int true "3D模型ID"
// @Security ApiKeyAuth
// @Success 200 {object} Services.response
// @Router /editmodel [get]
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

// @Summary 查询模型列表
// @Description 查询模型列表接口
// @Tags 模型相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param groupId query int false "3D模型组ID"
// @Security ApiKeyAuth
// @Success 200 {object} Services.response
// @Router /editmodel [get]
func GetAllModel(c *gin.Context) {

	groupId := c.DefaultQuery("groupId", "0")

	mdoels := Models.GetAllModel(Services.StringToInt(groupId))

	responseServer := Services.Gin{Ctx: c}
	responseServer.Response(0, "查询成功", mdoels)
}

// @Summary 场景模型保存
// @Description 场景模型保存接口
// @Tags 模型相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param sceneId path integer true "场景ID"
// @Param data body []Services.SceneToModel true "3D模型数组"
// @Security ApiKeyAuth
// @Success 200 {object} Services.response
// @Router /editmodel [put]
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

// @Summary 查询场景绑定模型
// @Description 查询场景绑定模型接口
// @Tags 模型相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param sceneId query integer true "场景ID"
// @Security ApiKeyAuth
// @Success 200 {object} Services.response
// @Router /editmodel [get]
func GetSceneModels(c *gin.Context) {
	sceneId := c.Param("sceneId")
	mdoels := Models.GetSceneModels(Services.StringToInt(sceneId))

	responseServer := Services.Gin{Ctx: c}
	responseServer.Response(0, "查询成功", mdoels)
}

func DeleteModel(c *gin.Context) {

	modelId := c.Param("id")
	responseServer := Services.Gin{Ctx: c}
	status := Models.DeleteModel(modelId)

	if status {
		responseServer.Response(0, "删除成功", nil)
	} else {
		responseServer.Response(1, "删除失败", nil)
	}

}
