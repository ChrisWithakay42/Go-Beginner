package controller

import (
	"github.com/gin-gonic/gin"
	"go_rest/data/request"
	"go_rest/data/response"
	"go_rest/service"
	"go_rest/utils"
	"net/http"
	"strconv"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (controller *UserController) Create(ctx *gin.Context) {
	createUserRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	utils.ErrorPanic(err)

	controller.userService.Create(createUserRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) Update(ctx *gin.Context) {
	updateUserRequest := request.UpdateUserRequest{}
	err := ctx.ShouldBindJSON(&updateUserRequest)
	utils.ErrorPanic(err)

	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	utils.ErrorPanic(err)
	updateUserRequest.Id = id

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) Delete(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	utils.ErrorPanic(err)
	controller.userService.Delete(id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) FindAll(ctx *gin.Context) {
	userResponse := controller.userService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) FindById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	utils.ErrorPanic(err)

	userResponse := controller.userService.FindById(id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
