package controllers

import (
	"gogorm/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gogorm/services"
)

type UserController struct {
	UserService *services.UserService
}

type UserCreateDto struct {
	Email string `json:"email" binding:"required"`
}

// CreateUser godoc
// @Summary      Create User
// @Description  Create User
// @Tags         users
// @Param UserCreateDto body UserCreateDto true "Create User"
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.User
// @Router       /users/ [post]
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var (
		userCreateDto UserCreateDto = UserCreateDto{}
		result        entity.User
		err           error
	)

	if err := ctx.ShouldBindJSON(&userCreateDto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user := entity.User{
		Email: userCreateDto.Email,
	}

	if result, err = uc.UserService.Create(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
	// return
}

// GetAllUser godoc
// @Summary      Get All User
// @Description  Get All User Record
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.User
// @Router       /users/ [get]
func (uc *UserController) GetAllUser(ctx *gin.Context) {
	var (
		users = []entity.User{}
		err   error
	)

	users, err = uc.UserService.GetAll()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
	// return
}

// GetUserById() godoc
// @Summary      Get User By ID
// @Description  Get User Record By ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  entity.User
// @Router       /users/{id} [get]
func (uc *UserController) GetUserById(ctx *gin.Context) {
	var (
		userId = ctx.Param("id")
		user   entity.User
	)

	userIdInt, _ := strconv.Atoi(userId)
	user, err := uc.UserService.GetUserById(uint(userIdInt))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "id not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
	// return
}

// DeleteUserById() godoc
// @Summary      Delete User By ID
// @Description  Delete User Record By ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  entity.User
// @Router       /users/{id} [delete]
func (uc *UserController) DeleteUserById(ctx *gin.Context) {
	userId := ctx.Param("id")
	userIdInt, _ := strconv.Atoi(userId)

	if err := uc.UserService.Delete(uint(userIdInt)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "id not found",
		})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
	// return
}

// UpdateUserById() godoc
// @Summary      Update User
// @Description  Update User
// @Tags         users
// @Param UserCreateDto body UserCreateDto true "Update User"
// @Param        id   path      int  true  "User ID"
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.User
// @Router       /users/{id} [put]
func (uc *UserController) UpdateUserById(ctx *gin.Context) {
	var userCreateDto UserCreateDto = UserCreateDto{}

	if err := ctx.ShouldBindJSON(&userCreateDto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userId := ctx.Param("id")
	userIdInt, _ := strconv.Atoi(userId)
	user := entity.User{
		Email: userCreateDto.Email,
	}

	err := uc.UserService.Update(uint(userIdInt), user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "successfully updated data",
	})
	// return
}
