package controllers

import (
	"be_crud/models"
	m "be_crud/models"
	"fmt"
	"net/http"
	"strconv"

	// "strings"

	"github.com/gin-gonic/gin"
	// "github.com/jackc/pgx/v5"
)

// @Description list all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {string} string "string"
// @Router /users [get]
func GetAllUsers(ctx *gin.Context) { 
	users := m.FindAllUser()
	

	if len(users) == 0 {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "data not found",
			Results: nil,
			Errors:  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "success get all users",
		Results: users,
	})
}

// @Description list all users
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} string "string"
// @Router /users/{id} [get]
func GetUserById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user := m.FindUserById(id)

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "success get user by id",
		Results: user,
	})
}

func CreateUser(ctx *gin.Context) {
	var req m.User
	currUsers := m.FindAllUser()

	err :=ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to bind json",
			"error": err,
		})
		return
	}

	for _, user := range currUsers {
		if user.Email == req.Email {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "email already exist",
			})
			return
		}
	}

	m.AddingNewUSer(req)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success create new user",
		"result": req,
	})
}

func UpdateUser(ctx *gin.Context) {
	var req m.User
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to convert id to int",
			"error": err,
		})
		return
	}

	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to bind json",
			"error": err,
		})
		return
	}

	fmt.Println(req, id)
	m.UpdateUser(&req, id)

	// update kalo sukses
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success update user",
		"result": req,
	})
}