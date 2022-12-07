package converts

import (
	"reflect"
	"strconv"
)

func AnyArrayToString(a []any) []string {
	res := make([]string, 0)
	for _, v := range a {
		var temp string
		switch reflect.TypeOf(v).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			temp = strconv.FormatInt(reflect.ValueOf(v).Elem().Int(), 10)
		case reflect.String:
			temp = v.(string)
		}
		res = append(res, temp)
	}
	return res
}
