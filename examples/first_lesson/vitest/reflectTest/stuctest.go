package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"monster_name"`
	Age   int    `json:"monster_age"`
	Score float64
}

func (m *Monster) Print() {
	fmt.Printf("方法print函数内打印: name: %s, age: %d, score: %f\n", m.Name, m.Age, m.Score)
}
func (m *Monster) Sayhello() {
	fmt.Printf("方法sayhello内部打印, %s!\n", m.Name)
}
func (m *Monster) Addsum(a int, b int) int {
	fmt.Printf("方法Addsum内部打印, %s!\n", m.Name)

	return a + b
}

func TestStruct(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println("valueof如传入进来的是地址，那么就是指向值的地址&｛xxx｝，如果是值就是结构体的值：\n", v)
	t1 := reflect.TypeOf(a)
	fmt.Println("typeof不管传入进来的是值还是地址都是main.Monster=", t1)
	var vv reflect.Value
	//fmt.Printf("valueof.kind=是个指针ptr=%v,v.elem()=%v\n", v.Kind(), v.Elem())
	//如果传入进来的不是地址，不能取v.elem
	fmt.Printf("valueof.kind=是个指针ptr，如果传入进来的不是指针那返回的＝struct：%v \n", v.Kind())
	// 如果是指针类型，通过 Elem() 获取指针指向的实际值
	if v.Kind() == reflect.Ptr {
		vv = v.Elem()
		fmt.Println("v.elem()＝指针所指的值=", vv)
	} else {
		vv = v
	}
	// 获取类型
	t := vv.Type()
	fmt.Println("vv.type如是值＝main.Monster：", t)
	kd := vv.Kind()
	fmt.Println("vv.kind如是值＝main.monster：", t)

	if kd != reflect.Struct {
		fmt.Println("not struct")
		return
	}

	//获取结构体的所有字段
	num := vv.NumField()
	v.Elem().Field(0).SetString("白骨精") //修改字段的值

	fmt.Printf("这个结构体有多少个字段: %d\n", num)
	//遍历结构体所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("field %d: 值为＝%v\n", i, vv.Field(i))
		//这里可以获取struct标签，reflect.type来获取tag标签的值
		tagval := t.Field(i).Tag.Get("json")
		if tagval != "" {
			fmt.Printf("fied %d tag: %v\n", i, tagval)
		}
	}

	//获取结构体所有的方法 , v是结构体的指针
	nummethod := v.NumMethod()
	//nummethod := reflect.ValueOf(a).NumMethod() //要以这个方式才能获取到方法个数
	fmt.Printf("传入的结构体指针总共有多少个方法: %d\n", nummethod)
	// 如果方法存在，调用第一个方法 Print()
	if nummethod > 0 {
		v.Method(1).Call(nil)
	}

	var param []reflect.Value
	param = append(param, reflect.ValueOf(10))
	param = append(param, reflect.ValueOf(20))
	//res := v.Method(0).Call(param)
	//如果传入进来的不是地址不能取方法？
	res := v.Method(0).Call(param) // 这里 Method(2) 调用 Addsum 方法
	fmt.Printf("结构体方法add method total: %d result:%v\n", len(res), res[0].Int())
}

func main() {
	var a Monster = Monster{
		Name:  "老鼠精",
		Age:   400,
		Score: 90.8,
	}
	TestStruct(&a)
}
