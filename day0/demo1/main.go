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
	http.HandleFunc("/hello/{go}", hellogo)
	http.HandleFunc("/hello/{golang}/", hellolang)
	var p1 people
	http.Handle("/p1", p1)
	// 第二个参数是路由分发器，如果为nil，则会使用默认的http.DefaultServeMux
	// 您可以传入自己实现的自定义 Handler 或使用第三方路由库（如 gin、echo）生成的实例，
	// 此时 http.HandleFunc 注册的路由将失效，因为请求会被您传入的 Handler 完全接管。
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func hellogo(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

func hellolang(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

type people struct {
}

func (receiver people) ServeHTTP(repo http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(repo, "p1")
	return
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
