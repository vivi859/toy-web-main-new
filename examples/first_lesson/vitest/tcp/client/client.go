package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	//连接服务端的地址
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("连接服务端成功", conn.RemoteAddr())
	//循环输入，直到输入exit退出
	for {
		//发送数据到服务端,客户端终端输入os.stdin标准输入
		reader := bufio.NewReader(os.Stdin)
		//从终端读取一行用户输入，发送给服务器　\n是中止符一行的
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readstring error ", err)
		}
		// 去除两端的空白字符，包括换行符
		line = strings.TrimSpace(line) //这行很重要

		if line == "exit" {
			return
		}
		n, err := conn.Write([]byte(line)) //强转　line转换成byte切片
		if err != nil {
			fmt.Println("写数据发送给服务端出错了", err)
		}
		fmt.Printf("客户端发送数据成功了,字节数＝%v\n", n)
	}
}
