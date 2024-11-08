package main

import (
	"fmt"
)

// 结构体测试,指针 slice map都是nil 没有分配空间，引用要make
type Cat struct {
	Name   string //=字段=属性=filed 可以是数组和引用类型，没有赋值就对应一个零值，print(*ptr.age)不能这样写.的优先级比＊高
	Age    int
	Colors string
	//test   [3]int            //数组
	ptr   *int              //指针 默认nil打印出就是nil
	slice []int             //切片 默认nil打印是空
	map1  map[string]string //切片 默认nil打印是空
}

// 双个结构提测试 ,在字段上面写个tag用于反射机制，用于序列化和反序列化
type Pointer struct {
	x int
	y int
}

// 双个结构提测试
type Pointer2 struct {
	leftup, rightdown Pointer
}
type Pointer3 Pointer

type Monster struct {
	Name string `json:"name"` //这个就是 struct tag
	Age  int    `json:"age"`
}

// dosome栈　有自己的空间 这个是方法　变量的方法　非普通函数
func (m *Monster) dosome() {
	//m.Name = "小伙子"
	fmt.Println("猴子会抓人了，抓了＝", m)
}

type Dog struct {
	Name string `json:"name"` //这个就是 struct tag
	Age  int    `json:"age"`
}

// dosome栈　有自己的空间 这个是方法　变量的方法　非普通函数
func (m *Dog) dosome() {
	//m.Name = "狗子"
	fmt.Println("狗子会叫", m)
}

// 接口测试 如果接口里面有多个函数，那么要实现这个接口多个函数的话就必须建一个结构体包含了接口中的所有接口函数，就实现了这个接口
type AllanimalDosome interface {
	dosome()
}

type Allanimal struct {
}

// 编写一个方法，接受接口的类型变量　实现了AllanimalDosome这个接口的所有方法
func (a Allanimal) Whichanimal(alll AllanimalDosome) {
	alll.dosome()
}

// main栈　有自己的空间
func main() {

	//接口测试
	animal := Allanimal{}
	dog := &Dog{
		Name: "狗子",
		Age:  2,
	}
	monster := &Monster{
		Name: "猴子",
		Age:  5,
	}
	var a2 AllanimalDosome = dog
	a2.dosome()

	animal.Whichanimal(dog)
	animal.Whichanimal(monster)

	//bankvar := packtest.Banktrucks{
	//	Cardno:   100200200,
	//	Balance:  800,
	//	Bankname: "chinabank",
	//	Pwd:      "pass123",
	//}
	//bankvar.Bankbalance("pass12")
	//
	//bankvar.BankAddmoney(250, "pass123")
	//
	//stu := packtest.Student{
	//	Name:  "xiao王",
	//	Score: 90,
	//	Id:    11,
	//}
	//fmt.Println()
	//stu.Setage(9)
	//age1 := stu.Getage()
	//fmt.Println(age1)
	////嵌入结构体测试SmallStudent　嵌入student结构体
	//Pupil := &packtest.SmallStudent{}
	//Pupil.Student.Name = "jerry"
	//Pupil.Student.Score = 100
	//Pupil.Student.Setage(99) //Pupil.Getage()也可以简化
	//
	//fmt.Println("Student.Getage()", Pupil.Student.Getage())

	////学生的结构体
	//stu := packtest.Student{
	//	Name:  "xiao王",
	//	Score: 90,
	//	Age:   10,
	//	Id:    11,
	//}
	//stu.Dosome()
	////返回结构体的指针类型
	//var stu2 = &packtest.Student{
	//	Name:  "xiao",
	//	Score: 90,
	//	Age:   10,
	//}
	//fmt.Println(stu2, *stu2) //取的就是地址 取值就是＊
	//fmt.Println()
	//var stu3 = packtest.New_stutt_factory("小李里", 10)
	//fmt.Println(stu3.Age, stu3.Dosome2())

	////双个结构提测试　ptrs的四个int变量字段地址是连续分布，读取速度会更快
	////ptrs := Pointer2{Pointer{1, 2}, Pointer{3, 4}}
	////fmt.Println(ptrs)
	//monsters := Monster{
	//	Name: "小猴子",
	//	Age:  20,
	//}
	//monsters.dosome()
	//fmt.Println(monsters.Name)
	//
	////将monsters序列化成json，转换出是byte
	//jsonvar, err := json.Marshal(monsters)
	//if err != nil {
	//	fmt.Println("json处理错误", err)
	//}
	//fmt.Println(string(jsonvar)) //通过tag后打印出来的Name 就是小写name
	//
	////结构体测试
	//var cat1 Cat
	//cat1.Name = "小白"
	//cat1.Age = 1
	//cat1.Colors = "白色哦"
	//cat1.slice = make([]int, 2)
	//cat1.slice[0] = 22
	//cat1.map1 = make(map[string]string) //这个大小不需要设置会自动增长
	//cat1.map1["key1"] = "value12"
	//
	//cat2 := cat1 //值赋值　就是在内存中是另外一个内存空间，复制了一份cat内存和多一个变量cat2
	//cat2.Colors = "黄色哦"
	//
	//cat3 := Cat{
	//	Name:   "小花花",
	//	Age:    1,
	//	Colors: "黄黄的",
	//}
	//
	//var cat4 *Cat = new(Cat)
	//var cat5 *Cat = &Cat{} //也可以这样定义
	//
	//cat4.Name = "指针猫咪" //*cat4.name=xx　标准应该这样写，go底层为了简化做了优化
	//cat4.Age = 2
	//cat4.Colors = "指针猫咪颜色花"
	//
	//fmt.Printf("结构体测试＝＝＝＝＝＝", cat5)
	////fmt.Println(&cat5)
	//fmt.Println(cat4)
	//fmt.Println(*cat4)
	//fmt.Println(cat2)
	//fmt.Println(cat3)
	//fmt.Println(cat1)

}
