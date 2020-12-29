package serializer

import "backend/models"

// Levels level 树结构
type Levels struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Comment   *string   `json:"comment"`
	ParentID  int       `json:"parent_id"`
	OrderID   int       `json:"order_id"`
	CreatedAt int64     `json:"created_at"`
	Children  []*Levels `json:"children,omitempty"`
}

// 构建parent_id为key的map结构.
func buildData(list []*Levels) map[int]map[int]*Levels {
	data := make(map[int]map[int]*Levels)
	for _, v := range list {

		id := v.ID                   // 主ID
		pid := v.ParentID            // 父ID
		if _, ok := data[pid]; !ok { // 如果不存在则创建一个新节点
			data[pid] = make(map[int]*Levels)
		}
		data[pid][id] = v

	}
	return data
}

// BuildTree 构建树的结构.
// a. 判断parent_id是否存在.
// b. 如果parent_id存在继续递归.至到data没有找到parent_id节点的数据.
func buildTree(parentID int, data map[int]map[int]*Levels) []*Levels {
	node := make([]*Levels, 0)
	for id, item := range data[parentID] {
		if data[id] != nil {
			item.Children = buildTree(id, data)
		}
		node = append(node, item)
	}
	return node
}

//BuildLevels 测试序列化器
func BuildLevels(levels []*models.Level) (list []*Levels) {
	for _, v := range levels {
		var l Levels
		l.ID = v.ID
		l.Name = v.Name
		l.Comment = v.Comment
		l.CreatedAt = v.CreatedAt.Unix()
		l.ParentID = v.ParentID
		l.OrderID = v.OrderID
		list = append(list, &l)
	}
	data := buildData(list)

	list = buildTree(0, data)
	return list
}
