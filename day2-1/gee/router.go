package gee

import (
	"log"
	"net/http"
)

type router struct {
	hanlders map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		make(map[string]HandlerFunc),
	}
}

// 添加路由
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.hanlders[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handle, ok := r.hanlders[key]; ok {
		handle(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
