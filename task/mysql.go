package main

import (
	"aurora/blog/api/common"
	_ "github.com/go-sql-driver/mysql"
)

// 初始化数据库，创建对应数据表
func main() {
	common.SetupConfig()
	common.SetupDB()
	common.SetupTable()
}
