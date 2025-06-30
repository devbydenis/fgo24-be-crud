package main

import (
	"be_crud/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)
// @title           fgo24-be-crud
// @version         1.0
// @description     This is a sample server celler server.
// @BasePath         /
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
