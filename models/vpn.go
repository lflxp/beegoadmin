package models

import "time"

func init() {
	// vpn := Vpn{}s
	Register(new(Vpn), new(Machine), new(Cdn))
	// orm.RegisterModel(new(Vpn))
}

type Vpn struct {
	Id   int64  `name:"id" search:"true"`
	Vpn  string `name:"vpn" search:"true"`
	Name string `name:"name" search:"false"`
	Ip   string `name:"ip" search:"false"`
}

type Cdn struct {
	Id       int64  `name:"id" search:"true"`
	Cdn_name string `name:"cdn_name" search:"true"`
	Type     string `name:"type" search:"false"`
}

type Machine struct {
	Id     int64     `name:"id" search:"true"`
	Sn     string    `xorm:"sn" name:"sn" search:"true"`
	Mac    string    `xorm:"mac" name:"mac" search:"true"`
	Ip     string    `xorm:"ip" name:"ip" search:"true"`
	Name   string    `xorm:"name" name:"name" search:"true"`
	Create time.Time `xorm:"created"` //这个Field将在Insert时自动赋值为当前时间
	Update time.Time `xorm:"updated"` //这个Field将在Insert或Update时自动赋值为当前时间
}
