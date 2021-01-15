package model

import "aurora/blog/api/constant"

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

func (self *Category) Mapping() map[string]interface{} {
	return map[string]interface{}{
		"id":          self.ID,
		"name":        self.Name,
		"description": self.Description,
		"extra":       self.Extra,
		"createdAt":   self.CreatedAt.Format(constant.SimpleTimeFormat),
		"updatedAt":   self.UpdatedAt.Format(constant.SimpleTimeFormat),
	}
}
