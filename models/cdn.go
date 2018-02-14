package models

import (
	"time"
	// "github.com/astaxie/beego/orm"
)

//auto register
func init() {
	Register(new(Cdn), new(Machine))
	// orm.RegisterModel(new(Cdn), new(Machine))
}

type Cdn struct {
	Id       int64  `name:"id" search:"true"`
	Cdn_name string `name:"cdn_name" search:"true"`
	Type     string `name:"type" search:"false"`
}

type Machine struct {
	Id     int64     `xorm:"autoincr" name:"id" search:"true"`
	Sn     string    `xorm:"sn" name:"sn" search:"true"`
	Mac    string    `xorm:"mac" name:"mac" search:"true"`
	Ip     string    `xorm:"ip" name:"ip" search:"true"`
	Name   string    `xorm:"name" name:"name" search:"true"`
	Create time.Time `xorm:"created"` //这个Field将在Insert时自动赋值为当前时间
	Update time.Time `xorm:"updated"` //这个Field将在Insert或Update时自动赋值为当前时间
}
