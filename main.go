package main

import (
	"net/http"

	"gee"
)

func main() {
	r := gee.New()

	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.Json(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
