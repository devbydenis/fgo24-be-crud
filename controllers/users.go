package controllers

import (
	"be_crud/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) { 
	users := models.FindAllUser()
	fmt.Println("",users)

	if len(users) == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success get all data",
		"data": users,
	})
}