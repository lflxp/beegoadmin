package models

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
)

var Registered []map[string]string

func init() {
	Registered = []map[string]string{}
}

func GetRegisterByName(name string) map[string]string {
	for _, maps := range Registered {
		if strings.ToUpper(maps["Struct"]) == strings.ToUpper(name) {
			return maps
		}
	}
	return nil
}

//http://code.lflxp.cn/life/services/src/master/util/db/xormTest.go
//收集struct信息
func RigsterStruct(data interface{}) map[string]string {
	tmp := map[string]string{}
	// ttt := reflect.TypeOf(data)
	// beego.Critical(ttt.Elem())
	// beego.Critical(ttt.NumMethod())
	// beego.Critical(ttt.MethodByName().Func.)

	vv := reflect.ValueOf(data)
	v := reflect.Indirect(vv)

	// beego.Critical("API", v.Type().Name(), v.Type().PkgPath())
	// vv.MethodByName().FieldByName()
	// vv.MethodByName().IsValid

	tmp["Struct"] = v.Type().Name()
	for i := 0; i < v.NumField(); i++ {
		//利用反射获取structTag
		tmp["Tag"] = fmt.Sprintf("%s", v.Type().Field(i).Tag)
		// html 字段类型
		if v.Type().Field(i).Tag.Get("colType") != "" {
			tmp["ColType"] = v.Type().Field(i).Tag.Get("colType")
		} else {
			tmp["ColType"] = v.Type().Field(i).Type.String()
		}
		beego.Critical("ColType", tmp["ColType"])
		switch tmp["ColType"] {
		case "radio":
			tmp["detail"] = v.Type().Field(i).Tag.Get("radio")
		case "manytomany":
			tmp["detail"] = v.Type().Field(i).Tag.Get("manytomany")
		case "select":
			tmp["detail"] = v.Type().Field(i).Tag.Get("select")
		case "multiselect":
			tmp["detail"] = v.Type().Field(i).Tag.Get("multiselect")
		}

		//字段名 供前端form表单使用
		if v.Type().Field(i).Tag.Get("name") != "" {
			tmp["Name"] = v.Type().Field(i).Tag.Get("name")
		} else {
			tmp["Name"] = v.Type().Field(i).Name
		}
		//收集昵称
		if v.Type().Field(i).Tag.Get("verbose_name") != "" {
			tmp["Col"] = fmt.Sprintf("%s %s:%s:%s:%s", tmp["Col"], v.Type().Field(i).Tag.Get("verbose_name"), tmp["ColType"], tmp["Name"], tmp["detail"])
		} else {
			tmp["Col"] = fmt.Sprintf("%s %s:%s:%s:%s", tmp["Col"], v.Type().Field(i).Name, tmp["ColType"], tmp["Name"], tmp["detail"])
		}

		//

		// fmt.Println("Tag", v.Type().Field(i).Tag)
		// fmt.Println("Tag", v.Type().Field(i).Tag)
		// fmt.Println("Type", v.Type().Field(i).Type)
		// fmt.Println("Name", v.Type().Field(i).Name)
		// fmt.Println("Index", v.Type().Field(i).Index)
		// fmt.Println("Offset", v.Type().Field(i).Offset)
		// fmt.Println("pkgpath", v.Type().Field(i).PkgPath)
		// fmt.Println(st.Field(i).Tag.Get("search"))
		// for _, y := range strings.Split(tmp["Tag"], " ") {
		// 	rs := strings.Split(y, ":")
		// 	tmp[v.Type().Field(i).Name+"_"+rs[0]] = rs[1]
		// }
		fmt.Println("DDDDDDDDDDD", v.Type().Field(i).Tag.Get("verbose_name"))
	}
	return tmp
}

func Register(data ...interface{}) error {
	for _, model := range data {
		Registered = append(Registered, RigsterStruct(model))
	}
	return nil
}

// func ReadStruct(data ...interface{}) []map[string]string {
// 	result := []map[string]string{}
// 	for _, x := range data {
// 		result = append(result, RigsterStruct(x))
// 	}
// 	return result
// }
