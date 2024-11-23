package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST("/api/setimg", getimg)

	go r.Run(serverPort)
}
