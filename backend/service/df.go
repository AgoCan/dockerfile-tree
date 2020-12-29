package service

import (
	"backend/models"
	"backend/serializer"
)

// GetDockerfile 获取dockerfile
type GetDockerfile struct {
}

// GetByLevelID 根据level的id获取dockerfile
func (service *GetDockerfile) GetByLevelID(levelID int) serializer.Response {
	var df models.Dockerfile
	if levelID < 0 {
		return serializer.Error(serializer.ErrCodeParameter)
	}
	sqlStr := ` select id,created_at,updated_at,dockerfile,level_id
				from dockerfile
			where level_id = ?`

	err := models.DB.Get(&df, sqlStr, levelID)

	if err != nil {
		return serializer.Error(serializer.ErrSQL)
	}
	res := serializer.BuildDockerfile(df)
	return serializer.Success(res)
}

// Dockerfile 更新和创建使用
type Dockerfile struct {
	Dockerfile string `json:"dockerfile" binding:"required"`
	LevelID    int    `json:"level_id" binding:"required"`
}

// CreateUpdate 创建或更新dockerfile
func (service *Dockerfile) CreateUpdate() serializer.Response {
	if service.LevelID < 0 {
		return serializer.Error(serializer.ErrCodeParameter)
	}

	var count int
	models.DB.Get(&count, "SELECT count(id) FROM dockerfile where level_id=?", service.LevelID)

	var sqlStr string
	if count != 0 {
		sqlStr = `UPDATE dockerfile SET dockerfile = ? where level_id = ?`
	} else {
		sqlStr = `INSERT INTO dockerfile(dockerfile, level_id) VALUES (?, ?)`
	}

	_, err := models.DB.Exec(sqlStr, service.Dockerfile, service.LevelID)
	if err != nil {
		return serializer.Error(serializer.ErrSQL)
	}

	return serializer.Success("ok")
}
