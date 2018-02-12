package models

type Cdn struct {
	Id       int64  `name:"id" search:"true"`
	Cdn_name string `name:"cdn_name" search:"true"`
	Type     string `name:"type" search:"false"`
}

type Machine struct {
	Id   int64  `name:"id" search:"true"`
	Sn   string `name:"sn" search:"true"`
	Mac  string `name:"mac" search:"true"`
	Ip   string `name:"ip" search:"true"`
	Name string `name:"name" search:"true"`
}
