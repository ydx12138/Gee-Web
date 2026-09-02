package gee

import (
	"net/http"
)

//ServeHTTP -> router.handle -> router.getRoute -> r.hanlders[key](c)

//GET -> addRoute -> router.addRoute -> parsePattern -> r.roots[method].insert -> r.hanlders[key] = handler

// gee.go是世界中心，不断从这里产出新方法，然后提取到其他文件里，只留下一个同名方法进行引用，可称之为中央集权制
type HandlerFunc func(c *Context)
type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (e *Engine) addRoute(method string, path string, handler HandlerFunc) {
	e.router.addRoute(method, path, handler)
}

func (e *Engine) GET(addr string, handle HandlerFunc) {
	e.addRoute("GET", addr, handle)

}
func (e *Engine) POST(addr string, handle HandlerFunc) {
	e.addRoute("POST", addr, handle)

}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

// 所有的请求会经过这里然后进行路由分配。这个方法的参数不能变，因为其是上层接口的实现
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}
