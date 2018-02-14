package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/lflxp/beegoadmin/utils"
)

func init() {
	// vpn := Vpn{}s
	utils.Register(new(Vpn))
	orm.RegisterModel(new(Vpn))
}

type Vpn struct {
	Id   int64  `name:"id" search:"true"`
	Vpn  string `name:"vpn" search:"true"`
	Name string `name:"name" search:"false"`
	Ip   string `name:"ip" search:"false"`
}
