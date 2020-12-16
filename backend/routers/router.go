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
	router.Use(middleware.GinLogger(middleware.Logger),
		middleware.GinRecovery(middleware.Logger, true))
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	dockerGroup := router.Group("/api/v1")
	{
		dockerGroup.GET("/level", controllers.GetLevelList)   // 获取level列表
		dockerGroup.POST("/level", controllers.CreateLevel)   // 创建level
		dockerGroup.DELETE("/level", controllers.DeleteLevel) // 创建level
		// dockerGroup.GET("/level/:id")                       // 获取单个level

		// dockerGroup.GET("/dockerfile/:id") // 获取单个dockerfile
		// dockerGroup.GET("/dockerfile/:id") // 获取dockerfile列表
		// dockerGroup.POST("/dockerfile")    // 创建dockerfile

		// dockerGroup.GET("/resource")  // 获取资源
		// dockerGroup.POST("/resource") // 创建资源
		// dockerGroup.GET("/config")    // 获取配置
		// dockerGroup.POST("/config")   // 创建配置

	}
	return router
}
