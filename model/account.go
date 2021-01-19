package model

type Account struct {
	Model

	UserName string `json:"userName"` // 用户名
	Password string `json:"password"` // 密码
}

func (Account) TableName() string {
	return "account"
}
