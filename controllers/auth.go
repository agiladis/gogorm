package controllers

import (
	"gogorm/services"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	UserService *services.UserService
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ac *AuthController) Login(ctx *gin.Context) {
	var loginDto LoginDto

	err := ctx.ShouldBindJSON(&loginDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	result, err := ac.UserService.GetUserByEmail(loginDto.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "unable to login, check your credential",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(loginDto.Password))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "please check your password",
		})
		return
	}

	claim := jwt.MapClaims{
		"id":    result.ID,
		"email": result.Email,
	}

	signedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	jwtToken, err := signedToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "unable to create jwt token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": jwtToken,
	})
}
