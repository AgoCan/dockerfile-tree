package service

import (
	"backend/models"
	"backend/serializer"
)

// ListConfigService 配置列表
type ListConfigService struct {
}

// List 列表资源
func (service *ListConfigService) List() serializer.Response {
	var configList []*models.Config
	sqlStr := `select id,created_at, updated_at, config_key, config_value, config_comment
	from config `
	err := models.DB.Select(&configList, sqlStr)
	if err != nil {
		return serializer.Error(serializer.ErrSQL)
	}

	res := serializer.BuildConfigs(configList)

	return serializer.Success(res)
}

// CreateConfig 创建配置文件
type CreateConfig struct {
	Key     string  `json:"config_key"`
	Value   *string `json:"config_value"`
	Comment *string `json:"config_comment"`
}

// Create 创建
func (service *CreateConfig) Create() serializer.Response {

	var id int
	err := models.DB.Get(&id, "SELECT count(id) FROM config where config_key=?", service.Key)
	if err != nil {
		return serializer.Error(serializer.ErrSQL)
	}

	if id != 0 {
		str := service.Key + serializer.GetMessage(serializer.ErrSQLExist)
		return serializer.ErrorStr(serializer.ErrSQLExist, str)
	}

	sqlStr := `INSERT INTO config(config_key, config_value, config_comment)
		VALUES (?,?,?)`
	_, err = models.DB.Exec(sqlStr, service.Key, service.Value, service.Comment)
	if err != nil {
		return serializer.Error(serializer.ErrSQL)
	}

	return serializer.Success("ok")
}
