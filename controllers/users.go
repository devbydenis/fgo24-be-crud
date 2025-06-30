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

// @summary List all users
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

// @summary Detail all users
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

// @summary Create a new user
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

// @summary Update user details
// @Description Update user details by user ID
// @Tags users
// @Accept  json
// @Produce json
// @Param id path int true "User ID"
// @Param user body m.UpdateUserType true "Update user"
// @Success 200 {object} m.UpdateUserType
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Security Token
// @Router /users [patch]
func UpdateUser(ctx *gin.Context) {
	var req m.UpdateUserType

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, m.Response{
			Success: false,
			Message: "failed to bind json",
			Errors:  err,
		})
		return
	}

	fmt.Println("id: ", req.ID)
	fmt.Println("email: ", req.Email)
	fmt.Println("name: ", req.Name)
	m.UpdateUser(&req)

	// update kalo sukses
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success update user",
		"result": req,
	})
}

// @Summary Delete a user
// @Description Delete user by user ID
// @Tags users
// @Accept  json
// @Produce json
// @Param user body m.DeleteUserType true "Delete User"
// @Success 204
// @Failure 404 {object} map[string]interface{}
// @Router /users [delete]
func DeleteUser(ctx *gin.Context) {
	var req m.DeleteUserType

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, m.Response{
			Success: false,
			Message: "failed to bind json",
			Errors:  err,
		})
		return
	}

	m.DeleteUser(req.ID)
	ctx.JSON(http.StatusOK, m.Response{
		Success: true,
		Message: "success delete user",
	})

}