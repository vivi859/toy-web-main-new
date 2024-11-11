package main

import "fmt"

func vi() {
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr[0:5]
	s2 := arr[5:10]
	s3 := arr[3:6]
	fmt.Println(arr, s1, s2, s3)
	s3[0] = 100
	fmt.Println(arr, s1, s2, s3)
	s3 = append(s3, 20, 21, 22, 23, 24)
	fmt.Println(arr, s1, s2, s3)
	s3[0] = 555
	fmt.Println(arr, s1, s2, s3)

	k1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	k2 := []int{88, 99, 77}
	copy(k1, k2)
	k1[1] = 5555

	fmt.Printf("k1=%d k2=%d\n", k1, k2)

}

func main() {
	//vi()

	a := 1
	b := 2
	arr1 := []*int{&a, &b}
	fmt.Printf("%T\n", arr1)
	fmt.Println(arr1)
	for _, val := range arr1 {
		fmt.Println(*val)
	}
	//s1 := []int{1, 2, 3, 4} // 直接初始化了 4 个元素的切片
	//fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))
	//
	//s2 := make([]int, 3, 4) // 创建了一个包含三个元素，容量为4的切片
	//fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))
	//
	//// s2 目前 [0, 0, 0], append（追加）一个元素，变成什么？
	//s2 = append(s2, 7) // 后边添加一个元素，没有超出容量限制，不会发生扩容
	//fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))
	//
	//s2 = append(s2, 8) // 后边添加了一个元素，触发扩容,超过容量后默认是自动增加一倍容量
	//fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))
	//
	//s3 := make([]int, 4) // 只传入一个参数，表示创建一个含有四个元素，容量也为四个元素的
	//// 等价于 s3 := make([]int, 4, 4)
	//fmt.Printf("s3: %v, len %d, cap: %d \n", s3, len(s3), cap(s3))
	//
	//// 按下标索引
	//fmt.Printf("s3[2]: %d", s3[2])
	//// 超出下标范围，直接崩溃
	//// runtime error: index out of range [99] with length 4
	//// fmt.Printf("s3[99]: %d", s3[99])
	//
	//SubSlice()
	//vivitest()
	//
	//arr := []int{1, 2, 3, 4}
	//fmt.Printf("\narraaaaaass: %v, len: %d, cap: %d \n", arr[0:2], len(arr), cap(arr))
	//fmt.Println(&arr[0])
	//fmt.Println(&arr[1])
	//ShareSlice()

	//shareArr()
}
func vivitest() {
	ss1 := []int{1, 2}
	fmt.Printf("vivitest : ss1: %v,len %d, cappp: %d,ss1[0]=%d,ss1[1]=%d", ss1, len(ss1), cap(ss1), ss1[0], ss1[1])

}
func SubSlice() {
	s1 := []int{2, 4, 6, 8, 10}
	s2 := s1[1:3]
	fmt.Printf("切片s1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("切片s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	s3 := s1[2:]
	fmt.Printf("s3: %v, len %d, cap: %d \n", s3, len(s3), cap(s3))

	s4 := s1[:3]
	fmt.Printf("s4: %v, len %d, cap: %d \n", s4, len(s4), cap(s4))
	s4 = append(s4, 10)
	fmt.Printf("s4: %v, len %d, cap: %d \n", s4, len(s4), cap(s4))
	fmt.Printf("s3: %v, len %d, cap: %d \n", s3, len(s3), cap(s3))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))
	s4[0] = 9999
	fmt.Printf("改了s4: %v, len %d, cap: %d \n", s4, len(s4), cap(s4))
	fmt.Printf("s3: %v, len %d, cap: %d \n", s3, len(s3), cap(s3))
	fmt.Printf("切片s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))
	fmt.Printf("切片s1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))

}

func ShareSlice() {

	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:]
	fmt.Printf("ssssss1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	s2[0] = 99
	fmt.Printf("ssssss1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	s2 = append(s2, 199)
	fmt.Printf("ssssss1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	s2[1] = 1999
	fmt.Printf("ssssss1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

}
