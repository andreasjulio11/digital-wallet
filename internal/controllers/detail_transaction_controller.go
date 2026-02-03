package controllers

import (
	"digital-wallet/internal/dto"
	"digital-wallet/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DetailTransactionController struct {
	DetailTransactionService *services.DetailTransactionService
	UserService              *services.UserService
}

func (ctrl *DetailTransactionController) Saldo(c *gin.Context) {
	var input dto.TransactionRequest

	val, _ := c.Get("user_id")
	floatID := val.(float64)
	userID := int(floatID)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	transaction, err := ctrl.DetailTransactionService.CreateTransaction(userID, input.Amount, input.TypeTransaction)
	user, _ := ctrl.UserService.Profile(int(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "saldo": user.Balanced})
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
		"data": gin.H{
			"id":               transaction.ID,
			"transaction_type": transaction.TransactionType,
			"amount":           transaction.Amount,
			"current_balance":  user.Balanced, // Menampilkan balance saja
			"created_at":       transaction.CreatedAt,
		},
	})

}
