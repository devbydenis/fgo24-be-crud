package main

import (
	"be_crud/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)
// @title           fgo24-be-crud
// @version         1.0
// @description     This is a collection of CRUD API.
// @BasePath         /

// @SecurityDefinitions.ApiKey  Token
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "FGO24-BE-CRUD",
		})
	})

	routers.CombineRouters(r)

	r.Run(":8800")
}
