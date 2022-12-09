package structs

import (
	"reflect"
	"tokowiwin/utils/array"
)

func GetColumns(field interface{}) []any {
	columns := make([]any, 0)
	if f := reflect.TypeOf(field); f.Kind() == reflect.Ptr {
		e := reflect.ValueOf(field).Elem()
		n := e.NumField()
		for i := 0; i < n; i++ {
			if e.Type().Field(i).Tag.Get("db") == "id" && e.Type().Field(i).Tag.Get("autoinc") == "true" {
				continue
			}
			columns = append(columns, e.Type().Field(i).Tag.Get("db"))
		}
	} else if f := reflect.TypeOf(field); f.Kind() == reflect.Slice || f.Kind() == reflect.Array {
		z := reflect.TypeOf(field).Elem().Kind()
		if f := reflect.TypeOf(z); f.Kind() == reflect.Ptr {
			e := reflect.ValueOf(z).Elem()
			n := e.NumField()
			for i := 0; i < n; i++ {
				if e.Type().Field(i).Tag.Get("db") == "id" && e.Type().Field(i).Tag.Get("autoinc") == "true" {
					continue
				}
				columns = append(columns, e.Type().Field(i).Tag.Get("db"))
			}
		}
	}
	return columns
}

func GetAddress(field interface{}) []any {
	columns := make([]any, 0)
	if f := reflect.TypeOf(field); f.Kind() == reflect.Ptr {
		e := reflect.ValueOf(field).Elem()
		n := e.NumField()
		for i := 0; i < n; i++ {
			fi := e.Field(i)
			columns = append(columns, fi.Addr().Interface())
		}
	}

	return columns
}

func GetAddressByFieldTag(field any, tag string, includes []string) []any {
	columns := make([]any, 0)
	if f := reflect.TypeOf(field); f.Kind() == reflect.Ptr {
		e := reflect.ValueOf(field).Elem()
		n := e.NumField()
		for i := 0; i < n; i++ {
			fi := e.Field(i)
			if !array.Contains(includes, e.Type().Field(i).Tag.Get(tag)) {
				continue
			}
			columns = append(columns, fi.Addr().Interface())
		}
	}

	return columns
}
