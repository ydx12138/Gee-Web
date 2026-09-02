package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/", hello)
	engine.Run(":8000")

}
func hello(repo http.ResponseWriter, req *http.Request) {
	
	fmt.Fprintf(repo, "Hello Gee")
}
