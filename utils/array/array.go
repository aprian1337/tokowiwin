package array

import (
	"reflect"
	"strconv"
	"strings"
)

func ToStringWithDelimiter(arr []any, delimiter string) string {
	var res string
	if reflect.TypeOf(arr).Kind() != reflect.Array && reflect.TypeOf(arr).Kind() != reflect.Slice {
		return ""
	}
	var tempArrStr []string
	for i := 0; i < reflect.ValueOf(arr).Len(); i++ {
		switch arr[i].(type) {
		case string:
			tempArrStr = append(tempArrStr, arr[i].(string))
		case int, int64:
			tempArrStr = append(tempArrStr, strconv.Itoa(arr[i].(int)))
		case bool:
			tempArrStr = append(tempArrStr, strconv.FormatBool(arr[i].(bool)))
		}
	}
	res = strings.Join(tempArrStr, delimiter)
	return res
}

func Int64ToString(arr []int64) []string {
	var res []string
	for _, v := range arr {
		str := strconv.FormatInt(v, 10)
		res = append(res, str)
	}
	return res
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
