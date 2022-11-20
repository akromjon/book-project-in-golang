package route

import (
	"github.com/gin-gonic/gin"
)

type RouteIterface interface {
	Routes(*gin.Engine)
}

func RegisterRoutes(r RouteIterface, e *gin.Engine) {
	r.Routes(e)
}
