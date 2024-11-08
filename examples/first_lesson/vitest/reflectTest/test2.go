package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{}) {

	rval := reflect.ValueOf(b)
	if rval.Kind() != reflect.Ptr {
		fmt.Println("传入的参数不是指针，无法修改其值")
		return
	}
	fmt.Printf("rval = %v\n", rval)

	//rval.Elem().SetInt(200)
	rval = rval.Elem()

	// 判断是否是int类型
	// 如果是结构体类型，尝试修改它的字段
	if rval.Kind() == reflect.Struct {
		// 获取 "Age" 字段
		field := rval.FieldByName("Age")
		if field.IsValid() && field.CanSet() && field.Kind() == reflect.Int {
			fmt.Printf("修改前的 age 值: %d\n", field.Int())
			field.SetInt(200) // 修改 age 的值
			fmt.Printf("修改后的 age 值: %d\n", field.Int())
		} else {
			fmt.Println("无法修改字段 'Age'")
		}
	} else if rval.Kind() == reflect.Int {
		// 如果是 int 类型，直接修改
		fmt.Printf("修改前的值: %d\n", rval.Int())
		rval.SetInt(200)
		fmt.Printf("修改后的值: %d\n", rval.Int())
	} else {
		fmt.Println("当前只支持修改 int 类型和 struct 的字段")
	}
}

type student struct {
	name string
	age  int
}

func main() {
	var num int = 100
	reflectTest(&num)
	fmt.Println(num)
	stu1 := student{
		name: "wang",
		age:  20,
	}

	reflectTest(&stu1)
}
