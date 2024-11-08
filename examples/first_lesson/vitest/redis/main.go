package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

// 全局pool 当启动程序就初始化连接池
var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8, //最大空闲连接数
		MaxActive:   0, //最大连接数　０不限制
		IdleTimeout: 100 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
	//c := pool.Get()
	//pool.Close()
}

func main() {
	//redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("redis连接成功", conn)
	//开始写数据到redis
	_, err = conn.Do("Set", "goname", "hello")
	if err != nil {
		fmt.Println(err)
	}
	_, err = conn.Do("expire", "goname", 10)
	if err != nil {
		fmt.Println(err)
	}
	//开始读 返回的值是interface{}任意类型 hgetall user1
	res, err := redis.String(conn.Do("get", "goname"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	res1, err := redis.Strings(conn.Do("hgetall", "user2"))
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range res1 {
		fmt.Println(i, v)
	}
	fmt.Println(res1)

	//连接池连接
	getconn := pool.Get()
	defer getconn.Close()
	_, err = conn.Do("set", "jerry", "hohoho")
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.String(conn.Do("get", "jerry"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}
