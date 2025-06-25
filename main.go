package main

import (
	"be_crud/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "FGO24-BE-CRUD",
		})
	})

	routers.CombineRouters(r)

	r.Run(":8888")
}
