package model

import "time"

// 模型基础结构
type Model struct {
	ID        int       `json:"id"`        // 主键ID
	CreatedAt time.Time `json:"createdAt"` // 创建日期
	UpdatedAt time.Time `json:"updatedAt"` // 更新日期
}
