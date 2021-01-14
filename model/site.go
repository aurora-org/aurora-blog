package model

import "time"

// 站点信息
type Site struct {
	Model

	Name    string    `json:"name"`    // 网站名称
	BeiAn   string    `json:"bei_an"`  // 备案号
	Powered string    `json:"powered"` // 站点引擎
	StartAt time.Time `json:"startAt"` // 站点启动日期
}

func (Site) TableName() string {
	return "site"
}
