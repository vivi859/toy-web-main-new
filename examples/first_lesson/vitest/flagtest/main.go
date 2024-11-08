package main

import (
	"flag"
	"fmt"
)

func main() {
	var user, pwd, host string
	var port int

	flag.StringVar(&user, "u", "", "user用户名，默认是空")
	flag.StringVar(&pwd, "p", "", "pwd密码，默认是空")
	flag.StringVar(&host, "h", "", "host，默认是空")
	flag.IntVar(&port, "P", 16303, "host，默认是16303")

	flag.Parse() //必须要执行这个进行转换
	fmt.Printf("user: %s, pwd: %s, host: %s, port: %d\n", user, pwd, host, port)

	// go build -o flat main.go
	//./flag -u root -p passwd8 -h localhost -P 16303
	//user: root, pwd: passwd8, host: localhost, port: 16303

}
