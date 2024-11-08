package main

import "fmt"

var Userid int
var Userpass string

func main() {

	//显示菜单
	var key int
	var loop = true

	for {
		fmt.Println("------------------------欢迎登录多人聊天系统------------------------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t ２ 注册用户")
		fmt.Println("\t\t\t 1 退出系统")
		fmt.Println("\t\t\t 请选择（１－３）：")
		fmt.Scanln("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("输入有误，重新输入")
		}

	}

	if key == 1 {
		fmt.Println("请输入你的ＩＤ：")
		fmt.Scanln("%d\n", &Userid)
		fmt.Scanln("%d\n", &Userpass)

	} else {

	}

}
