package model

import "aurora/blog/api/constant"

// 作者信息
type Author struct {
	Model

	Name     string `json:"name"`     // 名称
	NickName string `json:"nickName"` // 昵称
	Age      int    `json:"age"`      // 年龄
	Gender   int    `json:"gender"`   // 性别
	Hobby    string `json:"hobby"`    // 爱好
	Avatar   string `json:"avatar"`   // 头像地址
	Extra    string `json:"extra"`    // 拓展字段
}

func (Author) TableName() string {
	return "author"
}

func (self *Author) Mapping() map[string]interface{} {
	return map[string]interface{}{
		"name":      self.Name,
		"nickName":  self.NickName,
		"age":       self.Age,
		"gender":    self.Gender,
		"hobby":     self.Hobby,
		"avatar":    self.Avatar,
		"extra":     self.Extra,
		"createdAt": self.CreatedAt.Format(constant.SimpleTimeFormat),
		"updatedAt": self.UpdatedAt.Format(constant.SimpleTimeFormat),
	}
}
