// Code generated by github.com/SweetWhen/goarch/custompb. DO NOT EDIT.

package helloworld

import (
	context "context"
	errors "errors"
	gin "github.com/gin-gonic/gin"
	metadata "google.golang.org/grpc/metadata"
)

// This is a compile-time assertion to ensure that this generated file
// context.metadata.
//gin.errors.

type GreeterHTTPServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)

	SayHello2(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterHTTPServer(r gin.IRouter, srv GreeterHTTPServer) {
	s := Greeter{
		server: srv,
		router: r,
		resp:   defaultGreeterResp{},
	}
	s.RegisterService()
}

type Greeter struct {
	server GreeterHTTPServer
	router gin.IRouter
	resp   interface {
		Error(ctx *gin.Context, err error)
		ParamsError(ctx *gin.Context, err error)
		Success(ctx *gin.Context, data interface{})
	}
}

// Resp 返回值
type defaultGreeterResp struct{}

func (resp defaultGreeterResp) response(ctx *gin.Context, status, code int, msg string, data interface{}) {
	ctx.JSON(status, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// Error 返回错误信息
func (resp defaultGreeterResp) Error(ctx *gin.Context, err error) {
	code := -1
	status := 500
	msg := "未知错误"

	if err == nil {
		msg += ", err is nil"
		resp.response(ctx, status, code, msg, nil)
		return
	}

	type iCode interface {
		HTTPCode() int
		Message() string
		Code() int
	}

	var c iCode
	if errors.As(err, &c) {
		status = c.HTTPCode()
		code = c.Code()
		msg = c.Message()
	}

	_ = ctx.Error(err)

	resp.response(ctx, status, code, msg, nil)
}

// ParamsError 参数错误
func (resp defaultGreeterResp) ParamsError(ctx *gin.Context, err error) {
	_ = ctx.Error(err)
	resp.response(ctx, 400, 400, "参数错误", nil)
}

// Success 返回成功信息
func (resp defaultGreeterResp) Success(ctx *gin.Context, data interface{}) {
	resp.response(ctx, 200, 0, "成功", data)
}

func (s *Greeter) SayHello_0(ctx *gin.Context) {
	var in HelloRequest

	if err := ctx.ShouldBindQuery(&in); err != nil {
		s.resp.ParamsError(ctx, err)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(GreeterHTTPServer).SayHello(newCtx, &in)
	if err != nil {
		s.resp.Error(ctx, err)
		return
	}

	s.resp.Success(ctx, out)
}

func (s *Greeter) SayHello2_0(ctx *gin.Context) {
	var in HelloRequest

	if err := ctx.ShouldBindQuery(&in); err != nil {
		s.resp.ParamsError(ctx, err)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(GreeterHTTPServer).SayHello2(newCtx, &in)
	if err != nil {
		s.resp.Error(ctx, err)
		return
	}

	s.resp.Success(ctx, out)
}

// the f**k http register func
func (s *Greeter) RegisterService() {

	s.router.Handle("GET", "/api/v1/hello", s.SayHello_0)

	s.router.Handle("GET", "/api/v1/hello2", s.SayHello2_0)

}