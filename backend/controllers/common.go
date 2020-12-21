package controllers

// LevelInfo create level
type LevelInfo struct {
	ID       int    `json:"id"`
	ParentID int    `json:"parent_id"`
	Name     string `json:"name"`
	OrderID  int    `json:"order_id"`
	Comment  string `json:"comment,omitempty"`
}

// DelLevelInfo 删除信息
type DelLevelInfo struct {
	ID int `json:"id"`
}

// DockerfileInfo create DockerfileInfo
type DockerfileInfo struct {
	Dockerfile string `json:"dockerfile"`
	LevelID    int    `json:"level_id"`
}
