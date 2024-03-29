package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	_ "fmt"
	"log"
	"net/http"
	_ "net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewCampaignRepository(db)
	transactionRepository := transaction.NewRepository(db)

	userService := user.NewService(userRepository)
	campaignService := campaign.NewCampaignService(campaignRepository)
	transactionService := transaction.NewService(transactionRepository, campaignRepository)
	autService := auth.NewServices()

	userHandler := handler.NewUserHandler(userService, autService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmail)
	api.POST("/avatars", authMiddleware(autService, userService), userHandler.UploadAvatar)

	api.GET("/campaigns", campaignHandler.GetCampaign)
	api.GET("/campaigns/:id", campaignHandler.GetCampaignById)
	api.POST("/campaigns", authMiddleware(autService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware(autService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", authMiddleware(autService, userService), campaignHandler.UploadImage)

	api.GET("/campaigns/:id/transactions", authMiddleware(autService, userService), transactionHandler.GetCampaignTransaction)
	api.GET("/transactions", authMiddleware(autService, userService), transactionHandler.GetUserTransaction)
	api.POST("/transactions", authMiddleware(autService, userService), transactionHandler.CreateTransaction)
	router.Run()
}

func authMiddleware(authServices auth.Service, userService user.UserService) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authServices.ValidateToke(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "Salah", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "Gagal", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))

		user, err := userService.GetUserById(userId)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "Passs", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
