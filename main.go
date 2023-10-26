package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func handleFunction(ctx *gin.Context) {
	ctx.JSON(200, "This is a simple webpage")
}

func main() {
	fmt.Print("What's up")
	r := gin.Default()
	r.GET("/", handleFunction)
	err := r.Run()
	if err != nil {
		return
	}
}
