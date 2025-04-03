package gee

import (
	"net/http"
)

type Engine struct {
	r *router
}

func New() *Engine {
	return &Engine{r: newRouter()}
}

func (e *Engine) addRouter(method, pattern string, h HandlerFunc) {
	e.r.addRoute(method, pattern, h)
}
func (e *Engine) Post(pattern string, h HandlerFunc) *Engine {
	e.addRouter("POST", pattern, h)
	return e
}
func (e *Engine) Get(pattern string, h HandlerFunc) *Engine {
	e.addRouter("GET", pattern, h)
	return e
}

// 所有请求有处理
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(req, w)
	e.r.handle(c)
}
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
