package routers

import (
	"github.com/gin-gonic/gin"
	docs "be_crud/docs"
  swaggerfiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"
)

func CombineRouters(r *gin.Engine) {
	usersRouter(r.Group("/users"))
	
	// Swagger 
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))		// kita punya path /docs yang bakal munculin swagger yang udah kita dokumentasiin tadi

}