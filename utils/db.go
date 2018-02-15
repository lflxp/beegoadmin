package utils

/*
import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "sqlite3", "./sqlite.db", maxIdle, maxConn)
}
*/

import (
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/lflxp/beegoadmin/models"
	_ "github.com/mattn/go-sqlite3"
)

var (
	Engine *xorm.Engine
	err    error
)

func init() {
	Engine, err = xorm.NewEngine("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}
	Engine.ShowSQL(true)
	Engine.Logger().SetLevel(core.LOG_DEBUG)
	Engine.SetMaxIdleConns(300)
	Engine.SetMaxOpenConns(300)
	Engine.SetMapper(core.SnakeMapper{})
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "admin_")
	Engine.SetTableMapper(tbMapper)
	Engine.SetColumnMapper(core.SameMapper{})

	err = Engine.Sync2(new(models.Machine), new(models.Vpn), new(models.Cdn))
	if err != nil {
		panic(err)
	}
}
