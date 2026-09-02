package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", healthy)
	r.GET("/hello", hello)

	r.POST("/login", login)

	r.Run(":8000")
}

func healthy(c *gee.Context) {
	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
}
func hello(c *gee.Context) {
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
}
func login(c *gee.Context) {
	c.JSON(http.StatusOK, gee.H{
		"username": c.PostForm("username"),
		"password": c.PostForm("password"),
	})
}
