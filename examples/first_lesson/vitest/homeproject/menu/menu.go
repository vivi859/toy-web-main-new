package menu

import "fmt"

type User struct {
	Username string
	Password string
}

func NewUser() *User {
	return &User{
		Username: "vivi",
		Password: "vivipass",
	}
}

type Menu struct {
	key string
	//退出for循环的变量
	loop    bool
	total   float64
	money   float64
	node    string
	flag    bool
	details string
}

func NewMenu() *Menu {
	return &Menu{
		key:     "",
		loop:    true,
		total:   10000.0,
		money:   0.0,
		node:    "",
		flag:    false,
		details: "收支\t账户余额\t收支金额\t说明",
	}
}

func (m *Menu) Allmoney() {
	if m.flag == false {
		fmt.Println("目前还没有收支明细，请增加吧")
	} else {
		fmt.Println("-------当前收支明细记录-------")
		fmt.Println(m.details)
	}
}

func (m *Menu) Addmoney() {
	fmt.Printf("登记收入")
	fmt.Printf("本次收入金额: ")
	fmt.Scanln(&m.money)
	m.total += m.money
	fmt.Printf("本地收入说明: ")
	fmt.Scanln(&m.node)
	m.details += fmt.Sprintf("\n收入\t%v   \t%v     \t%v", m.total, m.money, m.node)
	m.flag = true
}

func (m *Menu) Delmoney() {
	fmt.Println("登记支出")
	fmt.Printf("本次支出金额: ")
	fmt.Scanln(&m.money)
	if m.total < m.money {
		fmt.Println("取出金额大于总额")
	} else {
		m.total -= m.money
	}
	m.details += fmt.Sprintf("\n支出\t%v   \t%v      \t%v", m.total, m.money, m.node)
	m.flag = true
}

func (m *Menu) Mainmanu() {
	for {
		fmt.Println()
		fmt.Println("++++++++++++++++++++++++")
		fmt.Println("家庭收支明细")
		fmt.Println("1 收支明细")
		fmt.Println("2 登记收入")
		fmt.Println("3 登记支出")
		fmt.Println("4 退出")
		fmt.Println("++++++++++++++++++++++++")
		fmt.Printf("请输入你的选择：")
		fmt.Scanln(&m.key)
		switch m.key {
		case "1":
			m.Allmoney()
		case "2":
			m.Addmoney()
		case "3":
			m.Delmoney()
		case "4":
			m.loop = false
		default:
			fmt.Println("请输入正确选项")
		}
		if !m.loop {
			break
		}
	}
}
