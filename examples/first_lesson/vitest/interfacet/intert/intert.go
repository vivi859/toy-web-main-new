package intert

import "fmt"

type Monkey struct {
	Name string
}

func (m *Monkey) Tree() {
	fmt.Printf("猴子会上树 %s\n", m.Name)
	fmt.Println()
}

type Fish interface {
	Swiming()
}

func Fishfunc(f Fish) {
	f.Swiming()
}

type LittleMonkey struct {
	Monkey
}

func (m *LittleMonkey) Swiming() {
	fmt.Printf("动物游泳 %s\n", m.Name)
	fmt.Println()
}

type Duck struct {
	Duck string
}

func (m *Duck) Swiming() {
	fmt.Printf("鸭子游泳 %s\n", m.Duck)
	fmt.Println()
}
