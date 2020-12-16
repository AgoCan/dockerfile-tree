package db

// Levels level 树结构
type Levels struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Comment  *string   `json:"comment"`
	ParentID int       `json:"parent_id"`
	OrderID  int       `json:"order_id"`
	Children []*Levels `json:"children,omitempty"`
}
