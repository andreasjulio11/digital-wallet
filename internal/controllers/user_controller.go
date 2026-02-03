package controllers

import (
	"digital-wallet/internal/dto"
	"digital-wallet/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func (ctrl *UserController) Register(c *gin.Context) {
	var input dto.RegisterRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed", "error": err.Error()})
		return
	}

	user, err := ctrl.UserService.Register(input.Name, input.Email, input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registrasi berhasil", "data": user})
}

func (ctrl *UserController) Login(c *gin.Context) {
	var input dto.LoginRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.UserService.Login(input.Email, input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
	})

}

func (ctrl *UserController) Profile(c *gin.Context) {
	val, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Tidak Ada User ID tersebut",
		})
		fmt.Println(exist)
		return
	}

	floatID := val.(float64)

	userID := int(floatID)

	user, err := ctrl.UserService.Profile((userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "User Tidak Ada",
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"name":  user.Name,
			"saldo": user.Balanced,
		},
	})
}
