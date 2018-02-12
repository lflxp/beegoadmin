package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func Register(data ...interface{}) []map[string]string {
	return ReadStruct(data)
}

func ReadStruct(data ...interface{}) []map[string]string {
	result := []map[string]string{}
	for _, x := range data {
		tmp := map[string]string{}
		st := reflect.TypeOf(&x).Elem()
		tmp["Struct"] = st.String()
		for i := 0; i < st.NumField(); i++ {
			tmp["Tag"] = fmt.Sprintf("%s", st.Field(i).Tag)
			tmp["Type"] = fmt.Sprintf("%s", st.Field(i).Type)
			tmp["Name"] = st.Field(i).Name
			for _, y := range strings.Split(tmp["Tag"], " ") {
				rs := strings.Split(y, ":")
				tmp[st.Field(i).Name+"_"+rs[0]] = rs[1]
			}
		}
		result = append(result, tmp)
	}
	return result
}
