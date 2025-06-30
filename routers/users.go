package routers

import (
	c "be_crud/controllers"

	"github.com/gin-gonic/gin"
)

func usersRouter(r *gin.RouterGroup) {
	r.GET("", c.GetAllUsers)
	r.GET(":id", c.GetUserById)
	r.POST("", c.CreateUser)
	r.PUT("/:id", c.UpdateUser)
	// r.DELETE("/:id", c.DeleteUser)

}