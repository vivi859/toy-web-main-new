package service

//业务逻辑
//　要有一个list方法获取所有切片信息
import (
	"geektime/toy-web/examples/first_lesson/vitest/custermer/model"
)

type CustermerService struct {
	Custers   []model.CusterMer
	Cutersnum int
}

// 返回指针类型切片
func NewcusterS() *CustermerService {
	custers := &CustermerService{}
	custers.Cutersnum = 1
	//初始化一个用户
	custer := model.NewCustermer(1, "顾客１", 22, 1389483492, "vivi@gmail.com") //用户信息库里的工厂模式方法
	custers.Custers = append(custers.Custers, *custer)
	return custers
}

// 返回客户切片
func (this *CustermerService) GetCusters() []model.CusterMer {
	return this.Custers
}

func (this *CustermerService) Addcuster(custer model.CusterMer) bool {
	this.Cutersnum++
	custer.Id = this.Cutersnum
	this.Custers = append(this.Custers, custer)
	return true
}

func (this *CustermerService) Getid(id int) int {
	//idd := -1
	//fmt.Println(len(this.Custers))
	for i := 0; i < len(this.Custers); i++ {
		if this.Custers[i].Id == id {
			return i
		}
	}
	return -1
}

func (this *CustermerService) DelCusters(id int) bool {
	id = this.Getid(id)
	if id != -1 {
		//fmt.Printf("\nidddd=%v\n,%v", id, this.Custers[id])
		this.Custers = append(this.Custers[:id], this.Custers[id+1:]...)
		return true
	} else {
		return false
	}
}
