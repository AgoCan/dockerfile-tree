package db

import (
	"backend/models"
	"fmt"
)

// GetDockerfileByID 获取单个Dockerfile
func GetDockerfileByID(levelID int) (df models.Dockerfile, err error) {
	if levelID < 0 {
		err = fmt.Errorf("invalid parameter,article_id:%d", levelID)
		return
	}
	sqlStr := ` select id,created_at,updated_at,dockerfile,level_id
				from dockerfile
			where level_id = ?`

	err = models.DB.Get(&df, sqlStr, levelID)

	if err != nil {
		return df, err
	}
	return df, nil
}

// CreateUpdateDockerfileByID 获取单个Dockerfile
func CreateUpdateDockerfileByID(content string, levelID int) (err error) {
	if levelID < 0 {
		err = fmt.Errorf("invalid parameter,article_id:%d", levelID)
		return
	}
	var sqlStr string
	var id int
	err = models.DB.Get(&id, "SELECT count(id) FROM dockerfile WHERE level_id=?", levelID)
	if id != 0 {
		sqlStr = `UPDATE dockerfile SET dockerfile = ? where level_id = ?`
	} else {
		sqlStr = `INSERT INTO dockerfile(dockerfile, level_id) VALUES (?, ?)`
	}
	_, err = models.DB.Exec(sqlStr, content, levelID)
	if err != nil {
		return err
	}
	return nil
}
