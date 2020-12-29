package serializer

import "backend/models"

// Dockerfile 序列化
type Dockerfile struct {
	Dockerfile string `json:"dockerfile"`
	LevelID    int    `json:"level_id"`
	CreatedAt  int64  `json:"created_at"`
}

//BuildDockerfile 测试序列化器
func BuildDockerfile(dockerfile models.Dockerfile) Dockerfile {
	return Dockerfile{
		Dockerfile: dockerfile.Dockerfile,
		LevelID:    dockerfile.LevelID,
		CreatedAt:  dockerfile.CreatedAt.Unix(),
	}
}
