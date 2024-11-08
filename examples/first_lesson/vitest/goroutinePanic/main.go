package main

import (
	"fmt"
	"time"
)

func Sayhello() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("say hello %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func Errtest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("=====errtest panic =====")
		}
	}()
	var myCap map[int]string
	myCap[0] = "hello" //error
}

func main() {
	go Sayhello()
	go Errtest()

	for i := 0; i < 10; i++ {
		fmt.Printf("main print%d\n", i)
		time.Sleep(1 * time.Second)
	}
}
