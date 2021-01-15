package model

import "aurora/blog/api/constant"

// 标签信息
type Tag struct {
	Model

	Name  string `json:"name"`  // 标签名称
	Extra string `json:"extra"` // 拓展字段
}

func (Tag) TableName() string {
	return "tag"
}

func (self *Tag) Mapping() map[string]interface{} {
	return map[string]interface{}{
		"id":        self.ID,
		"name":      self.Name,
		"extra":     self.Extra,
		"createdAt": self.CreatedAt.Format(constant.SimpleTimeFormat),
		"updatedAt": self.UpdatedAt.Format(constant.SimpleTimeFormat),
	}
}
