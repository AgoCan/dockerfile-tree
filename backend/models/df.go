package models

// 构建 状态
const (
	BuildFailed = iota
	BuildSuccess
)

// 推送 状态
const (
	PushFailed = iota
	PushSuccess
)

// Dockerfile dockerfile
type Dockerfile struct {
	baseModel
	Dockerfile string `db:"dockerfile" json:"dockerfile"`
	LevelID    int    `db:"level_id" json:"level_id"`
}

// Level 层级表
type Level struct {
	baseModel
	Name     string  `db:"name" json:"name"`
	Comment  *string `db:"comment" json:"comment"`
	OrderID  int     `db:"order_id" json:"order_id"`
	ParentID int     `db:"parent_id" json:"parent_id"`
}

// LevelCombinationTask 关联表
type LevelCombinationTask struct {
	baseModel
	RecordDockerfile  string `db:"record_dockerfile" json:"record_dockerfile"`
	CombinationTaskID int    `db:"combination_task_id" json:"combination_task_id"`
	LevelID           int    `db:"level_id" json:"level_id"`
}

// CombinationTask 组合任务表
type CombinationTask struct {
	baseModel
	LastLevel   string `db:"last_level" json:"last_level"`
	BuildStatus int    `db:"build_status" json:"build_status"`
}

// Record 组合记录表
type Record struct {
	baseModel
	ImageInfo         string `db:"image_info" json:"image_info"`
	ImageName         string `db:"image_name" json:"image_name"`
	PushStatus        int    `db:"push_status" json:"push_status"`
	CombinationTaskID int    `db:"combination_task_id" json:"combination_task_id"`
}

// Resource 资源表述表
type Resource struct {
	baseModel
	ImageName         string  `db:"image_name" json:"image_name"`
	DockerfileURLPath *string `db:"dockerfile_url_path" json:"dockerfile_url_path"`
}

// Config 配置表
type Config struct {
	baseModel
	Key   string `db:"key" json:"key"`
	Value string `db:"value" json:"value"`
}
