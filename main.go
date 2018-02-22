package main

import (
	"github.com/astaxie/beego"
	_ "github.com/lflxp/beegoadmin/routers"
	_ "github.com/lflxp/beegoadmin/utils"
	_ "github.com/mattn/go-sqlite3"
)

// func init() {
// 	orm.RegisterDriver("sqlite", orm.DRSqlite)
// 	// 参数4(可选)  设置最大空闲连接
// 	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
// 	maxIdle := 30
// 	maxConn := 30
// 	orm.RegisterDataBase("default", "sqlite3", "./db.sqlite3333", maxIdle, maxConn)
// }

func main() {
	beego.Run()
}
