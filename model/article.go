package model

// 文章信息
type Article struct {
	Model

	Title    string `json:"title"`                    // 文章标题
	Content  string `json:"content" gorm:"type:text"` // 文章内容
	Banner   string `json:"banner"`                   // 文章主图
	Author   string `json:"author"`                   // 文章作者
	Tag      string `json:"tag"`                      // 文章标签
	Category string `json:"category"`                 // 文章分类
	Times    int    `json:"times"`                    // 阅读次数
	Visible  bool   `json:"visible"`                  // 是否可见
	Extra    string `json:"extra"`                    // 拓展字段
}

func (Article) TableName() string {
	return "article"
}
