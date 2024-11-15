package main

import "fmt"

func main() {
	// 创建了一个预估容量是2的 map
	m := make(map[string]string, 2)
	// 没有指定预估容量
	m1 := make(map[string]string)
	// 直接初始化
	m2 := map[string]string{
		"Tom": "Jerry",
	}

	// 赋值
	m["hello"] = "world"
	m1["hello"] = "world"
	// 赋值
	m2["hello"] = "world"
	// 取值
	val := m["hello"]
	println(val)

	// 再次取值，使用两个返回值，后面的ok会告诉你map有没有这个key
	val, ok := m["invalid_key"]
	if !ok {
		println("key not found")
	}

	for key, val := range m {
		fmt.Printf("%s => %s \n", key, val)
	}

	myMap := make(map[string]map[string]string)

	// 初始化并赋值
	myMap["user1"] = make(map[string]string)
	myMap["user1"]["name"] = "Alice"
	myMap["user1"]["age"] = "30"

	myMap["user2"] = make(map[string]string)
	myMap["user2"]["name"] = "Bob"
	myMap["user2"]["age"] = "25"

	// 打印 map
	fmt.Println(myMap)

}
