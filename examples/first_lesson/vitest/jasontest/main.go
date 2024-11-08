package main

import (
	"encoding/json"
	"fmt"
)

func unmarshalStuct() {
	//假设下马的str是读到的数据，本文中需要\转义,一般都是程序读取
	str := "{\"name_people\":\"wang\",\"age_people\":18,\"phone_people\":138293}"
	var people People
	err := json.Unmarshal([]byte(str), &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("unmarshar")
	fmt.Println(people)

}

func unmarshalslice() {
	str := "[{\"Name\":\"xiaoming\"},{\"Age\":18,\"Name\":\"wang\"," +
		"\"address\":[\"北极\",\"上海\",\"北\"]}]"

	var a []map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("unmarshar_slice")
	fmt.Println(a)

}

type People struct {
	Name  string `json:"name_people"` //tag　反射机制
	Age   int    `json:"age_people"`
	Phone int    `json:"phone_people"`
}

func Peoplefun() {
	people := People{
		Name:  "小李",
		Age:   20,
		Phone: 123828777,
	}
	fmt.Println(people)
	data, err := json.Marshal(people)
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println("序列化后")
	fmt.Println(string(data))

}
func testmap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["Name"] = "xiaoming"
	a["Age"] = 18
	a["address"] = "six snidfn road"
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

}

func testslice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["Name"] = "xiaoming"
	slice = append(slice, m1)
	m2 := make(map[string]interface{})
	m2["Name"] = "wang"
	m2["Age"] = 18
	m2["address"] = []string{"北极", "上海", "北"}
	slice = append(slice, m2)
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func main() {
	Peoplefun()
	testmap()
	testslice()
	unmarshalStuct()
	unmarshalslice()
}
