package model

// 标签信息
type Tag struct {
	Model

	Name  string `json:"name"`  // 标签名称
	Extra string `json:"extra"` // 拓展字段
}

func (Tag) TableName() string {
	return "tag"
}
