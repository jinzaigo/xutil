package xutil

import (
	"reflect"
	"strings"
)

type IStruct interface {
	GetStructData() interface{}
}

//StructToMap
//使用反射实现，完美地兼容了json标签的处理
func StructToMap(st IStruct) map[string]interface{} {
	m := make(map[string]interface{})
	in := st.GetStructData()
	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return m
	}

	relType := val.Type()
	for i := 0; i < relType.NumField(); i++ {
		name := relType.Field(i).Name
		tag := relType.Field(i).Tag.Get("json")
		if tag != "" {
			index := strings.Index(tag, ",")
			if index == -1 {
				name = tag
			} else {
				name = tag[:index]
			}
		}
		m[name] = val.Field(i).Interface()
	}
	return m
}
