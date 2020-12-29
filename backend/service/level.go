package service

import (
	"backend/models"
	"backend/serializer"
)

// ListLevelService level服务
type ListLevelService struct{}

// List 列表ListLevel
func (service *ListLevelService) List() serializer.Response {
	var levelList []*models.Level
	sqlStr := ` select id,created_at,updated_at,name,comment,order_id,parent_id from level`
	err := models.DB.Select(&levelList, sqlStr)
	if err != nil {
		return serializer.Error(serializer.ErrSQL)
	}

	res := serializer.BuildLevels(levelList)

	return serializer.Success(res)
}

// CreateLevelService 创建level
type CreateLevelService struct {
	ParentID int    `json:"parent_id"`
	Name     string `json:"name" binding:"required,min=2,max=64"`
	OrderID  int    `json:"order_id"`
	Comment  string `json:"comment,omitempty"`
}

// Create 创建
func (service *CreateLevelService) Create() serializer.Response {

	// var id int
	// err := models.DB.Get(&id, "SELECT count(id) FROM level where name=?", service.Name)

	// if id != 0 {
	// 	str := service.Name + serializer.GetMessage(serializer.ErrSQLExist)
	// 	return serializer.ErrorStr(serializer.ErrSQLExist, str)
	// }

	sqlStr := `INSERT INTO level(name, order_id, parent_id)
		VALUES (?,?,?)`
	_, err := models.DB.Exec(sqlStr, service.Name, service.OrderID, service.ParentID)
	if err != nil {
		return serializer.Error(serializer.ErrSQL)
	}
	return serializer.Success("ok")
}

// UpdateLevelService 更新level
type UpdateLevelService struct {
	ID       int    `json:"id"`
	ParentID int    `json:"parent_id"`
	Name     string `json:"name" binding:"required,min=2,max=64"`
	OrderID  int    `json:"order_id"`
	Comment  string `json:"comment,omitempty"`
}

// Update 更新level
func (service *UpdateLevelService) Update(id int) serializer.Response {
	// var id int
	// err := models.DB.Get(&id, "SELECT count(id) FROM level where name=?", service.Name)

	// if id != 0 {
	// 	str := service.Name + serializer.GetMessage(serializer.ErrSQLExist)
	// 	return serializer.ErrorStr(serializer.ErrSQLExist, str)
	// }
	sqlStr := `UPDATE level SET 
				   name = ?,
				   parent_id = ?,
				   order_id = ?
			   where 
			   		id = ?`
	_, err := models.DB.Exec(sqlStr, service.Name, service.ParentID, service.OrderID, id)
	if err != nil {
		return serializer.Error(serializer.ErrSQL)
	}
	return serializer.Success("ok")
}

// DeleteLevelService 删除level
type DeleteLevelService struct{}

// Delete 删除level
func (service *DeleteLevelService) Delete(id int) serializer.Response {
	var count int
	models.DB.Get(&count, "SELECT count(id) FROM level where id=?", id)

	if count == 0 {
		return serializer.Error(serializer.ErrSQLNotExist)
	}

	sqlStr := `DELETE FROM level where id = ?`
	res, err := models.DB.Exec(sqlStr, id)
	if err != nil {
		return serializer.Error(serializer.ErrSQL)
	}
	_, err = res.RowsAffected()
	if err != nil {
		return serializer.Error(serializer.ErrSQL)
	}
	return serializer.Success("ok")
}
