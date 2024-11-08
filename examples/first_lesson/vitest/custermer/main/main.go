package main

//view界面

import (
	"fmt"
	"geektime/toy-web/examples/first_lesson/vitest/custermer/model"
	"geektime/toy-web/examples/first_lesson/vitest/custermer/service"
)

type view struct {
	key           int
	loop          bool
	custerservice *service.CustermerService //要调用service里面的方法所以要有这个结构体字段
}

// 显示所有客户的信息

func (this *view) list() {
	custers := this.custerservice.GetCusters()
	fmt.Println("-------------客户列表------------")
	for i := 0; i < len(custers); i++ {
		fmt.Println(custers[i].Getinfo())
	}
	fmt.Println("-------------客户列表完成------------")
}

// 添加用户
func (this *view) add() {
	fmt.Println("-------添加客户------")
	fmt.Printf("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Printf("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("手机号:")
	phone := 0
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)
	newcuster := model.NewCustermer2(name, age, phone, email)
	this.custerservice.Addcuster(*newcuster)
}

func (this *view) del() {
	fmt.Println("-------删除客户------")
	fmt.Printf("请输入要删除的客户ＩＤ:")
	id := 0
	fmt.Scanln(&id)
	fmt.Printf("请确认是否删除:y or n")
	check := ""
	fmt.Scanln(&check)
	if check == "y" {
		res := this.custerservice.DelCusters(id)
		if res {
			fmt.Sprintln("删除成功")
		} else {
			fmt.Println("id不存在，删除不成功")
		}
	}

}

func (v *view) Viewf() {
	for {
		fmt.Println("--------------客户信息管理软件-----------------")
		fmt.Println("　　　　　　　　　1　添加客户　　　　　　　　　")
		fmt.Println("　　　　　　　　　2　修改客户　　　　　　　　　")
		fmt.Println("　　　　　　　　　3　删除客户　　　　　　　　　")
		fmt.Println("　　　　　　　　　4　客户列表　　　　　　　　　")
		fmt.Println("　　　　　　　　　5　退　　出　　　　　　　　　")
		fmt.Print(" 请选中(1-5)")

		fmt.Scanln(&v.key)
		switch v.key {
		case 1:
			fmt.Println("请输入新增客户信息:")
			v.add()
		case 2:
			fmt.Println("修改客户信息:")
		case 3:
			fmt.Println("请输入要删除的客户:")
			v.del()
		case 4:
			fmt.Println("所有客户列表:")
			v.list()
		case 5:
			v.loop = true
		}
		if v.loop {
			break
		}
	}
}

func main() {

	var view view
	view.key = 1
	view.loop = false
	//需要对它进行初始化 view　custer * service.CustermerService
	view.custerservice = service.NewcusterS() //这个函数返回的是一个指针切片
	view.Viewf()
}
