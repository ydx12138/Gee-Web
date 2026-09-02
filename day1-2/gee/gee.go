package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(repo http.ResponseWriter, req *http.Request)
type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	engine := &Engine{
		router: make(map[string]HandlerFunc),
	}
	return engine
}

func (e *Engine) ServeHTTP(repo http.ResponseWriter, req *http.Request) {
	//解析请求的路径，查找路由映射表，如果查到，就执行注册的处理方法。如果查不到，就返回 404 NOT FOUND 。
	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(repo, req)
	} else {
		fmt.Fprintf(repo, "404 NOT FOUND：%s\n", req.URL)
	}
}
func (e *Engine) addRoute(method string, path string, handler HandlerFunc) {
	key := method + "-" + path
	e.router[key] = handler
}

func (e *Engine) GET(addr string, handle HandlerFunc) {
	e.addRoute("GET", addr, handle)

}
func (e *Engine) POST(addr string, handle HandlerFunc) {
	e.addRoute("POST", addr, handle)

}

func (e *Engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}
