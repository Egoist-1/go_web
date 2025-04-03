package gee

import "net/http"

type HandlerFunc func(*Context)
type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, pattern string, h HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = h
}
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Req.URL.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	}
	c.String(http.StatusNotFound, "404 page not found: %s\n", c.Path)
}
