package controllers

import (
	"backend/dal/db"
	"backend/utils/response"

	"github.com/gin-gonic/gin"
)

// CreateLevel 创建level
func CreateLevel(c *gin.Context) {

}

// GetLevelList 获取level list
func GetLevelList(c *gin.Context) {

	levelList, err := db.GetLevelList()
	if err != nil {
		response.Error(c, response.ErrSQLList)
		panic(err)
	}
	response.Success(c, levelList)
}
