package api

import (
	"backend/serializer"
	"backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListConfig 配置列表
func ListConfig(c *gin.Context) {
	service := service.ListConfigService{}
	res := service.List()
	c.JSON(http.StatusOK, res)
}

// UpdateConfig 更新配置
func UpdateConfig(c *gin.Context) {
	service := service.CreateConfig{}
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.Error(serializer.ErrCodeParameter))
	} else {
		res := service.Create()
		c.JSON(200, res)
	}
}
