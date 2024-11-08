package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Userid string `json:userid`
	Name   string
}

func main() {

	//TestReflectTest()

	var (
		model *User
		st    reflect.Type
		elem  reflect.Value
	)
	st = reflect.TypeOf(model)                                       //获取类型*user 是个指针
	fmt.Println("reflect.TYpeof", st.Kind().String())                //ptr
	st = st.Elem()                                                   //指向的类型
	fmt.Println("reflect.of.eleme", st.Kind().String())              //struct
	elem = reflect.New(st)                                           //返回一个values类型值，该值持有一个指向类型为type的新申请的零值的指针
	fmt.Println("reflect.elem===", elem)                             //&{}空指针
	fmt.Println("reflect.new.of", elem.Kind().String())              //ptr
	fmt.Println("reflect.new.of.eleme", elem.Elem().Kind().String()) //struct
	model = elem.Interface().(*User)                                 //model赋值空接口类型断言＊user类型，model是*user它的指向跟elem是一样的
	elem = elem.Elem()                                               //elem刚开始指向一个*user地址的指针，要修改，必须要elem()指向结构体值，才能修改
	elem.FieldByName("Userid").SetString("newuseriddd")
	fmt.Println("mode model.name", model, model.Name)

	num := elem.NumField()
	fmt.Println("num==", num)

	for i := 0; i < num; i++ {
		fmt.Printf("field %d: 值为＝%v\n", i, elem.Field(i))
	}

}
