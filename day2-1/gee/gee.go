package gee

import (
	"net/http"
)

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
	//解析请求的路径，查找路由映射表，如果查到，就执行注册的处理方法。如果查不到，就返回 404 NOT FOUND 。
	//key := req.Method + "-" + req.URL.Path
	//if handler, ok := e.router.hanlders[key]; ok {
	//	handler(&Context{
	//		Writer:     w,
	//		Req:        req,
	//		Path:       req.URL.Path,
	//		Method:     req.Method,
	//		StatusCode: 0,
	//	})
	//} else {
	//	fmt.Fprintf(w, "404 NOT FOUND：%s\n", req.URL)
	//}
	c := newContext(w, req)
	e.router.handle(c)
}
