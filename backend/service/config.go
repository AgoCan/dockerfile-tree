package service

import (
	"backend/models"
	"backend/serializer"
	"fmt"

	"github.com/jmoiron/sqlx"
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
	sqlStr = `select 
	l.updated_at,name,comment,order_id,parent_id,dockerfile
		from 
			level l,dockerfile d
		where 
		l.id in (?) and l.id=d.level_id; `
	query, args, err := sqlx.In(sqlStr, []int{1, 2, 3, 8})
	if err != nil {
		println(err)
	}
	rows, err := models.DB.Queryx(query, args...)
	fmt.Println(err)
	var levelList []*models.LevelDockerfile
	for rows.Next() {
		var level models.LevelDockerfile
		fmt.Println(1, rows.StructScan(&level))
		levelList = append(levelList, &level)
	}
	fmt.Println(levelList[0].Name, levelList[0].UpdatedAt)
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
