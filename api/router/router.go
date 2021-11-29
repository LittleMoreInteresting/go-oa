package router

import (
	"go-oa/api/v1/member"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	Router.GET("/", member.GetMemberByID)
}
