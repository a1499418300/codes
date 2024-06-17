package main

import (
	"fmt"
	"reflect"
	"time"
)

type MyStruct struct {
	Field1 string
	Field2 int
	Field3 float64
}

func main() {
	var i interface{} = MyStruct{"Hello, World!", 42, 3.14}
	begin := time.Now()
	fmt.Println(fmt.Sprintf("耗时：%v, 耗时： %s", time.Since(begin.Add(-time.Hour*24)), time.Since(begin.Add(-time.Hour))))
	v := reflect.ValueOf(i)
	// 检查我们是否有一个结构体
	if v.Kind() == reflect.Struct {
		// 遍历结构体的字段
		for i := 0; i < v.NumField(); i++ {
			// 获取结构体字段的反射类型对象
			fieldType := v.Type().Field(i)
			// 输出字段名和值
			fmt.Printf("Name: %s, Value: %v\n", fieldType.Name, v.Field(i).Interface())
		}
	}
}
