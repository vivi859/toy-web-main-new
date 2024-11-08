package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{}) {
	//获取reflect.type类型b的类型
	rtype := reflect.TypeOf(b)
	//获取reflect.value类型
	rval := reflect.ValueOf(b)
	//获取到b对应的类别
	//kd := rtype.Kind()
	//获取是获取指针指向的实际值。　rval.Elem() 只有在 b 是指针时才可以使用
	elemb := rtype.Elem()
	fmt.Printf("b.type= %T ,rtype.kind=%v,rval.kind=%v,elemb.kind=%v\n", reflect.TypeOf(rtype), rtype.Kind(), rval.Kind(), elemb.Kind())
	rval1 := rval.Elem().Interface() //一定要是elem
	rval2, ok := rval1.(float64)
	if ok {
		fmt.Printf("类型断言float64:", rval2)
	}
}

func main() {
	var num float64
	num = 0.111
	reflectTest(&num) //一定要是地址
}
