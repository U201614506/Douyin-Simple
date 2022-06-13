package main

import (
	"allright-tiktok/controller"

	"github.com/gin-gonic/gin"
)

func main() {


	r := gin.Default()

	controller.Connection()

	initRouter(r)


	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
