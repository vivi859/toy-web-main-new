package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	file, err := os.Open("/opt/vivi/gitvivi/toy-web-main/examples/first_lesson/vitest/file/file.txt")
	defer file.Close() //要及时关闭文件句柄　否则会有内存泄漏

	if err != nil {
		fmt.Println("打开文件错误了", err)
	} else {
		fmt.Println("打开了", file)
		reader := bufio.NewReader(file) //默认４０９６　读一部分缓冲一部分 如果文件比较小可以用ioutil一次性读取到内存．
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			fmt.Println(line)
		}
	}
	file1Path := "/opt/vivi/gitvivi/toy-web-main/examples/first_lesson/vitest/file/11.txt"
	file2Path := "/opt/vivi/gitvivi/toy-web-main/examples/first_lesson/vitest/file/22.txt"
	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Println(err)
		return
	}
	ioutil.WriteFile(file2Path, data, 0666)

	fmt.Println("命令行的参数有", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}
	//运行测试
	//go build -o osargs main.go
	//./osargs tom sdx/file 999
}
