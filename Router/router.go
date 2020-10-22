package Router

import (
	"github.com/gin-gonic/gin"
	"goGinPro/Controllers"
	"goGinPro/Middlewares"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "goGinPro/docs"
)

func InitRouter() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	r.Use(Middlewares.Cors())

	// 路由组
	v1 := r.Group("v1")
	{
		v1.POST("/createmodel", Controllers.CreateModel)
		v1.PUT("/savemodels/:sceneId", Controllers.SaveModels)
		v1.GET("/getmodellist", Controllers.GetAllModel)
		v1.GET("/getscenemodels", Controllers.GetSceneModels)
		v1.GET("/getmodelitem/:id", Controllers.GetModelItem)

		v1.PUT("/editmodel/:id", Controllers.EditModel)

		v1.DELETE("/deletemodel/:id", Controllers.DeleteModel)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8088")
}
