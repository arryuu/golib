package goutil

import (
	"reflect"
)

/**
 * 循环结构体里所有的方法
 *
 * valueOf.Type().Method(i).Name 方法名称
 * valueOf.Method(i).Interface().(func()) 方法内容
 */
func EachStructMethod(st interface{}, fn func(reflect.Value, int)) {
	valueOf := reflect.ValueOf(st)
	if valueOf.IsValid() {
		for i := 0; i < valueOf.NumMethod(); i++ {
			fn(valueOf, i)
		}
	}
}
