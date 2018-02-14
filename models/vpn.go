package models

func init() {
	// vpn := Vpn{}s
	Register(new(Vpn))
	// orm.RegisterModel(new(Vpn))
}

type Vpn struct {
	Id   int64  `name:"id" search:"true"`
	Vpn  string `name:"vpn" search:"true"`
	Name string `name:"name" search:"false"`
	Ip   string `name:"ip" search:"false"`
}
