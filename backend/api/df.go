package api

import (
	"backend/serializer"
	"backend/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetDockerfileByLevelID 获取dockerfile
func GetDockerfileByLevelID(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("level_id"))
	if err != nil {
		c.JSON(200, serializer.Error(serializer.ErrCodeParameter))
	}
	service := service.GetDockerfile{}
	res := service.GetByLevelID(id)
	c.JSON(http.StatusOK, res)
}

// CreateUpdateDockerfile 创建或更新dockerfile
func CreateUpdateDockerfile(c *gin.Context) {
	service := service.Dockerfile{}
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.Error(serializer.ErrCodeParameter))
	} else {
		res := service.CreateUpdate()
		c.JSON(200, res)
	}
}
