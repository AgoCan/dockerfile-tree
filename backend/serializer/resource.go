package serializer

import (
	"backend/models"
)

// Resources 资源表
type Resources struct {
	ID                int     `json:"id"`
	ImageName         string  `json:"image_name"`
	DockerfileURLPath *string `json:"dockerfile_url_path"`
	CreatedAt         int64   `json:"created_at"`
}

// BuildResources 返回资源列表
func BuildResources(resources []*models.Resource) (list []*Resources) {
	for _, v := range resources {
		var l Resources
		l.ID = v.ID
		l.ImageName = v.ImageName
		l.DockerfileURLPath = v.DockerfileURLPath
		l.CreatedAt = v.CreatedAt.Unix()
		list = append(list, &l)
	}
	return list
}
