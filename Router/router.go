package Router

import (
    "github.com/gin-gonic/gin"
    "goGinPro/Controllers"
    "goGinPro/Middlewares"
)

func InitRouter() {
    r := gin.Default()

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
    }

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.Run(":8088")
}
