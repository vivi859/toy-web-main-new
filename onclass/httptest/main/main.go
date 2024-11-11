package main

import (
	"fmt"
	"net/http"
	"toy-web-main/onclass/httptest/server"
)

func main() {
	servers := server.NewHttpServer("test-server")
	//server.Route("/", home)
	//server.Route("/user", user)
	//server.Route("/user/create", createUser)
	//server.Route("/order", order)
	servers.Route("/user/signup", server.Signup)
	servers.Start(":8080")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是主页")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是用户")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是创建用户")
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是订单")
}

//context package//context package//context package//context package//context package//context package
//context package//context package//context package//context package//context package//context package
//context package//context package//context package//context package//context package//context package

// server package //server package//server package//server package//server package
// server package //server package//server package//server package//server package
// server package //server package//server package//server package//server package
