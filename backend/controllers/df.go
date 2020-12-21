package controllers

import (
	"backend/dal/db"
	"backend/utils/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetDockerfileByID 获取单个dockerfile
func GetDockerfileByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("level_id"))
	if err != nil {
		if id != 0 {
			response.Error(c, response.ErrCodeParameter)
			panic(err)
		}
	}
	con, err := db.GetDockerfileByID(id)
	if err != nil {
		response.Error(c, response.ErrSQL)
		panic(err)
	}
	response.Success(c, con)
}

// CreateUpdateDockerfileByID 创建dockerfile
func CreateUpdateDockerfileByID(c *gin.Context) {
	var df DockerfileInfo
	err := c.BindJSON(&df)

	if err != nil {
		response.Error(c, response.ErrCodeParameter)

		panic(err)
	}
	err = db.CreateUpdateDockerfileByID(df.Dockerfile, df.LevelID)
	if err != nil {
		response.Error(c, response.ErrSQL)

		panic(err)
	}
	response.Success(c, "ok")

}
