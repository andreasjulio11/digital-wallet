package main

import (
	"digital-wallet/internal/config"
	"digital-wallet/internal/controllers"
	"digital-wallet/internal/middleware"
	"digital-wallet/internal/models"
	"digital-wallet/internal/repository"
	"digital-wallet/internal/services"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//cek env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env", err)
	}

	//konek ke database
	db := config.ConnectDB()

	//auto migrate menggunakan gorm
	err = db.AutoMigrate(&models.User{}, &models.DetailTransaction{})
	if err != nil {
		log.Fatal("Gagal Migrasi:", err)
	}
	fmt.Println("AutoMigrate Berhasil!")

	//DI
	userRepo := &repository.UserRepository{DB: db}
	userService := &services.UserService{UserRepository: userRepo}
	userController := &controllers.UserController{UserService: userService}

	dtlTransactionRepo := &repository.DetailTransactionRepository{DB: db}
	dtlTrasanctionService := &services.DetailTransactionService{DetailTransactionRepository: dtlTransactionRepo, UserRepository: userRepo}
	dtlTransactionController := &controllers.DetailTransactionController{DetailTransactionService: dtlTrasanctionService, UserService: userService}

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/register", userController.Register)
		api.POST("/login", userController.Login)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/profile", userController.Profile)

			protected.POST("/topup", dtlTransactionController.Saldo)

			protected.POST("/withdraw", dtlTransactionController.Saldo)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)
	r.Run(":" + port)
}
