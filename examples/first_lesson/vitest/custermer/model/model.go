package model

//客户信息custermer
import "fmt"

type CusterMer struct {
	Id    int
	Name  string
	Age   int
	Phone int
	Email string
}

// 工厂模式
func NewCustermer(id int, name string, age int,
	phone int, email string) *CusterMer {
	return &CusterMer{
		Id:    id,
		Name:  name,
		Age:   age,
		Phone: phone,
		Email: email,
	}
}

// 工厂模式 没有ＩＤ的
func NewCustermer2(name string, age int,
	phone int, email string) *CusterMer {
	return &CusterMer{
		Name:  name,
		Age:   age,
		Phone: phone,
		Email: email,
	}
}

// 返回用户信息
func (this CusterMer) Getinfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\n", this.Id, this.Name, this.Age, this.Phone, this.Email)
	return info
}
