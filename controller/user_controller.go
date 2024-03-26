package controller

import (
	"jin-example/data/request"
	"jin-example/data/response"
	"jin-example/helper"
	service "jin-example/service/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UserControllers struct {
	userService service.UserService
}

func NewUserController(service service.UserService) *UserControllers {
	return &UserControllers{
		userService: service,
	}
}

// CreateUser 接口文档
// @Summary 创建用户
// @Description 在数据库中保存用户数据
// @Tags user
// @Accept  json
// @Produce  application/json
// @Param user body request.CreateUserRequest true "Create User"
// @Success 200 {object} response.Response
// @Router /user [post]
func (controller *UserControllers) CreateUser(ctx *gin.Context) {
	log.Info().Msg("create user")

	createUserRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helper.ErrorPanic(err)

	controller.userService.CreateUser(createUserRequest)

	createUserResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, createUserResponse)
}

// UpdateUser 接口文档
// @Summary 更新用户
// @Description 更新用户数据
// @Tags user
// @Accept  json
// @Produce  application/json
// @Param tagId path string true "update user by id"
// @Param user body request.CreateUserRequest true "Update User"
// @Success 200 {object} response.Response
// @Router /user/{tagId} [patch]
func (controller *UserControllers) UpdateUser(ctx *gin.Context) {
	log.Info().Msg("update user")

	updateUserRequest := request.UpdateUserRequest{}
	err := ctx.ShouldBindJSON(&updateUserRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateUserRequest.Id = id

	controller.userService.UpdateUser(updateUserRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// DeleteUser 接口文档
// @Summary 删除用户
// @Description 删除用户数据
// @Tags user
// @Accept  json
// @Produce  application/json
// @Param userId path string true "delete user by id"
// @Success 200 {object} response.Response
// @Router /user/{userId} [delete]
func (controller *UserControllers) DeleteUser(ctx *gin.Context) {
	log.Info().Msg("delete user")
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	deleteUserRequest := request.DeleteUserRequest{
		Id: id,
	}
	controller.userService.DeleteUser(deleteUserRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindUserById 接口文档
// @Summary 查找用户
// @Description 查找用户数据
// @Tags user
// @Accept  json
// @Produce  application/json
// @Param tagId path string true "find user by id"
// @Success 200 {object} response.Response
// @Router /user/{tagId} [get]
func (controller *UserControllers) FindUserById(ctx *gin.Context) {
	log.Info().Msg("find user by id")
	userId := ctx.Param("tagId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	findUserByIdRequest := request.FindUserRequest{
		Id: id,
	}

	userResponse := controller.userService.FindUserById(findUserByIdRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAllUser 接口文档
// @Summary 查找所有用户
// @Description 查找所有用户数据
// @Tags user
// @Accept  json
// @Produce  application/json
// @Success 200 {object} response.Response
// @Router /user [get]
func (controller *UserControllers) FindAllUser(ctx *gin.Context) {
	log.Info().Msg("find all user")

	userResponse := controller.userService.FindAllUser()

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
