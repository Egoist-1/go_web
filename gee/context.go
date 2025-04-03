package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}
type Context struct {
	Req    *http.Request
	Res    http.ResponseWriter
	Method string
	Path   string
	Status int
}

func newContext(req *http.Request, res http.ResponseWriter) *Context {
	return &Context{Req: req, Res: res}
}
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}
func (c *Context) PostFrom(key string) string {
	return c.Req.PostFormValue(key)
}
func (c *Context) SetStatus(code int) {
	c.Status = code
	c.Res.WriteHeader(code)
}
func (c *Context) SetHeader(key, val string) {
	c.Res.Header().Set(key, val)
}
func (c *Context) String(code int, key string, val ...string) {
	c.SetStatus(code)
	c.SetHeader("Content-Type", "text/plain")
	c.Res.Write([]byte(fmt.Sprintf(key, val)))
}
func (c *Context) JSON(code int, obj any) {
	c.SetStatus(code)
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Res)
	err := encoder.Encode(obj)
	if err != nil {
		http.Error(c.Res, err.Error(), 500)
	}
}
