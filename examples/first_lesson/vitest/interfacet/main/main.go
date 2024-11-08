package main

import (
	"fmt"
	"geektime/toy-web/examples/first_lesson/vitest/interfacet/intert"
	"math/rand"
	"sort"
)

type Student struct {
	name  string
	age   int
	score int
}
type Student2 struct {
	Student
}

// 切片类型
type studentSlice []Student

// 要实现这个接口，必须实现它的３个方法，func Sort(data Interface) 这个接口调用，
func (s *studentSlice) Len() int {
	return len(*s)
}
func (s *studentSlice) Less(i, j int) bool {
	return (*s)[i].score < (*s)[j].score
}

func (s *studentSlice) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func main() {

	//切片数组　引用类型　数组排序
	var intss = []int{0, 1, 2, 3, 15, 8}
	sort.Ints(intss) //ints排序
	fmt.Println(intss)

	var studentslice studentSlice
	for i := 0; i < 10; i++ {
		students := Student{
			name:  fmt.Sprintf("student%d", i),
			age:   i,
			score: rand.Intn(100),
		}
		studentslice = append(studentslice, students)
	}

	for _, v := range studentslice {
		fmt.Println(v)
	}
	//sort.sort直接调用
	sort.Sort(&studentslice)
	fmt.Println(studentslice)

	student2 := Student2{Student{name: "小学生", age: 1, score: 90}}
	studentslice = append(studentslice, student2.Student) // 使用 student2.Student 添加到切片
	fmt.Println(studentslice)

	mm := intert.LittleMonkey{intert.Monkey{
		Name: "孙悟空",
	}}

	mm.Tree()
	mm.Swiming()
	dd := intert.Duck{
		Duck: "丫丫子",
	}
	dd.Swiming()
}
