package main

import (
	"fmt"
)

func main() {
	carmelaa := Camera{}
	pann := Pan{}
	usbb := UsbWorker{}
	usbb.UsbWorker(&carmelaa)
	usbb.UsbWorker(&pann)
	//定义多态数组 结构体变量
	var allusb []Usb
	allusb = append(allusb, &carmelaa, &pann)
	var UsbWorker = UsbWorker{}
	// 遍历多态数组并执行 start 和 stop 方法
	for _, v := range allusb {
		fmt.Printf("%v", v)
		UsbWorker.UsbWorker(v)
	}

}

type Usb interface {
	start()
	stop()
}
type UsbWorker struct{}

// 多态
func (u *UsbWorker) UsbWorker(usb Usb) {
	usb.start()
	//断言进行判断
	if pan, ok := usb.(*Pan); ok {
		pan.copy()
	}
	usb.stop()
}

type Camera struct {
}

func (c *Camera) start() {
	fmt.Println("Camera start")
}
func (c *Camera) stop() {
	fmt.Println("Camera stop")
}

type Pan struct {
}

func (p *Pan) start() {
	fmt.Println("Pan start")
}
func (p *Pan) stop() {
	fmt.Println("Pan stop")
}
func (p *Pan) copy() {
	fmt.Println("Pan 复制东西")
}
