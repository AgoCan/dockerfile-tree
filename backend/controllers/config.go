package controllers

/*
一次性获取所有的配置文件，然后展示出来。并且value可以配置。而key是固定的
*/

import (
	"backend/dal/db"
	"backend/utils/response"

	"github.com/gin-gonic/gin"
)

// GetConfig 获取配置
func GetConfig(c *gin.Context) {
	configList, err := db.GetCinfig()
	if err != nil {
		response.Error(c, response.ErrSQL)
		panic(err)
	}
	response.Success(c, configList)
}

// // CreateConfig 创建配置项
// func CreateConfig(c *gin.Context) {

// }

// UpdateConfig 更新配置项
func UpdateConfig(c *gin.Context) {
	// 只能更新value
	var configs Configs

	err := c.BindJSON(&configs)
	if err != nil {
		response.Error(c, response.ErrCodeParameter)

		panic(err)
	}
	err = db.UpdateConfig(configs.Key, configs.Value, configs.Comment)
	if err != nil {
		response.Error(c, response.ErrSQL)

		panic(err)
	}
	response.Success(c, "ok")
}
