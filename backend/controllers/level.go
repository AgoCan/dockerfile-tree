package controllers

import (
	"backend/dal/db"
	"backend/utils/response"

	"github.com/gin-gonic/gin"
)

// CreateLevel 创建level
func CreateLevel(c *gin.Context) {
	var levelInfo LevelInfo

	err := c.BindJSON(&levelInfo)

	if err != nil {
		response.Error(c, response.ErrCodeParameter)

		panic(err)
	}
	id, err := db.CreateLevel(levelInfo.ParentID, levelInfo.OrderID, levelInfo.Name)
	if err != nil {
		response.Error(c, response.ErrSQL)

		panic(err)
	}
	response.Success(c, id)
}

// GetLevelList 获取level list
func GetLevelList(c *gin.Context) {

	levelList, err := db.GetLevelList()
	if err != nil {
		response.Error(c, response.ErrSQL)
		panic(err)
	}
	response.Success(c, levelList)
}

// DeleteLevel 删除level
func DeleteLevel(c *gin.Context) {
	var delInfo DelLevelInfo
	err := c.BindJSON(&delInfo)
	if err != nil {
		response.Error(c, response.ErrCodeParameter)
		panic(err)
	}
	err = db.DeleteLevel(delInfo.ID)
	if err != nil {
		response.Error(c, response.ErrSQL)
		panic(err)
	}
	response.Success(c, "ok")
}
