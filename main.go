package main

import (
	"fmt"

	"github.com/HDQOSH/arms-gin-vue/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.Run()
	fmt.Print("666")
}
