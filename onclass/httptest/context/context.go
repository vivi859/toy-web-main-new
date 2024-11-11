package context

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Context struct {
	//这个是接口 ，不需要用*指针
	W http.ResponseWriter
	// Request是一个结构体所以要用指针
	R *http.Request
}

func (c *Context) ReadJson(v interface{}) error {
	//读body处理
	body, err := io.ReadAll(c.R.Body)
	if err != nil {
		fmt.Fprintf(c.W, "read body failed: %v", err)
		// 记住要返回，不然就还会执行后面的代码
		return err
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		//fmt.Fprintf(w, "read the data one more time got error: %v", err)
		c.W.WriteHeader(http.StatusBadRequest) // 使用标准库的 http.StatusBadRequest 常量
		fmt.Fprintf(c.W, "Invalid JSON format: %v", err)
		// 打印错误详细信息到控制台
		fmt.Printf("Error in JSON unmarshalling: %v\n", err)
		fmt.Printf("Request body content: %s\n", string(body))
		return err
	}
	//这里不需要return req ?
	return nil
}

// 接受任意类型的话就用interface{} 空接口
func (c *Context) WriteJson(code int, resp interface{}) error {
	c.W.WriteHeader(code)
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	//err 前面申明过了就不需要再用:=了，
	_, err = c.W.Write(respJson)
	return err
}

func (c *Context) StatusOkJson(resp interface{}) error {
	return c.WriteJson(http.StatusOK, resp)
}
func (c *Context) BadrequestJson(resp interface{}) error {
	return c.WriteJson(http.StatusBadRequest, resp)
}
