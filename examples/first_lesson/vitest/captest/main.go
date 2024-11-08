package main

import "fmt"

type student struct {
	name string
	age  int
}

type allstudent struct {
	students   []student
	studentsum int
}

func (this *student) printstu() {
	fmt.Printf("this is a student %v \n", this.name)
}

func (this *allstudent) printstu() {
	fmt.Printf("this is all student %v \n", this.allstudent)
}

func main() {

}
