package db

import (
	"backend/models"
)

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

// 构建树的结构.
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

// GetLevelList 获取level列表
func GetLevelList() (list []*Levels, err error) {
	var levelList []*models.Level
	sqlStr := ` select id,created_at,updated_at,name,comment,order_id,parent_id
	from level`
	err = models.DB.Select(&levelList, sqlStr)
	if err != nil {
		return nil, err
	}

	for _, v := range levelList {
		var l Levels
		l.ID = v.ID
		l.Name = v.Name
		l.Comment = v.Comment
		l.ParentID = v.ParentID
		l.OrderID = v.OrderID
		list = append(list, &l)
	}
	data := buildData(list)

	list = buildTree(0, data)
	return list, nil
}

// CreateLevel 创建level
func CreateLevel(parentID int, orderID int, name string) (id int64, err error) {
	sqlStr := `
		INSERT INTO level(name, order_id, parent_id)
		VALUES (?,?,?)
	`
	res, err := models.DB.Exec(sqlStr, name, parentID, orderID)
	if err != nil {
		return 0, err
	}
	id, err = res.LastInsertId()
	return id, err
}

// DeleteLevel 删除level
func DeleteLevel(id int) (err error) {
	sqlStr := `DELETE FROM level where id = ?`
	res, err := models.DB.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

// UpdateLevel 更新
func UpdateLevel(id, parentID, orderID int, name string) (err error) {
	// 如何更新数据，比如第二级的数据插入第一级中，如何把所有的数据都更新了？
	// 方案1，根据修改时间，修改时间比较新的排前面，那get 请求也需要重新整理
	// 方案2，直接把ID靠后的全部取出来，全部修改，性能差
	return nil
}
