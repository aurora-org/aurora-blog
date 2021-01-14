package model

// 文章分类信息
type Category struct {
	Model

	Name        string `json:"name"`        // 分类名称
	Description string `json:"description"` // 分类描述
	Extra       string `json:"extra"`       // 拓展字段
}

func (Category) TableName() string {
	return "category"
}
