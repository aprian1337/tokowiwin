package db

import (
	"fmt"
	"reflect"
	"strings"
	"tokowiwin/utils/array"
)

func GetUpdateQuery(field interface{}, exclude ...string) string {
	var (
		query string
		q     = make([]string, 0)
	)

	if f := reflect.TypeOf(field); f.Kind() == reflect.Ptr {
		e := reflect.ValueOf(field).Elem()
		n := e.NumField()
		for i := 0; i < n; i++ {
			if !array.Contains(exclude, e.Type().Field(i).Tag.Get("db")) {
				switch e.Type().Field(i).Type.String() {
				case "sql.NullInt64":
					q = append(q, fmt.Sprintf("%v=%v", e.Type().Field(i).Tag.Get("db"), e.Field(i).FieldByName("Int64")))
				case "sql.NullInt32":
					q = append(q, fmt.Sprintf("%v=%v", e.Type().Field(i).Tag.Get("db"), e.Field(i).FieldByName("Int32")))
				case "sql.NullBool":
					q = append(q, fmt.Sprintf("%v=%v", e.Type().Field(i).Tag.Get("db"), e.Field(i).FieldByName("Bool")))
				case "string":
					q = append(q, fmt.Sprintf("%v='%v'", e.Type().Field(i).Tag.Get("db"), e.Field(i)))
				default:
					q = append(q, fmt.Sprintf("%v=%v", e.Type().Field(i).Tag.Get("db"), e.Field(i)))
				}
			}
		}
	}
	query = strings.Join(q, ",")

	return query
}
