package main

import (
	"fmt"
	"log"
	"net/http"
)

// http://localhost:8000/login/:ydx会进入哪个处理函数？？ "/login/{rest...}"还是"/login/{name}"？？
// 答案是"/login/{name}"，因为其更具体，当有两个路径都匹配时，会选择更具体的那个
func main() {

	http.HandleFunc("/", gang)                   // /是兜底路由，没有被接住的请求都会落到这里
	http.HandleFunc("/login/1/2/3/4", login)     // 匹配/login/1/2/3/4
	http.HandleFunc("/login/{rest...}", handler) // 会匹配所有/login/*
	http.HandleFunc("/login/{name}", name)       // 匹配/login/:name
	http.HandleFunc("/hello", hello)
	p1 := new(Engine)
	log.Fatal(http.ListenAndServe("localhost:8000", p1))

}

type Engine struct {
}

func (receiver *Engine) ServeHTTP(repo http.ResponseWriter, req *http.Request) {
	fmt.Println("ServeHTTP")
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(repo, "/")
	case "/login":
		fmt.Fprintf(repo, "/")
	case "/hello":
		fmt.Fprintf(repo, "/")
	}
}
func hello(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

func gang(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "你的请求路径不存在")
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "登录成功")
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "name = %v\n", r.PathValue("rest"))
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Fprintf(w, "这里是handler")
}
func name(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "name = %v\n", r.PathValue("name"))
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Fprintf(w, "这里是name")
}
