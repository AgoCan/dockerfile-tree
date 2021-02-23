package routers

import (
	"backend/api"
	"backend/middleware/log"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化gin入口，路由信息
func SetupRouter() *gin.Engine {
	router := gin.New()
	if err := log.InitLogger(); err != nil {
		panic(err)
	}
	router.Use(log.GinLogger(log.Logger),
		log.GinRecovery(log.Logger, true))

	v1Group := router.Group("/api/v1")
	{
		v1Group.GET("/level", api.ListLevel)                    // 获取level列表
		v1Group.POST("/level", api.CreateLevel)                 // 创建level
		v1Group.PUT("/level/:id", api.UpdateLevel)              // 创建level
		v1Group.DELETE("/level/:id", api.DeleteLevel)           // 删除level

		v1Group.GET("/dockerfile", api.GetDockerfileByLevelID)  // 获取单个dockerfile
		v1Group.POST("/dockerfile", api.CreateUpdateDockerfile) // 创建或更新dockerfile

		v1Group.GET("/resource", api.ListResource)              // 获取资源
		// v1Group.PUT("/resource/:id", api.UpdateResource)        // 更新资源
		v1Group.POST("/resource", api.CreateResource) // 创建资源
		// v1Group.DELETE("/resource/:id", api.DeleteResource) // 删除资源

		v1Group.GET("/config", api.ListConfig)    // 获取配置
		v1Group.POST("/config", api.CreateConfig) // 创建配置，该记录可能不存在，后期优化

		v1Group.POST("/build", api.CreateBuildJob)
	}
	return router
}
