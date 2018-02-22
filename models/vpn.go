package models

import (
	"time"
)

func init() {
	// vpn := Vpn{}s
	Register(new(Vpn), new(Machine), new(Cdn), new(More))
}

/*
name  字段名
verbose_name 标识
list_display 显示字段
search_fields 查询字段
manytomany 一对多字段 指定表明
colType 字段类型 -> string|int|file|textarea|radio|m2m|otm|o2o|time|select|multiselect
radio|select -> Name|value,Name|value,...
*/
type Vpn struct {
	Id   int64
	Vpn  string `xorm:"vpn" name:"vpn" verbose_name:"Vpn字段测试" list:"true" search:"true"`
	Name string `xorm:"name" name:"name" verbose_name:"姓名" list:"true" search:"false"`
	Ip   string `xorm:"ip" name:"ip" verbose_name:"ip信息" list:"true" search:"false"`
}

type Cdn struct {
	Id           int64  `xorm:"id" name:"id" search:"true"`
	Cdn_name     string `xorm:"cdn_name" name:"cdn_name" verbose_name:"cdn的名称" search:"true"`
	Type         string `xorm:"type" name:"vpn" verbose_name:"类型" search:"false" colType:"textarea"`
	Detail       string `xorm:"detail" name:"vpn" verbose_name:"VPN信息" list:"true" search:"false" manytomany:"vpn" colType:"o2m"`
	Radio        string `xorm:"raidodas" name:"radio23" verbose_name:"Radio单选" list:"true" search:"false" colType:"radio" radio:"男|man,女|girl,人妖|none"`
	Select       string `xorm:"ss" name:"rrad" verbose_name:"Select单选固定" list:"true" search:"false" colType:"select" select:"男11111111111111111111111111|man,女|girl,人妖|none"`
	MultiSelect  string `xorm:"ss1" name:"rrad1" verbose_name:"Multiselect多选" list:"true" search:"false" colType:"multiselect" multiselect:"男|man,女|girl,人妖|none,中|zhong,国|guo,人|ren,重|chong,Qing|qing"`
	MultiSelect2 string `xorm:"ss1" name:"rrad1" verbose_name:"Multiselect多选" list:"true" search:"false" colType:"multiselect" multiselect:"男|man,女|girl,人妖|none,中|zhong,国|guo,人|ren,重|chong,Qing|qing"`
	Files        string `xorm:"cdn_name" name:"cdn_name" verbose_name:"cdn的名称" search:"true" colType:"file"`
}

type Machine struct {
	Id     int64     `xorm:"id" name:"id" search:"true"`
	Sn     string    `xorm:"sn" name:"sn" search:"true"`
	Mac    string    `xorm:"mac" name:"mac" search:"true"`
	Ip     string    `xorm:"ip" name:"ip" search:"true"`
	Name   string    `xorm:"name" name:"name" search:"true"`
	Create time.Time `xorm:"created"` //这个Field将在Insert时自动赋值为当前时间
	Update time.Time `xorm:"updated"` //这个Field将在Insert或Update时自动赋值为当前时间
}

type More struct {
	Uid      int64  `xorm:"id pk not null autoincr"`
	Username string `xorm:"unique"`
	Alias    string `xorm:"-"`
	Vpn      `xorm:"vpn_id int(11)"`
}
