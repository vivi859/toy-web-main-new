package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{}) {

	//通过反射获取传入变量的type kind value ,
	//1第一步先获取reflect.Type　.value
	rtype := reflect.TypeOf(b)
	fmt.Println(rtype)
	fmt.Printf("b type=%v,rtype =%T \n", rtype, rtype)
	rval := reflect.ValueOf(b)
	//n2 := 2 + rval.Int()
	//fmt.Println(n2)
	fmt.Printf("b value=%v,rval =%T \n", rval, rval)
	// 将rval又转回成interface{}
	iv := rval.Interface()
	fmt.Printf("newnum=%v,newnum=%T \n", iv, iv)

	//断言 要判断是Int类型还是student类型，怎么写呢．
	newnum, ok := iv.(int)
	if ok {
		fmt.Printf("断言成功　newnum=%v,newnum=%T \n", newnum, newnum)
	}

	//for 循环断言
	// 使用类型断言判断是 int 类型还是 student 类型
	switch v := iv.(type) {
	case int:
		fmt.Printf("断言成功，类型是 int，值为：%v\n", v)
		v2 := 2 + v
		fmt.Printf("v2=%v v2=%T \n", v2, v2)
	case student:
		fmt.Printf("断言成功，类型是 student，值为：%v，名字：%s，年龄：%d\n", v, v.name, v.age)
	default:
		fmt.Printf("未知类型：%T\n", v)
	}
}

type student struct {
	name string
	age  int
}

func main() {
	var num int = 100
	reflectTest(num)
	stu1 := student{
		name: "xiaoxiao",
		age:  19,
	}
	reflectTest(stu1)
}
