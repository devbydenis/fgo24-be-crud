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
// @Param search query string false "Search by name"
// @Success 200 {string} string "string"
// @Router /users [get]
func GetAllUsers(ctx *gin.Context) { 
	queryName := ctx.DefaultQuery("search", "")
	
	users := m.FindAllUser(queryName)
	

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

// @Description detail all users
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} string "string"
// @Security Token
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

// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body m.CreateUser true "Request create user"
// @Success 201 {object} m.CreateUser
// @Failure 400 {object} m.Response{Success bool, Message string, Errors any}
// @Security Token
// @Router /users [post]
func CreateUser(ctx *gin.Context) {
	var req m.CreateUser
	
	err :=ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, m.Response{
			Success: false,
			Message: "failed to bind json",
			Errors:  err,
		})
		return
	}

	currUsers := m.FindAllUser(req.Name)
	if len(currUsers) > 0 {
		ctx.JSON(http.StatusInternalServerError, m.Response{
			Success: false,
			Message: "Email already exists",
		})
		return
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