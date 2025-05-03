package main

import (
	"tharun/configs"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	configs.connectdb()
	r.Run(":4000")
}
