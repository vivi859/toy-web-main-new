package main

import (
	"fmt"
	"time"
)

func Witedata(intchan chan int) {
	for i := 0; i < 30; i++ {
		intchan <- i
		fmt.Println("Witedata", i)
		time.Sleep(1 * time.Second)
	}
	close(intchan)
}

func Readdata(intchan chan int, exitchan chan bool) {
	for {
		v, ok := <-intchan
		if !ok {
			break
		}
		fmt.Println("Readdata", v)
	}
	exitchan <- true
	close(exitchan)
}

func main() {
	intchan := make(chan int, 50)
	exitchan := make(chan bool, 1)
	go Witedata(intchan)
	go Readdata(intchan, exitchan)
	// 等待 Readdata 协程执行完毕
	<-exitchan
	fmt.Println("主程序退出")

}
