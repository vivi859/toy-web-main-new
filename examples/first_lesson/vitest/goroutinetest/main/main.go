package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	//申明一个全局互斥锁
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i < n; i++ {
		res *= i
		fmt.Println("test print +" + strconv.Itoa(i))
		//time.Sleep(1 * time.Second)
	}
	//加锁
	lock.Lock()
	myMap[n] = res
	lock.Unlock()

}

type Cat struct {
	Name string
	Age  int
}

func main() {

	//我们这里开启多个协程完成这个任务
	for i := 0; i < 10; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 10)
	lock.Lock()

	for i, v := range myMap {
		//fmt.Println(i, v)
		fmt.Printf("myMap[%d]=%d \n", i, v)
	}
	lock.Unlock()

	//for i := 1; i < 10; i++ {
	//	fmt.Println("main  +", i)
	//	time.Sleep(1 * time.Second)
	//}

	mapchan := make(chan map[string]string, 10)
	m1 := make(map[string]string)
	m1["city1"] = "天津"
	m1["city2"] = "北京"
	m2 := make(map[string]string)
	m2["hero"] = "xiao"
	mapchan <- m1
	mapchan <- m2
	res1 := <-mapchan
	res2 := <-mapchan

	fmt.Println(res1["city1"], res2["hero"])

	//任意类型的chan
	allchan := make(chan interface{}, 10)
	cat1 := Cat{Name: "tome", Age: 10}
	allchan <- cat1
	allchan <- m2
	allchan <- "hello world"
	res3 := <-allchan
	if cat, ok := res3.(Cat); ok {
		fmt.Println(cat)
	} else {
		fmt.Println("断言失败")
	}
	close(allchan)

	//遍历管道
	intchan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intchan2 <- i
	}
	close(intchan2)
	for v := range intchan2 {
		fmt.Println(v)
	}
}
