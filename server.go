package web

import (
	"net"
	"net/http"
)

type HandleFunc func(ctx Context)
type Server interface {
	http.Handler
	Start(addr string) error
	AddRoute(method, path string, handleFunc HandleFunc)
}

type HttpServer struct {
}

func (h *HttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *HttpServer) Start(addr string) error {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	return http.Serve(listen, h)
}

// AddRoute 注册路由
func (h *HttpServer) AddRoute(method, path string, handleFunc HandleFunc) {
	//TODO implement me
	panic("implement me")
}
func (h *HttpServer) Get(path string, handleFunc HandleFunc) {
	h.AddRoute(http.MethodGet, path, handleFunc)
}
