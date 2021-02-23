package api

import (
	"backend/serializer"
	"backend/service"

	"github.com/gin-gonic/gin"
)

// CreateBuildJob 构建镜像
func CreateBuildJob(c *gin.Context) {
	service := service.BuildJob{}
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.Error(serializer.ErrCodeParameter))
	} else {
		res := service.Create()
		c.JSON(200, res)
	}
}
