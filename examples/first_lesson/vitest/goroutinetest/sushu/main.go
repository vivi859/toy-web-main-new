package main

import (
	"fmt"
	"time"
)

func Putdata(intchan chan int) {
	for i := 0; i < 5000; i++ {
		intchan <- i
	}
	close(intchan)
}

func Getdata(intchan chan int, reschan chan int, exitchan chan bool) {

	var flag bool
	for {
		num, ok := <-intchan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			reschan <- num
		}
	}
	fmt.Println("有一个协程取不到数据退出了") //这只是其中一个协程退出了，不能关闭reschan管道
	exitchan <- true
}

func main() {

	intchan := make(chan int, 1000)
	reschan := make(chan int, 2000)
	exitchan := make(chan bool, 4)

	time1 := time.Now().Unix()
	go Putdata(intchan)

	//要获取到４个退出协程信号
	for i := 0; i < 4; i++ { //等于是开启了４个协程
		go Getdata(intchan, reschan, exitchan)
	}
	//起的匿名函数　类似协程
	go func() {
		for i := 0; i < 4; i++ {
			<-exitchan
		}

		time2 := time.Now().Unix()
		fmt.Println("协程时间：", time2-time1)
		//当我们从exitchan管到取出４个结果　可以放心关闭reschan
		close(reschan)
	}()

	for {
		res, ok := <-reschan
		if !ok {
			break
		}
		fmt.Printf("素数是%d\n", res)
	}

}
