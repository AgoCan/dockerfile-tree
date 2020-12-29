package api

import (
	"backend/serializer"
	"backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListResource 资源列表
func ListResource(c *gin.Context) {
	service := service.ListResourceService{}
	res := service.List()
	c.JSON(http.StatusOK, res)
}

// CreateResource 创建资源
func CreateResource(c *gin.Context) {
	service := service.CreateResource{}
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.Error(serializer.ErrCodeParameter))
	} else {
		res := service.Create()
		c.JSON(200, res)
	}
}
