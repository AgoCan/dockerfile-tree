package routers

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化gin入口，路由信息
func SetupRouter() *gin.Engine {
	router := gin.New()
	if err := middleware.InitLogger(); err != nil {
		panic(err)
	}
	// router.Use(middleware.GinLogger(middleware.Logger),
	// 	middleware.GinRecovery(middleware.Logger, true))
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1Group := router.Group("/api/v1")
	{
		v1Group.GET("/level", controllers.GetLevelList)                     // 获取level列表
		v1Group.POST("/level", controllers.CreateLevel)                     // 创建level
		v1Group.DELETE("/level", controllers.DeleteLevel)                   // 创建level
		v1Group.GET("/dockerfile", controllers.GetDockerfileByID)           // 获取单个dockerfile
		v1Group.POST("/dockerfile", controllers.CreateUpdateDockerfileByID) // 创建或更新dockerfile
		v1Group.GET("/resource")                                            // 获取资源
		v1Group.POST("/resource")                                           // 创建资源
		v1Group.GET("/config")                                              // 获取配置
		v1Group.POST("/config")                                             // 创建配置
	}

	return router
}
