package main

import (
	"encoding/json"
	"fmt"
)

func testslice() string {
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
	return (string(data))
}
func unmarshalslice() ([]map[string]interface{}, error) {
	str := "[{\"Name\":\"xiaoming\"},{\"Age\":18,\"Name\":\"wang\"," +
		"\"address\":[\"北极\",\"上海\",\"北\"]}]"

	var a []map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("unmarshar_slice")
	fmt.Println(a)
	return a, nil
}
