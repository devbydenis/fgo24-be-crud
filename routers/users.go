package routers

import (
	c "be_crud/controllers"

	"github.com/gin-gonic/gin"
)

func usersRouter(r *gin.RouterGroup) {
	r.GET("", c.GetAllUsers)
	r.GET(":id", c.GetUserById)
	r.PATCH("", c.UpdateUser)
	r.POST("", c.CreateUser)
	r.DELETE("", c.DeleteUser)
	// r.DELETE("/:id", c.DeleteUser)

}