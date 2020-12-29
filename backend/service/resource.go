package service

import (
	"backend/models"
	"backend/serializer"
)

// ListResourceService 资源结构体
type ListResourceService struct {
}

// List 列表资源
func (service *ListResourceService) List() serializer.Response {
	var resourceList []*models.Resource
	sqlStr := ` select id,created_at,updated_at,image_name,dockerfile_url_path from resource`
	err := models.DB.Select(&resourceList, sqlStr)
	if err != nil {
		return serializer.Error(serializer.ErrSQL)
	}

	res := serializer.BuildResources(resourceList)

	return serializer.Success(res)
}

// CreateResource 创建资源
type CreateResource struct {
	ImageName         string  `json:"image_name" binding:"required"`
	DockerfileURLPath *string `json:"dockerfile_url_path"`
}

// Create 创建
func (service *CreateResource) Create() serializer.Response {

	var id int
	err := models.DB.Get(&id, "SELECT count(id) FROM resource where image_name=?", service.ImageName)
	if err != nil {
		return serializer.Error(serializer.ErrSQL)
	}

	if id != 0 {
		str := service.ImageName + serializer.GetMessage(serializer.ErrSQLExist)
		return serializer.ErrorStr(serializer.ErrSQLExist, str)
	}

	sqlStr := `INSERT INTO resource(image_name, dockerfile_url_path)
		VALUES (?,?)`
	_, err = models.DB.Exec(sqlStr, service.ImageName, service.DockerfileURLPath)
	if err != nil {
		return serializer.Error(serializer.ErrSQL)
	}

	return serializer.Success("ok")
}
