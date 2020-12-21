package db

import (
	"backend/models"
	"fmt"
)

/*
db.MustExec("BEGIN;")
db.MustExec(...)
db.MustExec("COMMIT;")
*/

// GetCinfig 获取配置
func GetCinfig() (configList []*models.Config, err error) {
	sqlStr := `select id,created_at, updated_at, config_key, config_value, config_comment
			from config `
	err = models.DB.Select(&configList, sqlStr)

	if err != nil {
		return nil, err
	}
	return configList, nil
}

// UpdateConfig 更新配置文件
func UpdateConfig(key string, value, comment *string) (err error) {
	sqlStr := `UPDATE config SET config_value = ?, config_comment = ? where config_key = ?`
	_, err = models.DB.Exec(sqlStr, value, comment, key)
	fmt.Println(value, comment, key)
	if err != nil {
		return err
	}
	return nil
}
