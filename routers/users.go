package routers

import (
	c "be_crud/controllers"

	"github.com/gin-gonic/gin"
)

func usersRouter(r *gin.RouterGroup) {
	r.GET("", c.GetAllUsers)
}