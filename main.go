package main

import (
	"go-oa/api/v1/member"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/", member.GetMemberByID)
	engine.Run(":8080")
}
