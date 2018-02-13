package utils

import (
	"fmt"
	"reflect"
	"strings"
)

var Registered []map[string]string

func init() {
	Registered = []map[string]string{}
}

func RigsterStruct(data interface{}) map[string]string {
	tmp := map[string]string{}
	v := reflect.ValueOf(data)

	tmp["Struct"] = v.Type().String()
	for i := 0; i < v.NumField(); i++ {
		//利用反射获取structTag
		tmp["Tag"] = fmt.Sprintf("%s", v.Type().Field(i).Tag)
		tmp["Type"] = fmt.Sprintf("%s", v.Type().Field(i).Type)
		tmp["Name"] = fmt.Sprintf("%s %s", tmp["Name"], v.Type().Field(i).Name)
		// fmt.Println("Tag", v.Type().Field(i).Tag)
		// fmt.Println("Tag", v.Type().Field(i).Tag)
		// fmt.Println("Type", v.Type().Field(i).Type)
		// fmt.Println("Name", v.Type().Field(i).Name)
		// fmt.Println("Index", v.Type().Field(i).Index)
		// fmt.Println("Offset", v.Type().Field(i).Offset)
		// fmt.Println("pkgpath", v.Type().Field(i).PkgPath)
		// fmt.Println(st.Field(i).Tag.Get("search"))
		for _, y := range strings.Split(tmp["Tag"], " ") {
			rs := strings.Split(y, ":")
			tmp[v.Type().Field(i).Name+"_"+rs[0]] = rs[1]
		}
	}
	return tmp
}

func Register(data interface{}) error {
	Registered = append(Registered, RigsterStruct(data))
	return nil
}

// func ReadStruct(data ...interface{}) []map[string]string {
// 	result := []map[string]string{}
// 	for _, x := range data {
// 		result = append(result, RigsterStruct(x))
// 	}
// 	return result
// }
