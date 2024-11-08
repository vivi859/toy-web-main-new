package main

import "fmt"

func Typefun(items ...interface{}) {
	for index, item := range items {
		switch item.(type) {
		case int:
			fmt.Printf("第%v个是　int type: %T\n", index, item)
		case string:
			fmt.Printf("第%v个是　string type: %T\n", index, item)
		case bool:
			fmt.Printf("第%v个是　bool type: %T\n", index, item)
		case student:
			fmt.Printf("第%v个是　student type: %T\n", index, item)
		case *student:
			fmt.Printf("第%v个是　*student type: %T\n", index, item)
		default:
			fmt.Printf("第%v个是　unknown type: %T\n", index, item)
		}
	}
}

type student struct{}

func main() {
	student1 := student{}
	student2 := &student{}

	Typefun(1, 2, 3, 4, true, "tome", student1, student2)

}
