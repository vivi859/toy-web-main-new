package main

import (
	"fmt"
	"geektime/toy-web/examples/first_lesson/vitest/homeproject/menu"
)

func main() {

	user1 := menu.User{
		Username: "vivi",
		Password: "xxx",
	}
	user2 := menu.NewUser()
	if user1.Username == user2.Username && user1.Password == user2.Password {
		Menu := menu.NewMenu()
		Menu.Mainmanu()
	} else {
		fmt.Println("输入的用户名和密码不对")
	}

	//key := ""
	////退出for循环的变量
	//loop := true
	//total := 10000.0
	//money := 0.0
	//node := ""
	//flag := false
	//details := "收支\t账户余额\t收支金额\t说明"
	//for {
	//	fmt.Println()
	//	fmt.Println("++++++++++++++++++++++++")
	//	fmt.Println("家庭收支明细")
	//	fmt.Println("1 收支明细")
	//	fmt.Println("2 登记收入")
	//	fmt.Println("3 登记支出")
	//	fmt.Println("4 退出")
	//	fmt.Println("++++++++++++++++++++++++")
	//	fmt.Printf("请输入你的选择：")
	//
	//	fmt.Scanln(&key)
	//	switch key {
	//	case "1":
	//		if flag == false {
	//			fmt.Println("目前还没有收支明细，请增加吧")
	//		} else {
	//			fmt.Println("-------当前收支明细记录-------")
	//			fmt.Println(details)
	//		}
	//
	//	case "2":
	//		fmt.Printf("登记收入")
	//		fmt.Printf("本次收入金额: ")
	//		fmt.Scanln(&money)
	//		total += money
	//		fmt.Printf("本地收入说明: ")
	//		fmt.Scanln(&node)
	//		details += fmt.Sprintf("\n收入\t%v\t%v\t%v", total, money, node)
	//		flag = true
	//	case "3":
	//		fmt.Println("登记支出")
	//		fmt.Printf("本次支出金额: ")
	//		fmt.Scanln(&money)
	//		if total < money {
	//			fmt.Println("取出金额大于总额")
	//		} else {
	//			total -= money
	//		}
	//		details += fmt.Sprintf("\n支出\t%v\t%v\t%v", total, money, node)
	//		flag = true
	//
	//	case "4":
	//		loop = false
	//	default:
	//		fmt.Println("请输入正确选项")
	//	}
	//	if !loop {
	//		break
	//	}
	//}
}
