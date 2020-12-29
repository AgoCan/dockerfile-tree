package serializer

import (
	"backend/models"
)

// Configs 配置表
type Configs struct {
	ID        int     `json:"id"`
	Key       string  `json:"config_key"`
	Value     *string `json:"config_value"`
	Comment   *string `json:"config_comment"`
	CreatedAt int64   `json:"created_at"`
}

// BuildConfigs 返回配置列表
func BuildConfigs(configs []*models.Config) (list []*Configs) {
	for _, v := range configs {
		var l Configs
		l.ID = v.ID
		l.Key = v.Key
		l.Value = v.Value
		l.Comment = v.Comment
		l.CreatedAt = v.CreatedAt.Unix()
		list = append(list, &l)
	}
	return list
}
