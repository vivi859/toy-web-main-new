package server

import (
	"net/http"
	"toy-web-main/onclass/httptest/context"
)

// 这个方法接口可以随便写到哪里，不需要引入包，就可以实现接口，在项目其他地方都可以实现？？？

type Server interface {
	//这里其实就是func 但是不用加func关键字，接口一般大写开头，全局使用，实现一般用小写。 接口是一组行为的抽象。
	Route(pattern string, handlerFunc func(ctx *context.Context))
	//启动服务器
	Start(address string) error
}

type SdkHttpServer struct {
	Name string
}

// 类似构造函数
func NewHttpServer(name string) Server {
	return &SdkHttpServer{Name: name}
}

// 注册路由
//func (s *SdkHttpServer) Route(pattern string, handlerFunc http.HandlerFunc) {
//	http.HandleFunc(pattern, handlerFunc)
//}

// 注册路由2
func (s *SdkHttpServer) Route(
	pattern string,
	handlerFunc func(ctx *context.Context)) {

	http.HandleFunc(pattern, func(write http.ResponseWriter, request *http.Request) {
		ctx := NewContext(write, request)
		handlerFunc(ctx)
	})
}

func NewContext(write http.ResponseWriter, request *http.Request) *context.Context {
	return &context.Context{
		R: request,
		W: write,
	}
}
func (s *SdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

// 用户注册
type SignUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type CommonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// 注册路由2
func Signup(ctx *context.Context) {

	req := &SignUpReq{}
	err := ctx.ReadJson(req)
	if err != nil {
		//fmt.Fprintf(, "err:%v", err)
		ctx.BadrequestJson(CommonResponse{BizCode: 400, Msg: err.Error(), Data: "12323"})
		return
	}
	//body, err := io.ReadAll(ctx.R.Body)
	//if err != nil {
	//	fmt.Fprintf(ctx.W, "read body failed: %v", err)
	//	// 要返回掉，不然就会继续执行后面的代码
	//	return
	//}
	//body读取出来是byte需要反序列化
	//body, err := json.Unmarshal(ctx.R.Body, req)
	//if err != nil {
	//	fmt.Fprintf(ctx.W, "deserialized failed: %v", err)
	//	// 要返回掉，不然就会继续执行后面的代码
	//	return
	//}
	resp := &CommonResponse{
		BizCode: 200,
		Msg:     "success",
		Data:    map[string]interface{}{"user_id": req.Email, "password": req.Password, "confirmed_password": req.ConfirmedPassword},
	}
	ctx.WriteJson(http.StatusOK, resp)
	// 返回一个虚拟的 user id 表示注册成功了 respJson 这里是byte类型，要转string 2者可以互转 func Marshal(v any) ([]byte, error) {
	//fmt.Fprintf(ctx.W, string(respJson)) //[]byte切片跟string可以互相转
	//fmt.Fprintf(w,"%d",231) //注册成功
}
