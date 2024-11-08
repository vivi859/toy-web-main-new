package packtest

import "fmt"

// 学生的结构体
type Student struct {
	Name  string
	age   int
	Id    int
	Score int
}

// 这个是嵌入了student结构体，匿名结构体，可以调用Student的所有方法,继承方法
type SmallStudent struct {
	Student        //abc Student这样是有名结构体　引用要用abc.xx带属性名字
	Name    string //　这里也有name，继承student也有name，但是采取就近赋值，会先用这个，包括方法相同也是一样的原理
}

func (s *Student) Setage(age int) {
	s.age = age + 5
}
func (s *Student) Getage() int {
	return (s.age)
}

func (s *Student) Dosome() {
	fmt.Printf("student的信息 %v,%v,%v,%v", s.Name, s.age, s.Id, s.Score)
}

// stutt_factory小写的结构体　如果在其他包要用可以用工厂模式
type stutt_factory struct {
	Name string
	Age  int
}

func (s *stutt_factory) Dosome2() string {
	fmt.Printf("student的信息 %v,%v", s.Name, s.Age)
	fmt.Println()
	return s.Name
}

// 新建一个方法把结构体的类型传入，返回是一个结构体指针，这个就是工厂模式，就是增加了一个方法
func New_stutt_factory(a string, b int) *stutt_factory {
	return &stutt_factory{
		Name: a,
		Age:  b,
	}
}

type Banktrucks struct {
	Cardno   int
	Bankname string
	Balance  float64
	Pwd      string
}

func (b *Banktrucks) Bankbalance(pwd string) {
	if b.Pwd != pwd {
		fmt.Println("密码不对")
	} else {
		fmt.Printf("余额：%v", b.Balance)
	}
}
func (b *Banktrucks) BankAddmoney(money float64, pwd string) {
	if b.Pwd != pwd {
		fmt.Println("密码不对")
	} else {
		b.Balance += money
		fmt.Printf("余额：%v", b.Balance)
	}
}
