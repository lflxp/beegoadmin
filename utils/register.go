package utils

func Register(data ...interface{}) func(...interface{}) []map[string]string {
	return ReadStruct(data...)
}

func ReadStruct(data ...interface{}) []map[string]string {
	return []map[string]string{}
}
