package api

import (
	"backend/serializer"
	"backend/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Level create level
type Level struct {
	ID       int    `json:"id"`
	ParentID int    `json:"parent_id"`
	Name     string `json:"name"`
	OrderID  int    `json:"order_id"`
	Comment  string `json:"comment,omitempty"`
}

// ListLevel 列出所有的level
func ListLevel(c *gin.Context) {
	service := service.ListLevelService{}
	res := service.List()
	c.JSON(http.StatusOK, res)
}

// UpdateLevel 列出所有的level
func UpdateLevel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(200, serializer.Error(serializer.ErrCodeParameter))
	}
	service := service.UpdateLevelService{}
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.Error(serializer.ErrCodeParameter))
	} else {
		res := service.Update(id)
		c.JSON(200, res)
	}
}

// CreateLevel 列出所有的level
func CreateLevel(c *gin.Context) {
	service := service.CreateLevelService{}
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.Error(serializer.ErrCodeParameter))
	} else {
		res := service.Create()
		c.JSON(200, res)
	}

}

// DeleteLevel 列出所有的level
func DeleteLevel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(200, serializer.Error(serializer.ErrCodeParameter))
	}
	services := service.DeleteLevelService{}
	res := services.Delete(id)
	c.JSON(http.StatusOK, res)
}
