package routers

import (
	docs "be_crud/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CombineRouters(r *gin.Engine) {
	usersRouter(r.Group("/users"))
	
	// Swagger 
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/docs", func(ctx *gin.Context) {		//kalo dia acess /docs, dia bakal redirect ke /docs/index.html
		ctx.Redirect(http.StatusPermanentRedirect, "/docs/index.html")
	})
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))		// kita punya path /docs yang bakal munculin swagger yang udah kita dokumentasiin tadi

}