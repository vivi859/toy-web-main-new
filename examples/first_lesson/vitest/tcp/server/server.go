package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close() //关闭很重要
	//循环读取
	for {
		buf := make([]byte, 1024)
		//这里协程会堵塞　会一直等客户端发数据过来，跟客户端连接已经形成，套接字形成
		n, err := conn.Read(buf) //Read(b []byte) (n int, err error)
		//fmt.Printf("循环等待客户端的%v输入数据\n", conn.RemoteAddr().String())
		if err != nil {
			fmt.Println("read error:", err)
			return
		}
		//显示客户端发送的内容到服务器的终端　不要打println
		fmt.Println()
		fmt.Print(string(buf[:n])) //这里的n很重要,如果不带n，后面很多内容不知道是啥，
	}
}

func main() {

	fmt.Println("服务器开始监听")
	listener, err := net.Listen("tcp", ":8889")
	if err != nil {
		fmt.Println("listen err=", err)
	}
	defer listener.Close()

	//循环连接等到客户端来连接
	for {
		fmt.Println("等待客户端连接")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err=", err) //其中一个连接失败了　但是不需要退出
		} else {
			fmt.Printf("连接客户端成功 con=%v,客户端的ip＝%v\n", conn, conn.RemoteAddr().String())
		}
		//这里要准备起来一个协程为客户端服务 接受客户端数据不要再这里写，会堵塞主线程工作
		go process(conn)
	}

	fmt.Println("listen succes=", listener)
}
