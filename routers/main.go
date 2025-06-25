package routers

import "github.com/gin-gonic/gin"

func CombineRouters(r *gin.Engine) {
	usersRouter(r.Group("/users"))
}