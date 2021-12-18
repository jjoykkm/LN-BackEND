package obsolete_utility

import (
	"reflect"
	"strings"
)

func RemoveDuplicateValues(listData interface{}) interface{} {
	source := reflect.ValueOf(listData)
	list := reflect.MakeSlice(source.Type(), 0, 0)
	visited := make(map[interface{}]struct{})
	for i := 0; i < source.Len(); i++ {
		value := source.Index(i)
		if _, ok := visited[value.Interface()]; ok {
			continue
		}
		visited[value.Interface()] = struct{}{}
		list = reflect.Append(list, value)
	}
	return list.Interface()
}

func ConvertListToStringIn(listData []string) string {
	result := "('" + strings.Join(listData, "','") + "')"
	return result
}
