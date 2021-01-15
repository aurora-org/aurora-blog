package model

import (
	"aurora/blog/api/constant"
	"time"
)

// 站点信息
type Site struct {
	Model

	Name    string    `json:"name"`    // 网站名称
	BeiAn   string    `json:"bei_an"`  // 备案号
	Powered string    `json:"powered"` // 站点引擎
	StartAt time.Time `json:"startAt"` // 站点启动日期
	Extra   string    `json:"extra"`   // 拓展字段
}

func (Site) TableName() string {
	return "site"
}

func (self *Site) Mapping() map[string]interface{} {
	return map[string]interface{}{
		"name":      self.Name,
		"beiAn":     self.BeiAn,
		"powered":   self.Powered,
		"startAt":   self.StartAt,
		"extra":     self.Extra,
		"createdAt": self.CreatedAt.Format(constant.SimpleTimeFormat),
		"updatedAt": self.UpdatedAt.Format(constant.SimpleTimeFormat),
	}
}
